package config

import "net/url"

type ResolverConfig struct {
	URL          *url.URL
	ProviderToIP map[string]string
}
