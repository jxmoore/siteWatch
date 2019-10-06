package controllers

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/jxmoore/siteWatch/models"
)

// Notify is responsible for notifying when failures exceed the threshold.
func notify(siteName string, count int, notifier *models.Notification) error {

	notifier.Kind = strings.ToLower(notifier.Kind)
	if notifier.Kind == "" || strings.Contains(notifier.Kind, "stdout") {
		fmt.Printf("The test for %s has failed %d times which exceeds the current threshold value.\n", siteName, count)
	} else if notifier.Kind == "log" {
		if notifier.Path == "" {
			notifier.Path = "./siteWatch.log"
		}

		logFile, err := os.OpenFile(notifier.Path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			return errors.New("Error opening log file")
		}

		defer logFile.Close()

		logger := log.New(logFile, "", log.LstdFlags)
		logger.Printf("The test for %s has failed %d times which exceeds the current threshold value.\n", siteName, count)

	} else {
		fmt.Printf("The test for %s has failed %d times which exceeds the current threshold value.\n", siteName, count)
	}

	return nil
}
