package controllers

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/jxmoore/siteWatch/models"
)

// httpSiteList is a struct containing a pointer to the main struct in models. This is done so it can be used as a method on the request function.
type httpSiteList struct {
	siteBlock *models.SiteBlock
}

// ServerSiteStats calls the handeler and begins listening on the specified port.
func ServerSiteStats(siteListRef *models.SiteBlock) error {

	serverSiteList := httpSiteList{siteListRef}
	http.HandleFunc("/", serverSiteList.statHandler)
	err := http.ListenAndServe(siteListRef.Port, nil)
	if err != nil {
		return errors.New("Error listening on port")
	}

	return nil
}

// statHandler is a method on httpSiteList that is responsible for returing basic stats when browsing root.
func (serveringList *httpSiteList) statHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Last checked at : %v\n", serveringList.siteBlock.LastChecked)
	for _, site := range serveringList.siteBlock.Sites {
		var siteUp string
		if site.Status {
			siteUp = "Site up"
		} else {
			siteUp = "Currently down"
		}
		fmt.Fprintf(w, "Site : %v \n Status : %v \n FailureCount : %v \n Threshold : %v \n Expected Response Code : %v \n Last Response Code : %v \n\n", site.Address, siteUp, site.Count, site.Threshold, site.Result, site.LastResultCode)
	}
}
