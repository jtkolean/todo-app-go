package config

import (
	"io/ioutil"
	"log"

	yaml "gopkg.in/yaml.v3"
)

type config struct {
	Server struct {
		Port     int16 `yaml:"port"`
		Database struct {
			Host        string `yaml:"host"`
			Port        int16  `yaml:"port"`
			Driver      string `yaml:"driver"`
			User        string `yaml:"user"`
			Password    string `yaml:"password"`
			DbName      string `yaml:"dbname"`
			SslMode     string `yaml:"sslmode"`
			SslRootCert string `yaml:"sslrootcert"`
			SslKey      string `yaml:"sslkey"`
			SslCert     string `yaml:"sslcert"`
		} `yaml:"database"`
	} `yaml:"server"`
}

func NewConfig(filename string) config {
	p := config{}

	b, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatalf("failed to read file %v, error %v", filename, err.Error())
	}

	if err := yaml.Unmarshal(b, &p); err != nil {
		log.Fatalf("failed to unmarshal file %v, error %v", filename, err.Error())
	}

	return p
}
