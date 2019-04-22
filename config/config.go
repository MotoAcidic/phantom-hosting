package config

import (
	"fmt"
	"github.com/icrowley/fake"
	uuid "github.com/satori/go.uuid"
	"os"
	"time"
	"errors"
)

type MasternodeString struct {
	Alias string `json:"alias"`
	Genkey string `json:"genkey"`
	IPv6 string `json:"ipv6"`
	Port int `json:"port"`
	TransactionID string `json:"txid"`
	TransactionIndex int `json:"tx_index"`
	EpochTime int64 `json:"epoch_time"`
}

func GenerateNodeDetails(m MasternodeString) (mnString string, err error) {
	m.EpochTime = time.Now().Unix()
	m.IPv6 = "[" + fake.IPv6() + "]:"
	m.Alias = uuid.NewV4().String()

	if m.TransactionID == "" {
		return "", errors.New("Transaction ID is required")
	}
	if m.TransactionIndex == 0 {
		return "", errors.New("Transaction Index is required")
	}
	if m.Port == 0 {
		return "", errors.New("Port is required")
	}
	if m.Genkey == "" {
		return "", errors.New("Masternode Genkey is required")
	}

	mnString = fmt.Sprintf("%s %s%d %s %s %d %d", m.Alias, m.IPv6, m.Port, m.Genkey, m.TransactionID, m.TransactionIndex, m.EpochTime)

	return mnString, nil
}

func GenerateConfigurationFile(path string) {
	file, err := os.Create(path)
	if err != nil {
		fmt.Println(err)
	}

	defer file.Close()
}

func AddMasternodeToConfigFile(path string, strMasternode string) {
	file, err := os.OpenFile(path, os.O_APPEND | os.O_WRONLY, 0600)
	if err != nil {
		fmt.Println(err)
	}

	defer file.Close()

	if _, err := file.Write([]byte(strMasternode + "\n")); err != nil {
		fmt.Println(err)
	}
}