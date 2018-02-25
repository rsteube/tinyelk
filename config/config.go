package config

import (
	"fmt"
	"github.com/go-yaml/yaml"
	"io/ioutil"
)

type Config struct {
	Grok struct {
		Patterns struct {
			Base      map[string]string
			Line      map[string]string
			Multiline map[string]string
		}
		Drop []string
	}
}

func Load() (c Config, err error) {
	if yamlFile, err := ioutil.ReadFile(".tinyelk.yml"); err == nil {
		err = yaml.Unmarshal(yamlFile, &c)
		if err != nil {
			// TODO duh!
			fmt.Println(err)
		}
	}
	return
}
