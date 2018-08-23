package avida

import (
	"os"
	"os/user"

	"github.com/BurntSushi/toml"
	"github.com/Sirupsen/logrus"
)

// Config represents the avida config file.
type Config struct {
	LogLevel  string
	Exchanges []struct {
		Name      string
		Enabled   bool
		APIKey    string
		APISecret string
	}
	ColdAssets []struct {
		Coin     string
		Location string
		Amount   float32
	}
}

// GetConfig reads the avida config file.
func GetConfig(logger *logrus.Logger) Config {
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

	logger.WithFields(logrus.Fields{
		"LogLevel":    conf.LogLevel,
		"#Exchanges":  len(conf.Exchanges),
		"#ColdAssets": len(conf.ColdAssets),
	}).Info("Config settings")

	return conf
}
