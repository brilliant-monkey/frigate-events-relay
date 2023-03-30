package config

import (
	"errors"
	"fmt"
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

func Load(configPathEnv string, out interface{}) (err error) {
	log.Println("Loading config...")

	path := os.Getenv(configPathEnv)
	bytes, err := os.ReadFile(path)
	if err != nil {
		message := fmt.Sprintf("%s environment variable is not set.", configPathEnv)
		err = errors.New(message)
		return
	}

	err = yaml.Unmarshal(bytes, out)
	if err != nil {
		return
	}

	log.Println("Config loaded successfully.")

	return
}
