package main

import (
	"log"
	"os"
)

func getEnv(key string) string {
	value := os.Getenv(key)
	if value == "" {
		log.Fatalf("Variabile di ambiente %s non trovata", key)
	}
	return value
}
