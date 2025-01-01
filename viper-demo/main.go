package main

import (
	"fmt"
	"strings"

	"github.com/spf13/viper"
)

func main() {
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv()

	fmt.Println(viper.GetString("ENV"))

	viper.SetConfigName("application")           // name of config file (without extension)
	viper.SetConfigType("yaml")                  // REQUIRED if the config file does not have the extension in the name
	viper.AddConfigPath("./config")              // path to look for the config file in
	viper.AddConfigPath("/etc/app/forex/config") // call multiple times to add many search paths

	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		panic(fmt.Errorf("fatal error config file: %w", err))
	}

	fmt.Println(viper.GetInt("server.port"))
	fmt.Println(viper.GetString("app.forex-rate-api-url"))

	config := Config{}
	viper.Unmarshal(&config)
	fmt.Printf("%+v\n", config)
}
