package blockchain

import (
	"coin/exam32/db"
	"coin/exam32/utils"
	"sync"
)

type blockchain struct {
	// 최근에 등록된 Hash
	NewestHash string `json:"newestHash"`
	// 블록의 수
	Height int `json:"height"`
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
	// 찾은 블록들을 저장할 Block 포인터를 저장할 슬라이스를 만든다.
	var blocks []*Block
	// 최근에 생성된 블록의 해시값을 가져온다.
	hashCursor := b.NewestHash
	for {
		// 최근 블록 부터 가져온다
		// 무조건 찾을 수 있기 때문에 error 처리를 따로 하지 않는다.
		block, _ := FindBlock(hashCursor)
		// 찾은 블록을 []*block에 넣는다.
		blocks = append(blocks, block)
		// 가져온 블록의 이전 해시값이 빈값이 아니면 이전 블록이 있는 것이므로
		// 이전 블록의 해시 값을 가져와 가르킨다.
		if block.PrevHash != "" {
			hashCursor = block.PrevHash
		} else {
			// PrevHash가 없는 Genesis Block면
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

	return b
}
