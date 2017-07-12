package utility

import (
	"encoding/json"
	"os"
)

// For Externalized Configuration
var Configuration EntConfiguration

type EntConfiguration struct {
	Secret string `json:"secret"`
	Port   string `json:"port"`
}

// Reading properti from ./env.json for externalizing application configuration
// When error, expecting application to stop (panic)
func InitializeConfig() {
	file, err := os.Open("./env.json")
	if err != nil {
		panic(err)
	}
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&Configuration)
	if err != nil {
		panic(err)
	}
}

// Another utility function needed by the application
func ContainsInt(intSlice []int, searchInt int) bool {
	for _, value := range intSlice {
		if value == searchInt {
			return true
		}
	}
	return false
}
