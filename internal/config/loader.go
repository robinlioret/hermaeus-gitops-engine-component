package config

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/viper"
)

func Load(configPath string) (*Config, error) {
	v := viper.New()

	setDefaults(v)
	bindEnvVars(v)

	if configPath != "" {
		v.SetConfigFile(configPath)
	} else if os.Getenv("HEGEC_CONFIG_PATH") != "" {
		v.SetConfigFile(os.Getenv("HEGEC_CONFIG_PATH"))
	} else {
		v.SetConfigName("config")
		v.SetConfigType("yaml")
		v.AddConfigPath(".")
		v.AddConfigPath("./config")
		v.AddConfigPath("./etc/hegec")
	}

	var cfg Config
	if err := v.Unmarshal(&cfg); err != nil {
		return nil, fmt.Errorf("failed to unmarshal config: %w", err)
	}

	if err := validate(&cfg); err != nil {
		return nil, fmt.Errorf("failed to validate config: %w", err)
	}

	return &cfg, nil
}

func setDefaults(v *viper.Viper) {
	// TODO: add new configuration defaults here
	v.SetDefault("dummy ", false)
}

func bindEnvVars(v *viper.Viper) {
	v.SetEnvPrefix("HEGEC")
	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	v.AutomaticEnv()

	// TODO: add new nested configuration keys configuration environment variables here
	enVars := []string{}
	for _, envVar := range enVars {
		_ = v.BindEnv(envVar)
	}
}
