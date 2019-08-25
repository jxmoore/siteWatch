// Package controllers similar to MVC, is responsible for controlling the flow and manipulating the data structure (SiteBlock)
package controllers

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/jxmoore/siteWatch/models"
)

// Poll loops throught the SiteBlock slice and polls each endpoint to determine its up down status.
func Poll(siteList *models.SiteBlock, HTTPS bool) {
	for _, site := range siteList.Sites {

		if !strings.Contains(strings.ToLower(site.Address), "http://") && !strings.Contains(strings.ToLower(site.Address), "https://") {
			if HTTPS {
				site.Address = "https://" + site.Address
			} else {
				site.Address = "http://" + site.Address
			}
		} else if HTTPS { // Accounting for hardcoded HTTP in JSON but -t at runtime.
			if !strings.Contains(strings.ToLower(site.Address), "https://") {
				site.Address = strings.Replace(site.Address, "http://", "https://", -1)
			}
		}

		testSite, err := http.Get(site.Address)
		if err != nil {
			site.Count++
			site.Status = false
		}

		if testSite.StatusCode != site.Result {
			site.Count++
			site.Status = false
			fmt.Printf("Error test %v : %v != %v \n", site.Address, testSite.StatusCode, site.Result)
		} else {
			fmt.Printf("Test %s : %d \n", site.Address, testSite.StatusCode)
		}

		// currently just stdout, add slack.
		if site.Count >= site.Threshold {
			notify(site.Address, site.Count)
			site.Count = 0
		}

	}
}

// notify is responsible for notifying when failures exceed the threshold.
func notify(siteName string, count int) {
	fmt.Printf("The test for %s has failed %d times which exceeds the current threshold value.", siteName, count)
}

// StartPoll is responsible for running the Poll func on a loop.
func StartPoll(siteList *models.SiteBlock, HTTPS bool) {
	for {
		Poll(siteList, HTTPS)
		time.Sleep(time.Duration(siteList.Intreval) * time.Second)
	}
}
