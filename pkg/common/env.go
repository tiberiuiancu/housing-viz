package common

import "github.com/joho/godotenv"

func LoadEnv() {
	if err := godotenv.Load(".env"); err != nil {
		panic(err)
	}
}
