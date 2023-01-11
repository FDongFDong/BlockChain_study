package wallet

import (
	"coin/exam40/utils"
	"crypto/ecdsa"
	"crypto/x509"
	"encoding/hex"
	"fmt"
	"math/big"
)

const (
	signature string = "46a31c89268b7475adce02358e0cea60749d2a3f4765f75a2fd0b61c8da5ebdc4a3f29a9f37ae574a8b87966818a7b04522b26e746617f92355aff176365d125"
	// "Hello Golang"의 Hash된 값
	hashedMessage string = "8d2caf9f544c5641e94f35c7dc32ebd5d70bd4c92084c5b6644b017df45406f6"
	privateKey    string = "3077020101042047289063d458e4d0b567f4c4386f98c98c6c20ee168f92f7621518b6c95755f5a00a06082a8648ce3d030107a14403420004b20a694c13a75f0d222eb69473cc64f1c5087dedfbc58ad8e15f1066e78ed6a6926c7760ee56f80c58dcd9b5d0c8d9a0a8bc55d3a206f873ecfd68608d725bac"
)

func Start() {

	privBytes, err := hex.DecodeString(privateKey)
	utils.HandleErr(err)

	private, err := x509.ParseECPrivateKey(privBytes)
	utils.HandleErr(err)

	sigBytes, err := hex.DecodeString(signature)
	utils.HandleErr(err)

	rBytes := sigBytes[:len(sigBytes)/2]
	sBytes := sigBytes[len(sigBytes)/2:]

	var bigR, bigS = big.Int{}, big.Int{}

	bigR.SetBytes(rBytes)
	bigS.SetBytes(sBytes)

	hashBytes, err := hex.DecodeString(hashedMessage)
	utils.HandleErr(err)

	ok := ecdsa.Verify(&private.PublicKey, hashBytes, &bigR, &bigS)
	fmt.Println(ok)

}
