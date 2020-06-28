package config

import (
	"github.com/ghodss/yaml"
	"io/ioutil"
	"testing"
)

func Test_config(t *testing.T) {
	yamlfile, err := ioutil.ReadFile("./example_config.yaml")
	if err != nil {
		t.Fatal(err)
	}

	c := BaseInfo{}
	err = yaml.Unmarshal(yamlfile, &c)
	if err != nil {
		t.Fatal(err)
	}

	if c.Mode != 1 {
		t.Fatal("config error")
	}

	if c.Localhost != "121.37.236.234:8080" {
		t.Fatal("config error")
	}
}
