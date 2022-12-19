package blockchain

import (
	"crypto/sha256"
	"fmt"
	"sync"
)

type block struct {
	data     string
	hash     string
	prevHash string
}
type blockchain struct {
	// block 타입의 포인터 슬라이스
	blocks []*block
}

var b *blockchain

// 단 한번만 실행되게 하기 위함
var once sync.Once

// Hash를 계산해주는 함수
func (b *block) calculateHash() {
	hash := sha256.Sum256([]byte(b.data + b.prevHash))
	b.hash = fmt.Sprintf("%x", hash)
}

// 이전 해시값을 생성해서 가져오는 함수
func getLastHash() string {
	totalBlocks := len(GetBlockchain().blocks)
	// 블록의 개수가 0개면  마지막 Hash 값이 없기 떄문에 아무것도 반환하지 않는다.
	if totalBlocks == 0 {
		return ""
	}
	// 마지막 블록체인의 Hash를 반환해준다.
	return GetBlockchain().blocks[totalBlocks-1].hash
}

func createBlock(data string) *block {
	newBlock := block{data, "", getLastHash()}
	newBlock.calculateHash()
	return &newBlock
}

func GetBlockchain() *blockchain {
	if b == nil {
		// 단 한번만 실행된다.
		// 병렬적으로 처리하든 뭘하든 해당 함수는 한번만 실행됨을 의미
		once.Do(func() {
			// 블록체인을 하나 생성하고
			b = &blockchain{}
			// 하나의 블록을 추가한다.
			b.blocks = append(b.blocks, createBlock("Genesis Block"))
		})
	}
	return b
}
