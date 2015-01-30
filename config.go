package main

import (
	"io/ioutil"

	"github.com/golang/glog"
	"gopkg.in/yaml.v2"
)

type Config struct {
	Host     string `yaml:"irc_host"`
	Nick     string `yaml:"irc_nick"`
	Name     string `yaml:"irc_name"`
	Password string `yaml:"password"`
	Channel  string `yaml:"default_channel"`
	Port     string `yaml:"http_port"`
	UseTLS   bool   `yaml:"use_tls"`
	CertFile string `yaml:"ssl_cert_file"`
	KeyFile  string `yaml:"ssl_key_file"`
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
