package main

import (
	"fmt"
	"gitlab.com/jackkdev/phantom-hosting-api/api"
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
	api.Start()
}