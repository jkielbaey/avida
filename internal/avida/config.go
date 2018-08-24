package avida

import (
	"os"
	"os/user"

	"github.com/BurntSushi/toml"
	log "github.com/sirupsen/logrus"
)

// Config represents the avida config file.
type Config struct {
	LogLevel  string
	Exchanges []Exchange
	Positions []Position
}

// GetConfig reads the avida config file.
func GetConfig(logger *log.Logger) Config {
	usr, err := user.Current()
	if err != nil {
		logger.Fatal(err)
	}

	var conf Config
	configFile := usr.HomeDir + string(os.PathSeparator) + ".avida.toml"
	if _, err := toml.DecodeFile(configFile, &conf); err != nil {
		logger.Fatal(err.Error())
	}

	if conf.LogLevel == "" {
		conf.LogLevel = "info"
	}

	logger.WithFields(log.Fields{
		"LogLevel":   conf.LogLevel,
		"#Exchanges": len(conf.Exchanges),
		"#Positions": len(conf.Positions),
	}).Info("Config settings")

	return conf
}
