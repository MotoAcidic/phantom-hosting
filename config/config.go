package config

import (
	"fmt"
	"github.com/icrowley/fake"
	uuid "github.com/satori/go.uuid"
	"time"
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

func GenerateNodeDetails(m MasternodeString) (mnString string) {
	m.EpochTime = time.Now().Unix()
	m.IPv6 = "[" + fake.IPv6() + "]:"
	m.Alias = uuid.NewV4().String()

	mnString = fmt.Sprintf("%s %s%d %s %s %d %d", m.Alias, m.IPv6, m.Port, m.Genkey, m.TransactionID, m.TransactionIndex, m.EpochTime)

	return mnString
}