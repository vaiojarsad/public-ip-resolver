package config

import (
	"sync"
)

var (
	instance Manager
	once     sync.Once
)

// Manager handle the different configurations
type Manager interface {
	GetSMTPConfig() *SMTPConfig
	GetResolverConfig() *ResolverConfig
}

/*type LoggerConfig struct {
	Writer io.Writer
	Prefix string
	Flag   int
}*/

// GetInstance returns the singleton instance, creating it if necessary
func GetInstance(file string) (Manager, error) {
	var err error = nil
	once.Do(func() {
		instance, err = newJSONFileBackedConfigManager(file)
	})
	return instance, err
}
