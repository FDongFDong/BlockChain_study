package blockchain

import (
	"coin/exam32/db"
	"coin/exam32/utils"
	"crypto/sha256"
	"errors"
	"fmt"
	"strings"
)

// 임시로 hash값 중 0이 2개로 시작하는 hash를 찾도록 난이도설정
const difficulty int = 2

type Block struct {
	Data       string `json:"data"`
	Hash       string `json:"hash"`
	PrevHash   string `json:"prevhash,omitempty"`
	Height     int    `json:"height"`
	Difficulty int    `json:"defficulty"`
	Nonce      int    `json:"nonce"`
}

func (b *Block) persist() {
	db.SaveBlock(b.Hash, utils.ToBytes(b))
}

var ErrNotFound = errors.New("block not found")

func (b *Block) restore(data []byte) {
	utils.FromBytes(b, data)
}

func FindBlock(hash string) (*Block, error) {
	blockBytes := db.Block(hash)
	if blockBytes == nil {
		return nil, ErrNotFound
	}
	block := &Block{}
	block.restore(blockBytes)
	return block, nil
}
func (b *Block) mine() {
	target := strings.Repeat("0", b.Difficulty)
	for {
		blockAsString := fmt.Sprint(b)
		// 블록을 string으로 바꾼 후 해쉬로 변환시킨다음 16진수 string으로 다시 변환한다.
		hash := fmt.Sprintf("%x", sha256.Sum256([]byte(blockAsString)))
		fmt.Printf("Block as String:%s\nHash:%s\nTarget:%s\nNonce:%d\n\n\n", blockAsString, hash, target, b.Nonce)
		if strings.HasPrefix(hash, target) {
			b.Hash = hash
			break
		} else {
			b.Nonce++
		}

	}
}

func createBlock(data string, prevHash string, height int) *Block {
	block := &Block{
		Data:       data,
		Hash:       "",
		PrevHash:   prevHash,
		Height:     height,
		Difficulty: difficulty,
		Nonce:      0,
	}
	block.mine()
	block.persist()
	return block
}
