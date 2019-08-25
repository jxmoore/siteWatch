// Package controllers similar to MVC, is responsible for controlling the flow and manipulating the data structure (SiteBlock)
package controllers

import (
	"fmt"
	"net/http"
	"strings"
	"sync"
	"time"

	"github.com/jxmoore/siteWatch/models"
)

// StartPoll is responsible for running the Poll func on a loop.
func StartPoll(siteList *models.SiteBlock, HTTPS bool) {
	cleanAddress(siteList, HTTPS)
	for {
		// fmt.Println("Calling poll...")
		poll(siteList)
		time.Sleep(time.Duration(siteList.Intreval) * time.Second)
	}
}

// cleanAddress is responsible for appending HTTP:// onto the site address, or converting them from HTTP:// to HTTPS://
func cleanAddress(siteList *models.SiteBlock, HTTPS bool) {

	for x, site := range siteList.Sites {
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

		siteList.Sites[x] = site

	}
}

// poll loops throught the SiteBlock slice and calls the sitecheck creating a new goroutine for each site
func poll(siteList *models.SiteBlock) {
	var wg sync.WaitGroup
	for x := range siteList.Sites {
		wg.Add(1)
		go siteCheck(siteList, x, &wg)
	}
	wg.Wait()
}

// siteCheck performs a get against the site to determine if the status code matches the resultcode in the SiteBlock struct
func siteCheck(siteList *models.SiteBlock, index int, wg *sync.WaitGroup) {

	defer wg.Done()
	site := siteList.Sites[index]

	testSite, err := http.Get(site.Address)
	if err != nil {
		site.Count++
		site.Status = false
	} else if testSite.StatusCode != site.Result {
		site.Count++
		site.Status = false
		fmt.Printf("Error test %v : %v != %v \n", site.Address, testSite.StatusCode, site.Result)
	} else {
		// fmt.Printf("Test %s : %d \n", site.Address, testSite.StatusCode)
		if !site.Status {
			site.Status = true
		}
	}

	if site.Count >= site.Threshold && site.Threshold != 0 {
		notify(site.Address, site.Count)
		site.Count = 0
	}

	siteList.Sites[index] = site
}

// notify is responsible for notifying when failures exceed the threshold.
func notify(siteName string, count int) {
	fmt.Printf("The test for %s has failed %d times which exceeds the current threshold value.\n", siteName, count)
}
