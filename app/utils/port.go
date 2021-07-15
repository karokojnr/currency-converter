package utils

import (
	"fmt"
)

func GetPort() string {
	var port = GoDotEnvVariable("PORT")
	if port == "" {
		port = "4000"
		fmt.Println("INFO: No PORT environment variable detected, defaulting to " + port)
	}
	return ":" + port
}