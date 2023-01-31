package blockchain

import (
	"coin/exam43/db"
	"coin/exam43/utils"
	"errors"
	"fmt"
	"strings"
	"time"
)

type Block struct {
	Hash         string `json:"hash"`
	PrevHash     string `json:"prevhash,omitempty"`
	Height       int    `json:"height"`
	Difficulty   int    `json:"defficulty"`
	Nonce        int    `json:"nonce"`
	Timestamp    int    `json:"timestamp"`
	Transactions []*Tx  `json:"transactions"`
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
		// time.Now().Unix() : int64를 반환한다. 1970년 1월 1일 UTC로부터 흐른 시간을 초단위로
		b.Timestamp = int(time.Now().Unix())
		hash := utils.Hash(b)
		fmt.Printf("\n\n\nTarget:%s\nHash:%s\nNonce:%d\n\n\n", target, hash, b.Nonce)

		if strings.HasPrefix(hash, target) {
			b.Hash = hash
			break
		} else {
			b.Nonce++
		}

	}
}

func createBlock(prevHash string, height int, diff int) *Block {
	block := &Block{

		Hash:       "",
		PrevHash:   prevHash,
		Height:     height,
		Difficulty: diff,
		Nonce:      0,
	}
	block.mine()
	block.Transactions = Mempool.TxToConfirm()
	block.persist()
	return block
}
