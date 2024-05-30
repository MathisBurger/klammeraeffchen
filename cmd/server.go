package main

import (
	"github.com/kelseyhightower/envconfig"
	"klammerAeffchen/internal"
	"klammerAeffchen/internal/configuration"
	"klammerAeffchen/pkg"
	"log"
)

// Initializes the web server application
func main() {
	var config configuration.Config
	err := envconfig.Process("klammeraeffchen", &config)
	if err != nil {
		log.Fatal(err)
	}
	authChannel := make(chan *pkg.ShortAuthMessage, 1)
	go pkg.ShortShortAuthHandler(authChannel)
	internal.InitializeWebServer(config, authChannel)
}
