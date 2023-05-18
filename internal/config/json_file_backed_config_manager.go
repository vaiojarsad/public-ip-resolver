package config

import (
	"encoding/json"
	"fmt"
	"net/url"
	"os"
	"path/filepath"
)

type jsonResolverConfig struct {
	URL          string            `json:"url,omitempty"`
	ProviderToIP map[string]string `json:"provider_to_ip,omitempty"`
}

type jsonSMTPConfig struct {
	Host   string `json:"host,omitempty"`
	Port   int    `json:"port,omitempty"`
	Sender string `json:"sender,omitempty"`
	Pass   string `json:"password,omitempty"`
	To     string `json:"to,omitempty"`
}

type jsonConfig struct {
	SMTPConfig     jsonSMTPConfig     `json:"smtp_config,omitempty"`
	ResolverConfig jsonResolverConfig `json:"resolver_config,omitempty"`
}

type jsonFileBackedConfigManager struct {
	smtpConfig     *SMTPConfig
	resolverConfig *ResolverConfig
}

func newJSONFileBackedConfigManager(file string) (Manager, error) {
	if file == "" {
		dir, err := os.UserHomeDir()
		if err != nil {
			return nil, err
		}
		file = filepath.Join(dir, ".public_ip_resolver.json")
	}
	data, err := os.ReadFile(file)
	if err != nil {
		return nil, fmt.Errorf("error reading configuration file: %w", err)
	}
	var c jsonConfig
	err = json.Unmarshal(data, &c)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling configuration file: %w", err)
	}
	smtpConfig := &SMTPConfig{
		Host:   c.SMTPConfig.Host,
		Port:   c.SMTPConfig.Port,
		Sender: c.SMTPConfig.Sender,
		Pass:   c.SMTPConfig.Pass,
		To:     c.SMTPConfig.To,
	}
	u, err := url.Parse(c.ResolverConfig.URL)
	if err != nil {
		return nil, fmt.Errorf("error parsing URL: %w", err)
	}
	resolverConfig := &ResolverConfig{URL: u, ProviderToIP: c.ResolverConfig.ProviderToIP}

	m := &jsonFileBackedConfigManager{
		resolverConfig: resolverConfig,
		smtpConfig:     smtpConfig,
	}
	return m, nil
}

func (cm *jsonFileBackedConfigManager) GetSMTPConfig() *SMTPConfig {
	return cm.smtpConfig
}

func (cm *jsonFileBackedConfigManager) GetResolverConfig() *ResolverConfig {
	return cm.resolverConfig
}
