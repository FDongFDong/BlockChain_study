package main

import (
	"coin/test/utils"
	"crypto/ecdsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/hex"
	"fmt"
	"math/big"
)

const (
	signature     string = "059d7cb78150b004c485803df8833937540361d3f350d26497410ec585998989397bdb4c41dc8efa8ce7de9cbd70486bdc4467b1a783b41c3b080d4afc842c83"
	privateKey    string = "30770201010420d6c0e023d53f9cf210ad59c82a6949a44942af11d0795ee34e243f4674693dc6a00a06082a8648ce3d030107a14403420004e58be678aba80d88b68da73bc468ff85eba2d398c2b932b7b5c63391263cfa6884cd57daf27f48f670541e0e579c53c4ad9b53315bf10426540f1564e245498d"
	hashedMessage string = "1c5863cd55b5a4413fd59f054af57ba3c75c0698b3851d70f99b8de2d5c7338f"
)

func main() {

	priByte, _ := hex.DecodeString(privateKey)

	private, _ := x509.ParseECPrivateKey(priByte)

	sigBytes, _ := hex.DecodeString(signature)

	rBytes := sigBytes[:len(sigBytes)/2]
	sBytes := sigBytes[len(sigBytes)/2:]

	var bigR, bigS = big.Int{}, big.Int{}

	bigR.SetBytes(rBytes)
	bigS.SetBytes(sBytes)

	fmt.Println(bigR, bigS)

	hashBytes, err := hex.DecodeString(hashedMessage)
	utils.HandleErr(err)
	ok := ecdsa.Verify(&private.PublicKey, hashBytes, &bigR, &bigS)
	fmt.Println(ok)

}

func Hash(i interface{}) string {
	s := fmt.Sprintf("%v", i)
	hash := sha256.Sum256([]byte(s))
	return fmt.Sprintf("%x", hash)
}

// 메세지에 서명하는 방법
// 1.메세지를 해시한다.
// 2. (private/public key)키 페어를 만든다.
// 3. 1번에서 만든 해시에 비밀키로 서명을 한다.
//  ("해시된 메세지" + 비밀키 = 서명)

// 서명된 메세지가 내것이람을 증명하는 방법(검증)
// 해시된 메세지 + 서명 + public key -> true/false
