package blockchain

import (
	"crypto/sha256"
	"errors"
	"fmt"
	"sync"
)

// main.go에서 가져다 쓰기위해 임시로 대문자로 변경
type Block struct {
	Data     string `json:"data"`
	Hash     string `json:"hash"`
	PrevHash string `json:"prevhash,omitempty"`
	Height   int    `json:"height"`
}
type blockchain struct {
	blocks []*Block
}

var b *blockchain

var once sync.Once

func (b *Block) calculateHash() {
	hash := sha256.Sum256([]byte(b.Data + b.PrevHash))
	b.Hash = fmt.Sprintf("%x", hash)
}

func getLastHash() string {
	totalBlocks := len(GetBlockchain().blocks)
	if totalBlocks == 0 {
		return ""
	}
	return GetBlockchain().blocks[totalBlocks-1].Hash
}

func createBlock(data string) *Block {
	newBlock := Block{data, "", getLastHash(), len(GetBlockchain().blocks) + 1}
	newBlock.calculateHash()
	return &newBlock
}

// export 함수
func (b *blockchain) AddBlock(data string) {
	fmt.Println("AddBlock = >", b.blocks)
	b.blocks = append(b.blocks, createBlock(data))

}

func GetBlockchain() *blockchain {
	if b == nil {
		once.Do(func() {
			b = &blockchain{}
			b.AddBlock("Genesis Block")
		})
	}
	return b
}

// 사용자에게 field를 드러내주는 function(singleton의 철학)
func (b *blockchain) AllBlocks() []*Block {
	fmt.Println("AllBlocks = >", b.blocks)
	return b.blocks

}

var ErrNotFound = errors.New("block not found")

func (b *blockchain) GetBlock(height int) (*Block, error) {
	if height > len(b.blocks) {
		return nil, ErrNotFound
	}
	return b.blocks[height-1], nil
}
