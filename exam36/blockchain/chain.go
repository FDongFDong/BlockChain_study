package blockchain

import (
	"coin/exam36/db"
	"coin/exam36/utils"
	"sync"
)

const (
	defaultDifficulty  int = 2
	difficultyInterval int = 5
	blockInterval      int = 2
	allowedRange       int = 2
)

type blockchain struct {
	// 최근에 등록된 Hash
	NewestHash string `json:"newestHash"`
	// 블록의 수
	Height int `json:"height"`
	// 현재 난이도
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

func (b *blockchain) AddBlock() {
	block := createBlock(b.NewestHash, b.Height+1)
	b.NewestHash = block.Hash
	b.Height = block.Height
	b.CurrentDifficulty = block.Difficulty
	b.persist()
}

func (b *blockchain) recalculateDifficulty() int {
	allBlocks := b.Blocks()
	// 블록 슬라이스에 가장 최근 블록이 앞에 들어간다.
	newestBlock := allBlocks[0]
	// 가장 최근에 난이도가 재설정된 블록은 allBlock[5-1]이다.
	lastRecalculatedBlock := allBlocks[difficultyInterval-1]
	// 두 블록 사이에 걸린시간 : 최근 생성된 블록의 시간 - 블록의 난이도가 재설정된 후 생성된 블록의 시간
	// 타임스탬프가 Unix Time이기에 초단위로 변경해줘야한다.
	actualTime := (newestBlock.Timestamp / 60) - (lastRecalculatedBlock.Timestamp / 60)
	// 예상 시간 : 5 * 2 = 10분 기준으로 난이도를 설정한다.
	expectedTime := difficultyInterval * blockInterval

	// 10분을 기준으로 앞뒤로 2분씩 범위안에만 들어오면 난이도를 유지한다.
	// 예상 시간 보다 빠르게 블록이 생성되면 난이도를 1만큼 증가시킨다.
	if actualTime <= (expectedTime - allowedRange) {
		return b.CurrentDifficulty + 1
	} else if actualTime >= (expectedTime + allowedRange) {
		// 예상 시간 보다 느리게 블록이 생성되면 난이도를 1만큼 감소시킨다.
		return b.CurrentDifficulty - 1
	}
	return b.CurrentDifficulty
}

func (b *blockchain) difficulty() int {
	// 제네시스 블록체인은 Difficulty가 2다.
	if b.Height == 0 {
		return defaultDifficulty
	} else if b.Height%difficultyInterval == 0 {
		// 비트코인은 2016개, 우리의 블록체인은 5개 블록마다 체크하여 난이도를 조정한다
		return b.recalculateDifficulty()
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
				b.AddBlock()
			} else {
				b.restore(checkpoint)
			}

		})
	}

	return b
}
