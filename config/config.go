package config

import (
	"github.com/ghodss/yaml"
	"io/ioutil"
)

type BaseInfo struct {
	Localhost string `yaml:"localhost"`
	Mode      int    `yaml:"mode"`
}

func (b *BaseInfo) LoadFile(path string) error {
	yamlfile, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}
	err = yaml.Unmarshal(yamlfile, b)
	if err != nil {
		return err
	}

	return nil
}
