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
// Updated: Time-stamp: <2018-07-09 15:59:25>
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

	log.SetLevel(log.InfoLevel)
	// If changing log level to Debug, we will see two log message
	// log.SetLevel(log.DebugLevel)
}

func main() {
	log_init()
	contextLogger := log.WithFields(log.Fields{"app": "test"})
	contextLogger.Info("Info msg")
	contextLogger.Debug("warn msg")
}
// File: log_basic.go ends
