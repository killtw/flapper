package config

import (
	"bytes"
	"github.com/spf13/viper"
	"io/ioutil"
)

var defaultConfig = []byte(`
home: "~/Downloads"
`)

type Configure struct {
	Home string `yaml:"path"`
}

func LoadConfig(path string) (Configure, error) {
	var conf Configure

	viper.SetConfigType("yaml")

	if path != "" {
		content, err := ioutil.ReadFile(path)

		if err != nil {
			return conf, err
		}

		if err := viper.ReadConfig(bytes.NewBuffer(content)); err != nil {
			return conf, err
		}
	} else {
		viper.AddConfigPath("$HOME")
		viper.AddConfigPath("$HOME/.flapper")
		viper.AddConfigPath(".")
		viper.SetConfigName("flapper")

		if err := viper.ReadInConfig(); err != nil {
			if err := viper.ReadConfig(bytes.NewBuffer(defaultConfig)); err != nil {
				return conf, err
			}
		}
	}

	conf.Home = viper.GetString("home")

	return conf, nil
}
