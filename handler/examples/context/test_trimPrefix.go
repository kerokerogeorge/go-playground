package handler

import (
	"fmt"
	"math/big"
	"strings"
)

func TrimPrefix() {
	address := "0xfffff"
	hexString := strings.TrimPrefix(address, "0x")
	fmt.Printf("result: %s", hexString)
	a := fmt.Sprintf("%064s", hexString)
	fmt.Println("result:", a)

	i, _ := new(big.Int).SetString("10000", 10)
	fmt.Println("result:", i)
}
