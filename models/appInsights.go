package models

import (
	"fmt"
	"os"
	"time"

	"github.com/microsoft/ApplicationInsights-Go/appinsights"
)

// Availability is a struct that holds the information regarding a specific Availability test along with the client that is used
// to send telemetry.
type Availability struct {
	Name    string
	Time    time.Duration
	Success bool
	Client  appinsights.TelemetryClient
	Start   time.Time
	End     time.Time
	Msg     string
}

// SendAvailibiltyStats is a method on the Availability struct. It uses the receiver passed in to send availibility test results
// to appinsights. It returns error but the track method used does not have a return so as it stands this will always return nil.
func (a Availability) SendAvailibiltyStats() error {

	hostname, err := os.Hostname()
	if err != nil {
		hostname = "Unknown"
	}
	location, err := time.LoadLocation("EST")
	if err != nil {
		fmt.Println(err)
	}

	availability := appinsights.NewAvailabilityTelemetry(a.Name, a.Time, a.Success)
	availability.RunLocation = "LSPAKS-" + hostname
	availability.Message = fmt.Sprintf("%v. Start - Stop %v : %v", a.Msg, a.Start.In(location), a.End.In(location))
	availability.Id = fmt.Sprintf("%v:%v", a.Name, a.End)
	availability.MarkTime(a.Start, a.End)

	a.Client.Track(availability)

	return nil
}
