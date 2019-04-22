package main

import (
	"fmt"
	"github.com/icrowley/fake"
	"github.com/satori/go.uuid"
	"os"
	"strconv"
	"time"
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

func getNodeDetails() (mnString string) {
	epochTime := time.Now().Unix()
	ipv6 := "[" + fake.IPv6() + "]:"
	port := "4918"
	uuid := uuid.NewV4().String()

	mnString = uuid + " " + ipv6 + port + " <GENKEY>" + " <TXID>" + " <TXINDEX>" + " " + strconv.FormatInt(epochTime, 10)

	return mnString
}

func main() {
	fmt.Println(getNodeDetails())
}