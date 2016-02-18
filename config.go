package goldshire

import (
	"encoding/json"
	"os"
)

type Config struct {
	Listen string       `json:"listen"`
	Domain string       `json:"domain"`
	Gitlab GitlabConfig `json:"gitlab"`
}

type GitlabConfig struct {
	ApiBase      string `json:"api_base"`
	PrivateToken string `json:"private_token"`
}

func GetConfig(name string) (*Config, error) {
	f, err := os.Open(name)
	if err != nil {
		return nil, err
	}

	cfg := &Config{}

	decoder := json.NewDecoder(f)
	if err := decoder.Decode(cfg); err != nil {
		return nil, err
	}

	return cfg, nil
}
