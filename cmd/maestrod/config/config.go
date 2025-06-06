package config

import (
	"encoding/json"
	"os"

	"github.com/wvdschel/compute-maestro/xdg"
)

const DIRNAME = "compute-maestro"

type Config struct {
	ModelData ModelData `json:"model_data,omitempty"`
	Services  []Service `json:"services,omitempty"`

	Hostname string `json:"hostname,omitempty"`
	Port     int    `json:"port,omitempty"`
	TLS      TLS    `json:"tls,omitempty"`
}

type ServiceType string

type Service struct {
	APIPath string         `json:"path,omitempty"`
	Type    ServiceType    `json:"type,omitempty"`
	Spec    map[string]any `json:"spec"`
	Port    int            `json:"port,omitempty"`
	Logging Logging        `json:"logging,omitempty"`
	Models  []string       `json:"models,omitempty"`
}

type ModelData struct {
	Location string `json:"location,omitempty"`
}

type Logging struct {
	Enabled bool `json:"enabled,omitempty"`
}

type TLS struct {
	Cert string `json:"cert,omitempty"`
	Key  string `json:"key,omitempty"`
}

func LoadFromFile(filename string) (*Config, error) {
	c := &Config{}

	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	err = json.NewDecoder(f).Decode(c)
	if err != nil {
		return nil, err
	}

	return c, nil
}

func Default() *Config {
	return &Config{
		ModelData: ModelData{
			Location: xdg.DataHome() + DIRNAME,
		},
		Services: []Service{
			{
				APIPath: "",
				Type:    "",
				Spec:    map[string]any{},
				Port:    0,
				Logging: Logging{},
				Models:  []string{},
			},
		},
		Hostname: "",
		Port:     18080,
		TLS: TLS{
			Cert: "",
			Key:  "",
		},
	}
}
