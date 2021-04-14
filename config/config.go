package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

type config struct {
	Pagesize int
}

var Conf config

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	Conf.Pagesize = 20
}

func GetName(name string) string {
	return os.Getenv(name)
}
