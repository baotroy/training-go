package cryp

import (
	"crypto/ecdsa"
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

func PrivateKeyToAddress(pk string) common.Address {
	// const privateKey2 = "7ee461c9c2c28a215d3fe103754ff4d56ceccf177602cc9adc1ca0570a389ef0"
	privateKey, err := crypto.HexToECDSA(pk)
	if err != nil {
		fmt.Print(err)
	}
	fmt.Println(privateKey)
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		fmt.Print("error casting public key to ECDSA")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	return fromAddress
}
