package main

import (
	"fmt"
	"strings"

	"github.com/spf13/viper"
)

func main() {
	mainViper := getMainConfig()
	// specificViper := getEnvSpecificConfig()

	fmt.Println(mainViper.GetInt("server.port"))
	// fmt.Println(specificViper.GetString("db.url"))

	// getEnvVariableConfig()
}

func getEnvVariableConfig() *viper.Viper {
	envVarViper := viper.New()
	envVarViper.SetEnvPrefix("SRV")
	// err := envVarViper.BindEnv("DB_USER")
	// if err != nil { // Handle errors reading the config file
	// 	panic(fmt.Errorf("fatal error config: %w", err))
	// }

	// fmt.Println(envVarViper.GetString("DB_USER"))

	envVarViper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	envVarViper.AutomaticEnv()

	fmt.Println(envVarViper.GetString("db.user"))

	return envVarViper
}

func getMainConfig() *viper.Viper {
	mainViper := viper.New()

	mainViper.SetEnvPrefix("SRV")

	fmt.Println(mainViper.GetString("DB_USER"))

	mainViper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	mainViper.AutomaticEnv()

	fmt.Println(mainViper.GetString("DB_USER"))

	mainViper.SetConfigName("application")           // name of config file (without extension)
	mainViper.SetConfigType("yaml")                  // REQUIRED if the config file does not have the extension in the name
	mainViper.AddConfigPath("./config")              // path to look for the config file in
	mainViper.AddConfigPath("/etc/app/forex/config") // call multiple times to add many search paths

	err := mainViper.ReadInConfig() // Find and read the config file
	if err != nil {                 // Handle errors reading the config file
		panic(fmt.Errorf("fatal error config file: %w", err))
	}

	// fmt.Println(mainViper.GetInt("server.port"))
	// fmt.Println(mainViper.GetString("app.default-base-currency"))
	// fmt.Println(mainViper.GetString("db.url"))

	config := Config{}
	mainViper.Unmarshal(&config)
	fmt.Printf("%+v\n", config)

	return mainViper
}

func getEnvSpecificConfig() *viper.Viper {
	envSpecificViper := viper.New()

	envSpecificViper.SetConfigName("application-sit")
	envSpecificViper.SetConfigType("yaml")
	envSpecificViper.AddConfigPath("./config")
	envSpecificViper.AddConfigPath("/etc/app/forex/config")

	err := envSpecificViper.ReadInConfig() // Find and read the config file
	if err != nil {                        // Handle errors reading the config file
		panic(fmt.Errorf("fatal error config file: %w", err))
	}

	// fmt.Println(envSpecificViper.GetString("db.url"))

	return envSpecificViper
}
