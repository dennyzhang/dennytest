//-------------------------------------------------------------------
// @copyright 2017 DennyZhang.com
// Licensed under MIT
//   https://www.dennyzhang.com/wp-content/mit_license.txt
//
// File: log_basic.go
// Author : Denny <https://www.dennyzhang.com/contact>
// Description : go run log_basic.go
// --
// Created : <2018-04-07>
// Updated: Time-stamp: <2018-07-09 16:20:02>
//-------------------------------------------------------------------
package main

import (
	"fmt"
	"os"
	log "github.com/sirupsen/logrus"
)

// https://github.com/sirupsen/logrus
func log_init() {
	log.SetOutput(os.Stdout)
	file, err := os.OpenFile("/tmp/output.log", os.O_CREATE|os.O_WRONLY, 0666)
	if err == nil {
		log.SetOutput(file)
	} else {
		log.Info("Failed to log to file, using default stderr")
	}

	log.SetLevel(log.InfoLevel)
	// If changing log level to Debug, we will see two log message
	// log.SetLevel(log.DebugLevel)
}

func main() {
	log_init()
	fmt.Println("Check log in /tmp/output.log")
	contextLogger := log.WithFields(log.Fields{"app": "test"})
	contextLogger.Info("Info msg")
	contextLogger.Debug("Warn msg")
}
// File: log_basic.go ends
