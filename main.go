package main

import (
	"flag"

	"github.com/jxmoore/siteWatch/models"
)

func main() {
	var file = flag.String("f", ".\\sites.json", "A relative path to the JSON file containing the sites to monitor.")
	//var example = flag.Bool("-e", false, "Display an example file.")
	flag.Parse()
	models.NewSiteStruct(*file)
}
