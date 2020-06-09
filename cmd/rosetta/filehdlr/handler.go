package filehdlr

import (
	"log"
	"os"
	"os/user"

	"github.com/human-caching/rosetta/constants"
)

// GetConfigPath Determines and returns current file path to Rosetta configuration file
func GetConfigPath() string {
	usr, err := user.Current()

	if err != nil {
		log.Fatal("Cannot determine current OS, configuration file check failed")
	}

	return usr.HomeDir + constants.PathSlash + constants.ConfigFileName
}

// ConfigFileCheck Checks if config file exist, makes checking whether or not the is initial running
func ConfigFileCheck() bool {
	currentConfigPath := GetConfigPath()

	_, fileErr := os.Stat(currentConfigPath)

	if os.IsNotExist(fileErr) {
		return false
	}

	return true
}

// CreateConfigurationFile is a Handler function for creating configuration file
func CreateConfigurationFile() bool {
	currentConfigPath := GetConfigPath()

	_, fileErr := os.Stat(currentConfigPath)

	if os.IsExist(fileErr) {
		return false
	}

	f, err := os.Create(currentConfigPath)

	defer f.Close()

	if err != nil {
		return false
	}

	return true
}
