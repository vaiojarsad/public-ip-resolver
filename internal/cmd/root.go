// Package cmd is used to define Cobra stuff
package cmd

import (
	"github.com/spf13/cobra"
	"github.com/vaiojarsad/public-ip-resolver/internal/environment"

	"github.com/vaiojarsad/public-ip-resolver/internal/config"
	"github.com/vaiojarsad/public-ip-resolver/internal/public_ip_resolver"
)

var (
	rootCmd = &cobra.Command{
		Use:   "public_ip_resolver",
		Short: "Resolve public IPs",
		Long:  "Resolve the public IP assigned by the given ISP",
		RunE:  rootRunE,
	}
	provider, configFile string
)

func init() {
	rootCmd.Flags().StringVarP(&provider, "provider", "p", "", "provider to resolve assigned public IP address for")
	rootCmd.Flags().StringVarP(&configFile, "config", "c", "", "config file")

	err := rootCmd.MarkFlagRequired("provider")
	if err != nil {
		cobra.CheckErr(err)
	}
}

func Execute() {
	_ = rootCmd.Execute()
}

func rootRunE(c *cobra.Command, _ []string) error {
	c.SilenceUsage = true
	cm, err := config.GetInstance(configFile)
	if err != nil {
		return err
	}
	env := environment.GetInstance()
	env.ConfigManager = cm
	return public_ip_resolver.Resolve(provider)
}
