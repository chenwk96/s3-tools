package config

type Config struct {
	Endpoint        string `yaml:"endpoint"`
	AccessKeyID     string `yaml:"access_key_id"`
	SecretAccessKey string `yaml:"secret_access_key"`
	UseSSL          bool   `yaml:"use_ssl"`
	Host            string `yaml:"host"`
	Id              string `yaml:"id"`
	Key             string `yaml:"key"`
	Region          string `yaml:"region"`
	Version         string `yaml:"version"`
}

var Cfg Config
