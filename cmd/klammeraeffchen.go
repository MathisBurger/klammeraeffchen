package main

import (
	"github.com/kelseyhightower/envconfig"
	"klammerAeffchen/internal"
	"log"
)

func main() {
	var config internal.Config
	err := envconfig.Process("klammeraeffchen", &config)
	if err != nil {
		log.Fatal(err)
	}
	go internal.InitializeDiscordBot(&config)
	internal.InitializeWebServer(config)
}
