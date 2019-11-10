// Package poll contains the functions for running the poll against the sites defined in a SiteConfig.
package poll

import (
	"crypto/tls"
	"fmt"
	"net/http"
	"sync"
	"time"

	"github.com/microsoft/ApplicationInsights-Go/appinsights"

	"github.com/jxmoore/siteWatch/models"
)

// RunSitePoll is an exported function that starts a go routine for each site in the SiteConfig struct.
// The WG is awaited for completion but the wait groups never signal completion.
func RunSitePoll(sites *models.SiteConfig, iKey string) error {

	fmt.Printf("Starting poll... \n\tConfig :%v\n\tKey : %v\n\n", sites, iKey)

	var wg sync.WaitGroup
	var client = appinsights.NewTelemetryClient(iKey)
	wg.Add(len(sites.SiteBlock))

	for _, s := range sites.SiteBlock {
		fmt.Printf("Starting go routine for test %v\n", s.Name)
		go sitePoll(s.Name, s.TestEndpoint, s.Intreval, s.Response, s.Timeout, client)
	}

	wg.Wait()
	return nil
}

// sitePoll is the main polling function, it loops indefinitely, probing the specific endpoint ever 'x' seconds. It reuses its own HTTP client
// that has a timeout defined by the siteconfig block.
func sitePoll(name, endpoint string, intreval, responseCode, timeout int, client appinsights.TelemetryClient) error {

	testResults := models.Availability{Client: client, Name: name}
	tr := &http.Transport{TLSClientConfig: &tls.Config{InsecureSkipVerify: true}}
	testSite := &http.Client{Timeout: time.Second * time.Duration(timeout), Transport: tr}

	for {

		testResults.Start = time.Now()
		resp, err := testSite.Get(endpoint)
		testResults.Time = time.Since(testResults.Start)
		testResults.End = time.Now()
		if err != nil {
			fmt.Println(err.Error())
			testResults.Success = false
			testResults.Msg = fmt.Sprintf("Test %v : returned an error : %v", name, err.Error())
		} else if resp.StatusCode != responseCode {
			defer resp.Body.Close()
			testResults.Msg = fmt.Sprintf("%v responded with the expected status code of %v.", endpoint, responseCode)
			testResults.Success = false
		} else {
			defer resp.Body.Close()
			testResults.Msg = fmt.Sprintf("%v responded with the expected status code of %v.", endpoint, responseCode)
			testResults.Success = true
		}

		fmt.Printf(testResults.Msg)
		fmt.Println("")

		_ = testResults.SendAvailibiltyStats()
		time.Sleep(time.Duration(intreval) * time.Second)

	}

	return nil

}
