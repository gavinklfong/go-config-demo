package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"testing"

	"github.com/knadh/koanf/parsers/dotenv"
	"github.com/knadh/koanf/providers/env"
	"github.com/knadh/koanf/providers/file"
	"github.com/knadh/koanf/v2"

	"github.com/stretchr/testify/assert"
)

func TestDotEnv(t *testing.T) {

	fmt.Print("Setup environment variable\n")
	os.Setenv("env", "sit")
	os.Setenv("SRV_DB_PASSWORD", "passme")

	// obtain a koanf instance. Use "." as the key path delimiter. This can be "/" or any character.
	k := koanf.NewWithConf(koanf.Conf{
		Delim: ".",
	})

	fmt.Print("Loading default dotenv\n")
	if err := k.Load(file.Provider("./.env"), dotenv.Parser()); err != nil {
		log.Fatalf("error loading config: %v", err)
	}

	environment := os.Getenv("env")
	if len(strings.TrimSpace(environment)) > 0 {
		fmt.Printf("Loading env specific dotenv - %s\n", environment)
		// Load Dot Env config
		if err := k.Load(file.Provider(fmt.Sprintf("./.env.%s", environment)), dotenv.Parser()); err != nil {
			log.Fatalf("error loading config: %v", err)
		}
	} else {
		fmt.Print("env is not defined, skip environment specific dotenv loading\n")
	}

	fmt.Print("Loading default dotenv\n")
	k.Load(env.Provider("SRV_", ".", func(s string) string {
		return strings.TrimPrefix(s, "SRV_")
	}), nil)

	fmt.Print("Retrieve configuration from Koanf\n")
	fmt.Printf("SERVER_PORT=%v\n", k.Int("SERVER_PORT"))
	fmt.Printf("DB_URL=%v\n", k.String("DB_URL"))
	fmt.Printf("DB_PASSWORD=%v\n", k.String("DB_PASSWORD"))

	assert.Equal(t, 8080, k.Int("SERVER_PORT"))
	assert.Equal(t, "jdbc:mysql://db.sit.example.come:3306/forex", k.String("DB_URL"))
	assert.Equal(t, "passme", k.String("DB_PASSWORD"))
}
