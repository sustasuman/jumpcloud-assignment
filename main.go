package main

import (
	"fmt"

	"github.com/spf13/viper"
)

func loadEnv() {
	viper.SetConfigName("config/config")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %w \n", err))
	}
	//TODO Files can be remote and can be loaded using viper.
}

func main() {
	loadEnv()
	StartServer()
}
