package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/knadh/koanf/parsers/yaml"
	"github.com/knadh/koanf/providers/env"
	"github.com/knadh/koanf/providers/file"
	"github.com/knadh/koanf/v2"
)

func main() {
	// obtain a koanf instance. Use "." as the key path delimiter. This can be "/" or any character.
	k := koanf.New(".")

	// Load YAML config.
	if err := k.Load(file.Provider("../config/application.yaml"), yaml.Parser()); err != nil {
		log.Fatalf("error loading config: %v", err)
	}

	// Load env specific config
	if err := k.Load(file.Provider("../config/application-sit.yaml"), yaml.Parser()); err != nil {
		log.Fatalf("error loading config: %v", err)
	}

	fmt.Printf("server.port=%v\n", k.Int("server.port"))
	fmt.Printf("app.default-base-currency=%v\n", k.String("app.default-base-currency"))
	fmt.Printf("db.url=%v\n", k.String("db.url"))
	fmt.Printf("db.user=%v\n", k.String("db.user"))

	os.Setenv("SRV_DB_USER", "testuser")

	k.Load(env.Provider("SRV_", ".", func(s string) string {
		return strings.Replace(strings.ToLower(
			strings.TrimPrefix(s, "SRV_")), "_", ".", -1)
	}), nil)

	fmt.Printf("db.user=%v\n", k.String("db.user"))

}
