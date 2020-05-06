package main

import (
	"bitbucket.org/telemetryapp/go_rabbitmq/config"
	"bitbucket.org/telemetryapp/go_rabbitmq/gateway"
	"gopkg.in/alecthomas/kingpin.v2"
	"log"
)

var configPath = kingpin.Flag("config", "Path to config file.").Short('c').ExistingFile()

func main () {

	// Load config
	conf := config.ReadConfig(*configPath)

	// Create and start a Gateway instance
	log.Println("Creating API gateway")
	gtw, err := gateway.New(conf)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(gtw)

}
