package blockchain

import (
	"coin/exam32/db"
	"coin/exam32/utils"
	"sync"
)

const (
	defaultDifficulty  int = 2
	difficultyInterval int = 5
)

type blockchain struct {
	// 최근에 등록된 Hash
	NewestHash string `json:"newestHash"`
	// 블록의 수
	Height            int `json:"height"`
	CurrentDifficulty int `json:"currentdifficulty"`
}

var b *blockchain

var once sync.Once

func (b *blockchain) restore(data []byte) {
	utils.FromBytes(b, data)

}

func (b *blockchain) persist() {
	db.SaveBlockchain(utils.ToBytes(b))
}
func (b *blockchain) Blocks() []*Block {
	var blocks []*Block
	hashCursor := b.NewestHash
	for {
		block, _ := FindBlock(hashCursor)
		blocks = append(blocks, block)
		if block.PrevHash != "" {
			hashCursor = block.PrevHash
		} else {
			break
		}
	}
	return blocks
}

func (b *blockchain) AddBlock(data string) {
	block := createBlock(data, b.NewestHash, b.Height+1)
	b.NewestHash = block.Hash
	b.Height = block.Height
	b.persist()
}

func (b *blockchain) difficulty() int {
	// 제네시스 블록체인은 Difficulty가 2다.
	if b.Height == 0 {
		return defaultDifficulty
	} else if b.Height%difficultyInterval == 0 {
		// 비트코인은 2016개, 우리의 블록체인은 5개 블록마다 체크하여 난이도를 조정한다

	} else {
		// 난이도가 변경된 후 블록이 5개가 추가되지 않았으면 현재 난이도를 그대로 유지한다.
		return b.CurrentDifficulty
	}
}
func Blockchain() *blockchain {
	if b == nil {
		once.Do(func() {

			b = &blockchain{
				Height: 0,
			}

			checkpoint := db.Checkpoint()
			if checkpoint == nil {
				//없으면 initialize한다.
				b.AddBlock("Genesis Block")
			} else {

				b.restore(checkpoint)
			}

		})
	}

	return b
}
