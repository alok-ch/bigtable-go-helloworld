package config

import (
	"github.com/mitchellh/mapstructure"
	"os"
)

type Config struct {
	Project  string
	Instance string
}

var AppConfigList = []string{
	"PROJECT",
	"INSTANCE",
}

type AppConfigProvider interface {
	ProvideEnv([]string) (map[string]string, error)
}

type EnvAppConfigProvider struct {
}

func (envAppCfgProvider *EnvAppConfigProvider) ProvideEnv(vars []string) (map[string]string, error) {
	envMap := make(map[string]string)
	for _, val := range vars {
		if os.Getenv(val) != "" {
			envMap[val] = os.Getenv(val)
		} else {
			panic("env config not found for " + val)
		}
	}
	return envMap, nil
}

func ConstructAppConfig(envMap map[string]string) *Config {
	cfg := &Config{}
	mapstructure.Decode(envMap, cfg)
	return cfg
}
