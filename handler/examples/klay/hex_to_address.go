package klay

import (
	"log"

	"github.com/klaytn/klaytn/common"
)

func HexToAddress() {
	address := common.HexToAddress("0xD38E8c7918c69ebDA0b9277E917421c3843B5E45")
	log.Println("address:", address)
}
