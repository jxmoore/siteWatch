package main

import (
	"flag"
	"log"

	"github.com/jxmoore/siteWatch/models"
	"github.com/jxmoore/siteWatch/poll"
)

func main() {
	var confFile = flag.String("p", "./config.json", "The config file for the sites to watch.")
	var aiKey = flag.String("k", "", "The instrumentation key to use.")

	flag.Parse()

	if *aiKey == "" || *confFile == "" {
		log.Fatal("Missing required flag.")
	}
	pollConfig := models.SiteConfig{}
	err := pollConfig.LoadSiteConfig(*confFile)
	if err != nil {
		log.Fatal(err.Error())
	}

	poll.RunSitePoll(&pollConfig, *aiKey)

}
