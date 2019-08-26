package controllers

import (
	"fmt"
	"net/http"

	"github.com/jxmoore/siteWatch/models"
)

// httpSiteList is a struct containing a pointer to the main struct in models. This is done so it can be used as a method on the request function.
type httpSiteList struct {
	siteBlock *models.SiteBlock
}

func ServerSiteStats(siteListRef *models.SiteBlock) {

	serverSiteList := httpSiteList{siteListRef}
	http.HandleFunc("/", serverSiteList.statHandler)
	http.ListenAndServe(":8080", nil)

}

func (serveringList *httpSiteList) statHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Last checked at : %v\n", serveringList.siteBlock.LastChecked)
	for _, site := range serveringList.siteBlock.Sites {
		var siteUp string
		if site.Status {
			siteUp = "Site up"
		} else {
			siteUp = "Currently down"
		}
		fmt.Fprintf(w, "Site : %v \n Status : %v \n Expected responseCode : %v \n FailureCount : %v\n\n", site.Address, siteUp, site.Result, site.Count)
	}
}
