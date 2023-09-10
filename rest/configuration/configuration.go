package configuration

import (
	"fmt"
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
	// declare struct
	conf := &Configuration{}

	// load yaml file
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	// transform from byte array to configuration struct
	err = yaml.Unmarshal(data, conf)
	if err != nil {
		return nil, err
	}
	return conf, err
}

func PrintConfiguration(configuration *Configuration) {
	fmt.Printf("host [%v]. port [%v]. user [%v]. password [%v]. name [%v].\n",
		configuration.DB.Host,
		configuration.DB.Port,
		configuration.DB.User,
		configuration.DB.Password,
		configuration.DB.Name)
}
