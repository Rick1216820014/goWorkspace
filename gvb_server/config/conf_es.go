package config

import "fmt"

type ES struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	DB       string `yaml:"db"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
}

func (es ES) Url() string {
	return fmt.Sprintf("%s:%d", es.Host, es.Port)
}
