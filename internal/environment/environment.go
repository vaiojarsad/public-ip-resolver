package environment

import (
	"sync"

	"github.com/vaiojarsad/public-ip-resolver/internal/config"
)

// Environment holds the execution context
type Environment struct {
	ConfigManager config.Manager
}

var (
	instance *Environment
	once     sync.Once
)

// GetInstance returns the singleton instance, creating it if necessary
func GetInstance() *Environment {
	once.Do(func() {
		instance = &Environment{}
	})
	return instance
}
