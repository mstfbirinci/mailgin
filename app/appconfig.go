package app

import (
	"encoding/json"
	"log"
	"os"
)

var instance Configuration

//ConfigurationInstance returns an configuration object instance
//If configuration object already initialize, it returns same instance
func ConfigurationInstance() *Configuration {

	if instance != (Configuration{}) {
		return &instance
	}

	file, openError := os.Open("app/config.json")

	if openError != nil {
		log.Fatal("Error occured while configuration file is reading. Exception: ", openError)
	}

	decoder := json.NewDecoder(file)

	decodeError := decoder.Decode(&instance)

	if decodeError != nil {
		log.Fatal("Error occured while configuration file is binding to config object. Exception: ", openError)
	}

	return &instance

}
