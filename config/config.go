package config

import (
	"libro-electronico/helper"
	"log"
	"os"
)
var IPPort, Net = helper.GetAddress()

func getEnv(key, fallback string) string {
    value, exists := os.LookupEnv(key)
    if !exists {
        log.Printf("Warning: Environment variable %s not set, using default value", key)
        return fallback
    }
    return value
}
