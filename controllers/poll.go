// Package controllers similar to MVC, is responsible for controlling the flow and manipulating the data structure (SiteBlock)
package controllers

import (
	"net/http"
	"sync"
	"time"

	"github.com/jxmoore/siteWatch/models"
)

// TimeFormat is a package wide variable used to format the time into a more user friendly string.
var timeFormat = "Jan 2, 2006 3:04pm (MST)"

// StartPoll is responsible for running the Poll func on a loop.
func StartPoll(siteList *models.SiteBlock) {

	if siteList.HTTP {
		go ServerSiteStats(siteList)
	}

	for {
		poll(siteList)
		time.Sleep(time.Duration(siteList.Intreval) * time.Second)
	}
}

// poll loops throught the SiteBlock slice and calls sitecheck on a new goroutine for each site
func poll(siteList *models.SiteBlock) {
	var wg sync.WaitGroup
	for x := range siteList.Sites {
		wg.Add(1)
		go siteCheck(siteList, x, &wg)
	}
	wg.Wait()
	siteList.LastChecked = time.Now().Format(timeFormat)
}

// siteCheck performs a get against the site to determine if the status code matches the resultcode in the SiteBlock struct
func siteCheck(siteList *models.SiteBlock, index int, wg *sync.WaitGroup) {

	defer wg.Done()
	site := siteList.Sites[index]

	testSite, err := http.Get(site.Address)
	if err != nil {
		site.Count++
		site.Status = false
		site.LastResultCode = 0
	} else if testSite.StatusCode != site.Result {
		site.Count++
		site.Status = false
		site.LastResultCode = testSite.StatusCode
	} else {
		if !site.Status {
			site.Status = true
			site.LastResultCode = testSite.StatusCode
		}
	}

	if site.Count >= site.Threshold && site.Threshold != 0 {
		notify(site.Address, site.Count, &siteList.Notification)
		site.Count = 0
	}

	siteList.Sites[index] = site
}
