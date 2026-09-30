package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

func Load() {
	err := godotenv.Load() 

	if err != nil {
		fmt.Println("Error in loading .env file")
	}
}

func GetString(key string, fallback string) string {
	// load() // if dotenv file is not already loaded then it load .env file

	value, ok := os.LookupEnv(key)

	if !ok {
		return fallback
	}

	return value
}

func getInt(key string, fallback int) int { // this is for getting the int type value
	// load()

	value, ok := os.LookupEnv(key)

	if !ok {
		return fallback
	}

	intValue, err := strconv.Atoi(value) // manual conversion string to number using go standard library strconv and atoi(ascii to interger)

	if err != nil {
		fmt.Printf("Error in coverting %s to int %v \n", key, err)
		return fallback
	}

	return intValue
}

func getBool(key string, fallback bool) bool {

	// load()

	value, ok := os.LookupEnv(key)

	if !ok {
		return fallback
	}

	valueBool, err := strconv.ParseBool(value)

	if err != nil {
		fmt.Printf("Error in converting %s to boolean %v \n", key, err)
		return fallback
	}

	return  valueBool
}