// Package poll
package poll

import (
	"fmt"
	"net/http"
	"sync"
	"time"

	"github.com/microsoft/ApplicationInsights-Go/appinsights"

	"github.com/jxmoore/siteWatch/models"
)

func RunSitePoll(sites *models.SiteConfig, iKey string) error {

	var wg sync.WaitGroup
	var client = appinsights.NewTelemetryClient(iKey)

	for _, s := range sites.SiteBlock {
		wg.Add(1)
		fmt.Printf("Starting go routine for test %v\n", s.Name)
		go sitePoll(s.Name, s.TestEndpoint, s.Intreval, s.Response, s.Timeout, &wg, client)
	}

	wg.Wait()
	return nil
}

func sitePoll(name, endpoint string, intreval, responseCode, timeout int, wg *sync.WaitGroup, client appinsights.TelemetryClient) error {

	testResults := models.Availability{Client: client, Name: name}
	testSite := &http.Client{Timeout: time.Second * time.Duration(timeout)}
	defer wg.Done()

	for {

		testResults.Start = time.Now()
		resp, err := testSite.Get(endpoint)
		testResults.Time = time.Since(testResults.Start)
		testResults.End = time.Now()

		defer resp.Body.Close()

		if err != nil {
			fmt.Println(err.Error())
			testResults.Success = false
			testResults.Msg = fmt.Sprintf("Test %v : returned an error : %v \n", name, err.Error())
		} else if resp.StatusCode != responseCode {
			testResults.Msg = fmt.Sprintf("Test %v : received status code : %v expected status code %v\n", name, resp.StatusCode, responseCode)
			testResults.Success = false
		} else {
			testResults.Msg = fmt.Sprintf("%v : %v took %v seconds and responded with the expected status code of %v\n", name, endpoint, testResults.Time, responseCode)
			testResults.Success = true
		}

		fmt.Printf(testResults.Msg)
		testResults.SendAvailibiltyStats()
		time.Sleep(time.Duration(intreval) * time.Second)
	}

}
