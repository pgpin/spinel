package spinel

import "gopkg.in/yaml.v2"
import "github.com/adam-hanna/randomstrings"

type ConfigAd struct {
	Host                 string
	Port                 int
	Dn                   string
	MaxRequestsPerSecond int
}

type ConfigHtml struct {
	Templates  string
	LoginTitle string
}

type Config struct {
	Debug   bool
	Socket  string
	Expires int64
	Secret  string
	Ad      ConfigAd
	Cidrs   []string
	Html    ConfigHtml
}

func ParseYamlConfiguration(yamlstr *[]byte) (*Config, error) {
	secret, _ := randomstrings.GenerateRandomString(1024)
	config := &Config{Debug: false,
		Socket:  "/tmp/spinel.sock",
		Expires: 24,
		Secret:  secret,
		Ad: ConfigAd{Host: "localhost",
			Port:                 389,
			Dn:                   "dc=CORP",
			MaxRequestsPerSecond: 10},
		Html: ConfigHtml{LoginTitle: "Login with ActiveDirectory credentials",
			Templates: "/usr/local/share/spinel/tmpl"}}
	err := yaml.Unmarshal([]byte(*yamlstr), config)
	if err != nil {
		return nil, err
	}
	return config, nil
}
