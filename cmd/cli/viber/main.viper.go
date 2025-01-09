package main

import (
	"fmt"

	"github.com/spf13/viper"
)

func main() {
	viper := viper.New()
	viper.AddConfigPath("./configs/") // path config
	viper.SetConfigName("local")      // name file
	viper.SetConfigType("yaml")

	// Read configuration
	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("fatal error config file : %w", err))
	}

	// Read server configuration
	fmt.Println("Server port ::", viper.Get("server.port"))
	fmt.Println("Name Database ::", viper.GetString("databases.dbname"))
}
