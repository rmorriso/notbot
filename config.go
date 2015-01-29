package main

import (
	"io/ioutil"

	"github.com/golang/glog"
	"gopkg.in/yaml.v2"
)

type Config struct {
	Port    string `yaml:"http_port"`
	Host    string `yaml:"irc_host"`
	Nick string `yaml:"irc_nick"`
	Name string `yaml:"irc_name"`
	Password string `yaml:"password"`
	Channel string `yaml:"default_channel"`
}

// Init unmarshalls Config from YAML configuration in filename
func Init(filename string) (*Config, error) {
	defer glog.Flush()
	var config = new(Config)
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return config, err
	}
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		return config, err
	}
	glog.V(5).Infof("read config %v\n", config)
	return config, err
}
