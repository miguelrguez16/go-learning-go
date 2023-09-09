package configuration

import (
	"github.com/go-yaml/yaml"
	"os"
)

type Database struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Name     string `yaml:"name"`
}

type Configuration struct {
	DB Database `yaml:"Database"`
}

func Load(filename string) (*Configuration, error) {

	conf := &Configuration{}

	// load yaml file
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	// transfor from byte array to configuration
	err = yaml.Unmarshal(data, conf)
	if err != nil {
		return nil, err
	}
	return conf, err
}
