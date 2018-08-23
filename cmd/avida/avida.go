package main

import (
	"fmt"
	"os"
	"os/user"

	"github.com/Sirupsen/logrus"
	"github.com/jkielbaey/avida/internal/avida"
)

var logger = logrus.New()

func main() {

	usr, err := user.Current()
	if err != nil {
		logger.Fatal(err)
	}

	// Create the logger file if doesn't exist. Append to it if it already exists.
	logFile := usr.HomeDir + string(os.PathSeparator) + "avida.log"
	file, err := os.OpenFile(logFile, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0600)
	Formatter := new(logrus.TextFormatter)
	Formatter.TimestampFormat = "02-01-2006 15:04:05"
	Formatter.FullTimestamp = true
	logger.Formatter = Formatter
	if err == nil {
		logger.Out = file
	} else {
		logger.Info("Failed to log to file, using default stderr")
	}
	defer file.Close()

	logger.Info("------------------ Avida Dollars. Let's count your money!!  ------------------")
	logger.SetLevel(logrus.InfoLevel)

	// Read config file
	conf := avida.GetConfig(logger)
	fmt.Println(conf)

}
