package config

import (
	"fmt"
	"github.com/spf13/viper"
)

func init() {

	viper.SetConfigFile(".env")
	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		panic(fmt.Errorf("fatal error config file: %w", err))
	}
	viper.SetDefault("page_size", 20)
	viper.Set("Taxonomies", map[string]string{"tag": "tags", "category": "categories"})
}

func Get(name string) interface{} {
	return viper.Get(name)
}
