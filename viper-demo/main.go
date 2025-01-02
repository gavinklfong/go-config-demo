package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/viper"
)

func main() {

	fmt.Println("Set env variable SRV_DB_PASSWORD=NewPassword")
	os.Setenv("SRV_DB_PASSWORD", "NewPassword")

	getMainConfig()
	getEnvSpecificConfig()
	getEnvVariableConfig()
}

func getEnvVariableConfig() *viper.Viper {
	fmt.Println("\n======== Env Variable Config ========")

	envVarViper := viper.New()
	envVarViper.SetEnvPrefix("SRV")
	envVarViper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	envVarViper.AutomaticEnv()

	fmt.Printf("db.password=%v\n", envVarViper.GetString("db.password"))

	return envVarViper
}

func getMainConfig() *viper.Viper {
	fmt.Println("\n======== application.yaml + Env Variables + Unmarshal ========")

	mainViper := viper.New()

	mainViper.SetEnvPrefix("SRV")

	mainViper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	mainViper.AutomaticEnv()
	fmt.Printf("Env[DB_PASSWORD]=%v\n", mainViper.GetString("DB_PASSWORD"))
	fmt.Printf("Env[db.password]=%v\n", mainViper.GetString("db.password"))

	mainViper.SetConfigName("application")           // name of config file (without extension)
	mainViper.SetConfigType("yaml")                  // REQUIRED if the config file does not have the extension in the name
	mainViper.AddConfigPath("../config")             // path to look for the config file in
	mainViper.AddConfigPath("/etc/app/forex/config") // call multiple times to add many search paths

	err := mainViper.ReadInConfig() // Find and read the config file
	if err != nil {                 // Handle errors reading the config file
		panic(fmt.Errorf("fatal error config file: %w", err))
	}

	fmt.Printf("server.port=%v\n", mainViper.GetInt("server.port"))
	fmt.Printf("app.default-base-currency=%v\n", mainViper.GetString("app.default-base-currency"))
	fmt.Printf("db.url=%v\n", mainViper.GetString("db.url"))

	config := Config{}
	mainViper.Unmarshal(&config)
	fmt.Printf("Unmarshal:\n%+v\n", config)

	return mainViper
}

func getEnvSpecificConfig() *viper.Viper {
	fmt.Println("\n======== application-sit.yaml ========")

	envSpecificViper := viper.New()

	envSpecificViper.SetConfigName("application-sit")
	envSpecificViper.SetConfigType("yaml")
	envSpecificViper.AddConfigPath("../config")
	envSpecificViper.AddConfigPath("/etc/app/forex/config")

	err := envSpecificViper.ReadInConfig() // Find and read the config file
	if err != nil {                        // Handle errors reading the config file
		panic(fmt.Errorf("fatal error config file: %w", err))
	}

	fmt.Printf("db.url=%v\n", envSpecificViper.GetString("db.url"))

	return envSpecificViper
}
