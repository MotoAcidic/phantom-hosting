package main

import (
	"fmt"
	"gitlab.com/jackkdev/phantom-hosting-api/config"
	"os"
)

func createConfig(path string) {
	_, err := os.Stat(path)
	if os.IsExist(err) {
		file, err := os.Create(path)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer file.Close()
	}
	fmt.Println("Config file created successfully")
}

func main() {
	fmt.Println(config.GenerateNodeDetails(config.MasternodeString{}))
}