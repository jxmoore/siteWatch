package models

import (
	"fmt"
	"os"
	"time"

	"github.com/microsoft/ApplicationInsights-Go/appinsights"
)

// Availability ...
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
// to appinsights. It returns error but the track method used does not have a return and there is no reason to fail on the hostname call.
// So as it exists the return will always be NIL
func (a Availability) SendAvailibiltyStats() error {

	hostname, err := os.Hostname()
	if err != nil {
		hostname = "Unknown"
	}

	availability := appinsights.NewAvailabilityTelemetry(a.Name, a.Time, a.Success)
	availability.RunLocation = "LSPAKS : " + hostname
	availability.Message = a.Msg
	availability.Id = fmt.Sprintf("%v:%v", a.Name, time.Now())
	availability.MarkTime(a.Start, a.End)

	a.Client.Track(availability)

	return nil
}
