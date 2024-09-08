// scripts/generate_config.go
package main

import (
	"log"
	"os"
	"text/template"

	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	tmpl, err := template.ParseFiles("mediamtx.yml.template")
	if err != nil {
		panic(err)
	}

	data := map[string]string{
		"RTSP_URL":     os.Getenv("RTSP_URL"),
		"KEYCLOAK_URL": os.Getenv("KEYCLOAK_URL"),
		"SSL_URL":      os.Getenv("SSL_URL"),
	}

	f, err := os.Create("mediamtx.yml")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	err = tmpl.Execute(f, data)
	if err != nil {
		panic(err)
	}
}
