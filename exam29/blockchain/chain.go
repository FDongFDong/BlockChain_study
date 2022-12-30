package blockchain

import (
	"coin/exam29/db"
	"coin/exam29/utils"
	"fmt"
	"sync"
)

type blockchain struct {
	NewestHash string `json:"newestHash"`
	Height     int    `json:"height"`
}

var b *blockchain

var once sync.Once

func (b *blockchain) restore(data []byte) {
	utils.FromBytes(b, data)

}

func (b *blockchain) persist() {
	db.SaveBlockchain(utils.ToBytes(b))
}

// AddBlock receiver
func (b *blockchain) AddBlock(data string) {
	block := createBlock(data, b.NewestHash, b.Height+1)
	b.NewestHash = block.Hash
	b.Height = block.Height
	b.persist()
}

func Blockchain() *blockchain {
	if b == nil {
		once.Do(func() {
			// NewestHash가 없고 Height는 0인 블록체인을 만들고
			b = &blockchain{"", 0}
			// checkpoint에 data가 있는지 확인한다.

			checkpoint := db.Checkpoint()
			if checkpoint == nil {
				//없으면 initialize한다.
				b.AddBlock("Genesis Block")
			} else {

				b.restore(checkpoint)
			}

		})
	}

	fmt.Println(b.NewestHash)
	return b
}
