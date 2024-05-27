package main

import (
	"github.com/kelseyhightower/envconfig"
	"klammerAeffchen/internal"
	"klammerAeffchen/internal/configuration"
	"log"
)

func main() {
	var config configuration.Config
	err := envconfig.Process("klammeraeffchen", &config)
	if err != nil {
		log.Fatal(err)
	}
	go internal.InitializeWebServer(config)
	internal.InitializeDiscordBot(&config)
}
