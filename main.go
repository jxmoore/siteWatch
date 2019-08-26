package main

import (
	"flag"
	"log"

	"github.com/jxmoore/siteWatch/controllers"
	"github.com/jxmoore/siteWatch/models"
)

func main() {

	var file = flag.String("f", ".\\sites.json", "A relative path to the JSON configuration file containing the sites to monitor and other settings.")
	var HTTPS = flag.Bool("t", false, "Get requests are attempted over HTTPS rather than HTTP")

	flag.Parse()
	watchList, err := models.LoadSiteConfig(*file, *HTTPS)
	if err != nil {
		log.Fatal(err)
	}

	controllers.StartPoll(watchList)
}
