package spinel

import "testing"

func TestParseYamlConfig(t *testing.T) {
	yamlstr := `
---
secret: foobar
listen: 127.0.0.1:9999
ad:
  host: hostname
  port: 123
  dsn: dsntest
cidrs:
  - 123.123.123.123/32
`
	b := []byte(yamlstr)
	config, err := ParseYamlConfiguration(&b)
	if err != nil {
		t.Error("could not parse yaml string", err)
	}
	if config.Secret != "foobar" {
		t.Error("configuration could not read secret")
	}
	if config.Ad.Host != "hostname" {
		t.Error("configuration could not read hostname")
	}
	if config.Cidrs[0] != "123.123.123.123/32" {
		t.Error("configuration could not read cidrs")
	}
	if config.Listen != "127.0.0.1:9999" {
		t.Error("configuration could not read listen")
	}
}
