package config

import "testing"

func TestLoadConfig(t *testing.T) {
	conf, err := LoadConfig("")
	if err != nil {
		t.Error(err)
	}
	if conf.Home != "~/Downloads" {
		t.Fatal("err")
	}
	conf, err = LoadConfig("./tests/config.yaml")
	if err != nil {
		t.Error(err)
	}
	if conf.Home != "~/Test" {
		t.Fatal("err")
	}
}
