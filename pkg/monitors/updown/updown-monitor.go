// Package UpdownMonitor adds updown website monitoring tool's support in IngressMonitorController
package updown

import (
	"log"
	"net/http"

	"github.com/antoineaugusti/updown"
	"github.com/stakater/IngressMonitorController/pkg/config"
	"github.com/stakater/IngressMonitorController/pkg/constants"
	"github.com/stakater/IngressMonitorController/pkg/models"
)

// UpdownMonitorService struct contains parameters required by updown go client
type UpdownMonitorService struct {
	apiKey        string
	url           string
	alertContacts string
	client        *updown.Client
	statusCodes   constants.HTTPStatusCodesMap
}

// Setup function will create a updown's go client object by using the configuration parameters
func (updownService *UpdownMonitorService) Setup(confProvider config.Provider) {
	// configuration parameters for creating a updown client
	updownService.apiKey = confProvider.ApiKey
	updownService.alertContacts = confProvider.AlertContacts
	// creating updown go client
	httpDefaultClient := http.DefaultClient
	updownService.client = updown.NewClient(updownService.apiKey, http.DefaultClient)
}

// GetAll function will return all checks object in an array
func (updownService *UpdownMonitorService) GetAll() []models.Monitor {

	var monitors []models.Monitor

	// getting all checks list
	updownChecks, httpResponse, err := updownService.client.Check.List()

	if (httpResponse.StatusCode == updownService.statusCodes.OK) && (err == nil) {

		// populating a monitors array using the updownChecks objects given in updownChecks slice
		for _, updownCheck := range updownChecks {
			newMonitor := models.Monitor{
				URL:   updownCheck.URL,
				NAME:  updownCheck.Alias,
				TOKEN: updownCheck.Token,
			}
			monitors = append(monitors, newMonitor)
		}
		return monitors
	} else {
		log.Println("Unable to get updown provider checks")
	}

}