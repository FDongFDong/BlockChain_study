package wallet

import (
	"coin/exam40/utils"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/x509"
	"os"
)

type wallet struct {
	privateKey *ecdsa.PrivateKey
}

var w *wallet

const (
	fileName string = "fdongfdong.wallet"
)

func hasWalletFile() bool {
	_, err := os.Stat(fileName)
	return !os.IsNotExist(err)
}

func createPrivKey() *ecdsa.PrivateKey {
	// 키페어를 생성해준다.
	privKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	utils.HandleErr(err)
	return privKey
}

// key를 받아서 복사-붙여넣기 가능한 형태로 변환 후 그 byte들을 파일로 저장한다.
func persistKey(key *ecdsa.PrivateKey) {
	bytes, err := x509.MarshalECPrivateKey(key)
	utils.HandleErr(err)
	// 0644 : 읽기와 쓰기 허용
	err = os.WriteFile(fileName, bytes, 0644)
	utils.HandleErr(err)
}

func Wallet() *wallet {
	if w == nil {
		w = &wallet{}
		// 사용자가 이미 지갑을 가지고 있는지 확인한다.
		if hasWalletFile() {

			// 만약 있다면 그 지갑을 파일로부터 복구한다.
		} else {
			// 만약 없다면..
			// 키페어를 만들어서
			key := createPrivKey()
			// 저장해둔다.
			persistKey(key)
			// 생서한 키페어를 등록한다.
			w.privateKey = key
		}
	}
	return w
}
