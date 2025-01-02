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

func main2() {
	fmt.Println("\n======== application.yaml + Env Variables + Unmarshal ========")

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

func main() {
	getMainConfig()
}

func getMainConfig() {
	fmt.Println("\n======== application.yaml + Env Variables + Unmarshal ========")

	// obtain a koanf instance. Use "." as the key path delimiter. This can be "/" or any character.
	k := koanf.NewWithConf(koanf.Conf{
		Delim: ".",
	})

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
	fmt.Printf("db.url=%v\n", k.String("db.url"))           // provided by application-sit.yaml
	fmt.Printf("db.password=%v\n", k.String("db.password")) // <placeholder>

	fmt.Printf("setup environment variable SRV_DB_PASSWORD=")
	os.Setenv("SRV_DB_PASSWORD", "NewPassword")

	k.Load(env.Provider("SRV_", ".", func(s string) string {
		return strings.Replace(strings.ToLower(
			strings.TrimPrefix(s, "SRV_")), "_", ".", -1)
	}), nil)

	fmt.Printf("db.password=%v\n", k.String("db.password"))

	config := Config{}
	k.UnmarshalWithConf("", &config, koanf.UnmarshalConf{Tag: "config"})
	fmt.Printf("Unmarshal config:\n%+v\n", config)

	flatConfig := FlatConfig{}
	k.UnmarshalWithConf("", &flatConfig, koanf.UnmarshalConf{Tag: "config", FlatPaths: true})
	fmt.Printf("Unmarshal flat config:\n%+v\n", flatConfig)
}
