package config

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"time"
)

type Config struct {
	HttpIPv4BindAddress string `yaml:"ipv4_http_bind_address"`
	HttpIPv4BindPort int `yaml:"ipv4_http_bind_port"`
	HttpReadTimeout time.Duration `yaml:"http_read_timeout"`
	HttpWriteTimeout time.Duration `yaml:"http_write_timeout"`
	HttpIdleTimeout time.Duration `yaml:"http_idle_timeout"`
}

func NewConfig() *Config {
	var C Config
	return &C
}

func (C *Config) FromEnv() (*Config, error) {
	envVars := []string{
		"HTTPRTR_IPV4_BIND_ADDR",
		"HTTPRTR_IPV4_BIND_PORT",
		"HTTPRTR_READ_TIMEOUT",
		"HTTPRTR_WRITE_TIMEOUT",
		"HTTPRTR_IDLE_TIMEOUT",
	}

	for _, i := range envVars {
		switch i {
		case "HTTPRTR_IPV4_BIND_ADDR":
			ipv4BindAddr := os.Getenv(i)
			C.HttpIPv4BindAddress = ipv4BindAddr
		case "HTTPRTR_IPV4_BIND_PORT":
			ipv4BindPort := os.Getenv(i)
			d, err := strconv.Atoi(ipv4BindPort)
			if err != nil {
				return nil, err
			}
			C.HttpIPv4BindPort = d
		case "HTTPRTR_READ_TIMEOUT":
			httpReadTimeout := os.Getenv(i)
			d, err := strconv.Atoi(httpReadTimeout)
			if err != nil {
				return nil, err
			}
			dur := time.Duration(d)
			C.HttpReadTimeout = dur
		case "HTTPRTR_WRITE_TIMEOUT":
			httpWriteTimeout := os.Getenv(i)
			d, err := strconv.Atoi(httpWriteTimeout)
			if err != nil {
				return nil, err
			}
			dur := time.Duration(d)
			C.HttpWriteTimeout = dur
		case "HTTPRTR_IDLE_TIMEOUT":
			httpIdleTimeout := os.Getenv(i)
			d, err := strconv.Atoi(httpIdleTimeout)
			if err != nil {
				return nil, err
			}
			dur := time.Duration(d)
			C.HttpIdleTimeout = dur
		}
	}

	return C, nil
}

func (C *Config) FromFile() (*Config, error) {
	log.Println("reading from configuration file 'appConfig.yaml'")
	info, err := os.Stat("appConfig.yaml")
	if err != nil {
		return nil, err
	}

	f, err := os.OpenFile("appConfig.yaml", os.O_RDONLY, info.Mode())
	if err != nil {
		return nil, err
	}
	defer f.Close()

	b, err := ioutil.ReadAll(f)

	err = yaml.Unmarshal(b, C)
	if err != nil {
		return nil, err
	}

	return C, nil
}