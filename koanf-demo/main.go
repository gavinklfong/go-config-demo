package main

import (
	"fmt"
	"log"

	"github.com/knadh/koanf/parsers/yaml"
	"github.com/knadh/koanf/providers/file"
	"github.com/knadh/koanf/v2"
)

// Global koanf instance. Use "." as the key path delimiter
var k = koanf.New(".")

func main() {
	// Load YAML config.
	if err := k.Load(file.Provider("../config/application.yaml"), yaml.Parser()); err != nil {
		log.Fatalf("error loading config: %v", err)
	}

	fmt.Printf("server.port=%v\n", k.Int("server.port"))
}
