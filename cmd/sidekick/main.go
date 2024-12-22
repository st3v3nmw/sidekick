package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/st3v3nmw/sidekick/internal/engine"
)

func main() {
	provider, err := mustGetEnv("PROVIDER")
	if err != nil {
		log.Fatal(err)
	}

	model, err := mustGetEnv("MODEL")
	if err != nil {
		log.Fatal(err)
	}

	apiKey, err := mustGetEnv("API_KEY")
	if err != nil {
		log.Fatal(err)
	}

	e, err := engine.NewEngine(provider, model, apiKey)
	if err != nil {
		log.Fatal(err)
	}

	req := strings.Join(os.Args[1:], " ")
	e.Loop(req)
}

func mustGetEnv(envVar string) (string, error) {
	fullEnvVar := fmt.Sprintf("SK_%s", envVar)
	value, ok := os.LookupEnv(fullEnvVar)
	if !ok {
		return "", fmt.Errorf("env var not set: %s", fullEnvVar)
	}

	return value, nil
}
