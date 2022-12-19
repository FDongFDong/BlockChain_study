package main

import (
	"crypto/sha256"
	"fmt"
)

type block struct {
	data     string
	hash     string
	prevHash string
}

func main() {
	genesisBlock := block{"Genisis Block", "", ""}
	// hash의 byte값을 가져온다.
	// []byte : string을 byte 슬라이스로 변환해준다.
	// Sum256(): byte slice를 받고 32byte Hash를 반환한다.
	hash := sha256.Sum256([]byte(genesisBlock.data + genesisBlock.prevHash))
	fmt.Println(hash)
	//[214 90 147 214 18 203 25 150 137 60 38 253 83 217 147 203 170 241 232 254 24 20 53 154 51 150 223 9 92 207 42 12]
	// Bitcoin, Ethereum은 16진수(base16) hash로 되어있다.
	// base16로 포맷된 string값으로 바꿔준다.
	hexHash := fmt.Sprintf("%x", hash)
	// 블록에 저장
	genesisBlock.hash = hexHash

	// 두번째 블록...
	// genesisBlock := block{"Genisis Block", "", genesisBlock.hash}
	// ...
	// ...
}
