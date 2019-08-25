// Package models is responsible for the creation and structure of the siteblock type.
package models

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
)

// Site is a struct representing the expected nested json objects with the addition of a count and status
// these hold the number of failures and current up/down status.
type Site struct {
	Address string
	Result  int
	Status  bool
	Count   int
}

//SiteBlock is the top level JSON object which contains an array of sites.
type SiteBlock struct {
	Sites []Site
}

// NewSiteStruct reads the file contents and returns a pointer to a SiteBlock.
func NewSiteStruct(filePath string) (*SiteBlock, error) {
	fileEx, err := os.Stat(filePath)
	if os.IsNotExist(err) || fileEx.IsDir() == true {
		return &SiteBlock{}, errors.New("The referenced file cannot be read in its current state")
	}

	contents, err := ioutil.ReadFile(filePath)
	if err != nil {
		return &SiteBlock{}, errors.New("Error reading file")
	}

	watchList := SiteBlock{}
	err = json.Unmarshal(contents, &watchList)
	if err != nil {
		fmt.Println(err)
		return &SiteBlock{}, errors.New("Error reading json")
	}
	return &watchList, nil
}
