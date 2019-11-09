// Package models conatins the structs, the SiteConfig (the configuration) and the Availability struct, along with their defined methods.
package models

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

// SiteConfig is the struct that represents the JSON that is passed in on runtime. It is used to determine what to probe, how often etc...
type SiteConfig struct {
	SiteBlock []struct {
		Name         string
		Address      string
		Route        string
		Response     int
		TestEndpoint string
		Intreval     int
		Timeout      int
	}
}

// LoadSiteConfig reads the file contents and unmarshells it into a pointer of the receiver
// it is exported as its called in main.
func (s *SiteConfig) LoadSiteConfig(filePath string) error {
	fileEx, err := os.Stat(filePath)
	if os.IsNotExist(err) || fileEx.IsDir() == true {
		return fmt.Errorf("The referenced file cannot be read in its current state : %v", err.Error())
	}

	contents, err := ioutil.ReadFile(filePath)
	if err != nil {
		return fmt.Errorf("Error reading file.\n%v", err.Error())
	}

	err = json.Unmarshal(contents, s)
	if err != nil {
		return fmt.Errorf("Error reading json\n%v", err.Error())
	}

	s.cleanAddress()

	return nil
}

// cleanAddress is a method on SiteConfig responsible for prefixing HTTPs:// onto the site address and supplying default values to the struct.
// Could be merged into LoadSite.... and then turned into an init
func (s *SiteConfig) cleanAddress() {

	for x, site := range s.SiteBlock {

		if !strings.Contains(strings.ToLower(site.Address), "http://") && !strings.Contains(strings.ToLower(site.Address), "https://") {
			site.Address = "https://" + site.Address
		} else if !strings.Contains(strings.ToLower(site.Address), "https://") {
			site.Address = strings.Replace(site.Address, "http://", "https://", -1)
		}

		if site.Name == "" {
			site.Name = site.Address
		}

		if site.Intreval == 0 {
			site.Intreval = 90
		}

		if site.Timeout == 0 {
			site.Timeout = 15
		}

		site.TestEndpoint = site.Address + site.Route
		s.SiteBlock[x] = site
	}

}
