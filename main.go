package main

import (
	"flag"
	"log"

	"github.com/jxmoore/siteWatch/controllers"
	"github.com/jxmoore/siteWatch/models"
)

func main() {

	var file = flag.String("f", ".\\sites.json", "A relative path to the JSON file containing the sites to monitor.")
	var HTTPS = flag.Bool("t", false, "Get requests are attempted over HTTPS rather than HTTP")

	flag.Parse()
	watchList, err := models.NewSiteStruct(*file)
	if err != nil {
		log.Fatal(err)
	}

	controllers.StartPoll(watchList, *HTTPS)
	// controllers.CleanAddress(watchList, *HTTPS)
	// controllers.Poll(watchList)
}
