// Package models is responsible for the creation and structure of the siteblock type.
package models

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"
)

// site is a struct representing the expected nested json objects with the addition of a count, threshold and status
// these hold the number of failures, the threshold for failures and current up/down status.
type site struct {
	Address   string
	Result    int
	Status    bool
	Count     int
	Threshold int
}

// Notification is the struct that represents the notification json body
type Notification struct {
	Kind    string
	Path    string
	Webhook string
}

// SiteBlock is a type containing the top level JSON object which contains an array of sites and the intreval to loop in seconds.
type SiteBlock struct {
	Sites        []site
	Notification Notification
	Intreval     int
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
		// fmt.Println(err)
		return &SiteBlock{}, errors.New("Error reading json")
	}

	// If the intreval is not the null value and is less than 5 seconds.
	if watchList.Intreval < 5 && watchList.Intreval != 0 {
		watchList.Intreval = 5
	}

	return &watchList, nil
}
