package main

import (
	"flag"
	"log"
	"github.com/Gregory-Development/http-router/config"
	"github.com/Gregory-Development/http-router/server"
)

func main(){
	var (
		conf *config.Config
		svr *server.Server
		err error
	)

	configSource := flag.String("config-src", "file", "The config options source (default: file - options: file, env)")
	flag.Parse()

	if *configSource == "file" {
		conf, err = config.NewConfig().FromFile()
		if err != nil {
			log.Fatalf("unable to gather configurations from file, err: %v", err)
		}
	} else if *configSource == "env" {
		conf, err = config.NewConfig().FromEnv()
		if err != nil {
			log.Fatalf("unable to gather configurations from environment, err: %v", err)
		}
	} else {
		conf = nil
		log.Fatalf("invalid 'config-src' option '%v' provided, valid options are 'file' or 'env'")
	}

	svr = server.NewServer(conf)

	svr.Run()
}
