package main

import (
	"github.com/spf13/viper"
)

func main() {
	viper.SetConfigFile(".env")
	viper.ReadInConfig()
	apiKey := viper.GetString("API_KEY")
	if apiKey == ""{
		panic("Missing API KEY")
	}

	ctx := 
}
