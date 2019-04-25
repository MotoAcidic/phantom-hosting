package config

import (
	"errors"
	"fmt"
	"github.com/icrowley/fake"
	"github.com/gofrs/uuid"
	"os"
	"time"
)

type MasternodeString struct {
	Alias            uuid.UUID `json:"alias"`
	Genkey           string    `json:"genkey"`
	IPv4             string    `json:"ipv6"`
	Port             int       `json:"port"`
	TransactionID    string    `json:"txid"`
	TransactionIndex int       `json:"tx_index"`
	EpochTime        int64     `json:"epoch_time"`
}

func GenerateNodeDetails(m MasternodeString) (mnString string, err error) {
	var alias = uuid.Must(uuid.NewV4())

	m.EpochTime = time.Now().Unix()
	m.IPv4 = fake.IPv4()
	m.Alias = alias 

	if m.TransactionID == "" {
		return "", errors.New("Transaction ID is required")
	}
	if m.TransactionIndex < 0 || m.TransactionIndex > 9 {
		return "", errors.New("Transaction Index is out of range")
	}
	if m.Port == 0 {
		return "", errors.New("Port is required")
	}
	if m.Genkey == "" {
		return "", errors.New("Masternode Genkey is required")
	}

	mnString = fmt.Sprintf("%v %s:%d %s %s %d %d", m.Alias, m.IPv4, m.Port, m.Genkey, m.TransactionID, m.TransactionIndex, m.EpochTime)

	return mnString, nil
}

func GenerateConfigurationFile(path string) (err error) {
	file, err := os.Create(path)
	if err != nil {
		return err
	}

	defer file.Close()

	return nil
}

func AddMasternodeToConfigFile(path string, strMasternode string) (err error) {
	file, err := os.OpenFile(path, os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		return err
	}

	defer file.Close()

	if _, err := file.Write([]byte(strMasternode + "\n")); err != nil {
		return err
	}

	return nil
}
