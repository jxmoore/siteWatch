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

	siteConfig := models.SiteConfig{}
	err := siteConfig.LoadSiteConfig(*confFile)
	if err != nil {
		log.Fatal(err.Error())
	}

	if *aiKey == "" {
		log.Fatal("Missing required applicationinsights instrumentation key.")
	}
	poll.RunSitePoll(&siteConfig, *aiKey)

}
