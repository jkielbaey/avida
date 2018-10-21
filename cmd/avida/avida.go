package main

import (
	"flag"
	"fmt"
	"os"
	"os/user"

	"github.com/jkielbaey/avida/internal/avida"
	log "github.com/sirupsen/logrus"
)

var logger = log.New()
var version, build string

func main() {

	versionFlag := flag.Bool("version", false, "prints avida version")
	flag.Parse()

	if *versionFlag {
		fmt.Printf("avida version: %s, build: %s\n", version, build)
		os.Exit(0)
	}

	usr, err := user.Current()
	if err != nil {
		logger.Fatal(err)
	}

	// Create the logger file if doesn't exist. Append to it if it already exists.
	logFile := usr.HomeDir + string(os.PathSeparator) + "avida.log"
	file, err := os.OpenFile(logFile, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0600)
	Formatter := new(log.TextFormatter)
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
	logger.SetLevel(log.InfoLevel)

	// Read config file
	conf := avida.GetConfig(logger)

	// Add all fixed positions from the configuration.
	var allPositions []avida.Position
	allPositions = append(allPositions, conf.Positions...)

	// Retrieve all positions on the exchanges.
	for _, exchange := range conf.Exchanges {
		positions, err := exchange.GetPositions()
		if err != nil {
			fmt.Println(err)
		}
		allPositions = append(allPositions, *positions...)
	}

	// Determine the USD value of all positions.
	totalValueUSD := 0.0
	for _, p := range allPositions {
		v, err := p.GetValueUSD()
		if err != nil {
			logger.Error(err)
			fmt.Printf("ERROR: %s\n", err)
		}
		fmt.Printf("%5s : %7.2f => $%7.2f\n", p.Symbol, p.Amount, v)
		totalValueUSD += v
	}
	fmt.Printf("%15s  ------------\n", " ")
	fmt.Printf(" %-14s    $%7.2f\n", "Total ==> ", totalValueUSD)
}
