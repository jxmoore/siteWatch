# siteWatch
A simple application for monitoring sites using HTTP GET requests. 
![RandomText](https://raw.githubusercontent.com/jxmoore/siteWatch/master/img/1.PNG)


This was a simple project to familiarize myself with some basic Go syntax and ideas. 

It takes a JSON file (see ex.json) and monitors the sites within. The _monitoring_ is a simple GET request against the adresses found in the JSON, if the site returns a status code that differs from whats defined in the JSON its deemed as a failure, which will fire an _alert_ (just a printf or log output) when it exceeds the threshold. 

Sitewatch can have this info over HTTP on a port defined in the JSON (as seen above), alert via STDOUT or just drop  the failures in a log file. 

There are a few other odds and ends but again, this was just for learning purposes. 
