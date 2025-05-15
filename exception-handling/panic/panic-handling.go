package main

import (
	"fmt"
	"os"
)

func main() {
	configData := loadConfig("config.json")

	fmt.Println("config loaded successfully", string(configData))
}

func loadConfig(filename string) []byte {

	data, err := os.ReadFile(filename)

	if err != nil {
		panic(fmt.Sprintf("failed to load config %v", err))
	}
	return data
}
