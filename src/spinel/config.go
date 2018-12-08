package spinel

import "gopkg.in/yaml.v2"

type ConfigAd struct {
	Host string
	Port int
	Dn  string
	MaxRequestsPerSecond int
}

type Config struct {
	Listen string
	Secret string
	Ad     ConfigAd
	Cidrs  []string
}

func ParseYamlConfiguration(yamlstr *[]byte) (*Config, error) {
	var config Config
	err := yaml.Unmarshal([]byte(*yamlstr), &config)
	if err != nil {
		return nil, err
	}
	return &config, nil
}
