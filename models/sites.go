package models

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
)

type Site struct {
	Address string
	Result  int
	Status  bool
}

type SiteBlock struct {
	Sites []Site
}

// NewSiteStruct reads the file contents and creates a slice of sites.
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
	fmt.Println(watchList.Sites)

	return &watchList, nil
}
