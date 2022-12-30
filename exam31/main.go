package main

import (
	"crypto/sha256"
	"fmt"
	"strings"
)

func main() {
	difficulty := 4
	// target := "0" * 2
	// 첫번째 인자값을 두번째 인자값 만큼 연결해서 출력해준다.
	target := strings.Repeat("0", difficulty)
	nonce := 1
	for {
		// 16진수 string으로 변환
		hash := fmt.Sprintf("%x", sha256.Sum256([]byte("hello"+fmt.Sprint(nonce))))
		fmt.Printf("Hash:%s\nTarget:%s\nNonce:%d\n\n", hash, target, nonce)
		if strings.HasPrefix(hash, target) {
			return
		} else {
			nonce++
		}
	}

}
