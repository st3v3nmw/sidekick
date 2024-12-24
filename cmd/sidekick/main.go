package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"

	"github.com/st3v3nmw/sidekick/internal/engine"
)

func main() {
	fix := flag.Bool("fix", false, "fix the failing command")
	flag.Parse()

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

	var request string
	if *fix {
		request = "fix the failing command:"
		content, err := getPaneContent()
		if err != nil {
			log.Fatal(err)
		}

		request += "\n" + content
	} else {
		request = strings.Join(os.Args[1:], " ")
	}

	e.Loop(request)
}

func mustGetEnv(envVar string) (string, error) {
	fullEnvVar := fmt.Sprintf("SK_%s", envVar)
	value, ok := os.LookupEnv(fullEnvVar)
	if !ok {
		return "", fmt.Errorf("env var not set: %s", fullEnvVar)
	}

	return value, nil
}

func getPaneContent() (string, error) {
	cmd := exec.Command("tmux", "capture-pane", "-p", "-S", "-1")
	out, err := cmd.Output()
	if err != nil {
		return "", err
	}

	return strings.TrimSpace(string(out)), nil
}
