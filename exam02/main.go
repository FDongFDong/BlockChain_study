package main

import (
	"crypto/sha256"
	"fmt"
)

type block struct {
	data     string
	hash     string
	prevHash string
}
type blockchain struct {
	blocks []block
}

// 마지막 블록의 해시를 불러오는 함수
func (b *blockchain) getLastHash() string {
	if len(b.blocks) > 0 {
		// 블록체인에 등록된 블록이 0개가 아니라면 새로 생성되는 블록은 제네시스 블록이 아니다.
		// 두번째 블록 부터는 이전 블록의 해시값을 블록에 함께 저장해준다.
		return b.blocks[len(b.blocks)-1].hash
	}
	// 0개라면 제네시스 블록이므로 이전 해시값을 저장하지 않는다.
	return ""
}

func (b *blockchain) addBlock(data string) {
	newBlock := block{data, "", b.getLastHash()}
	hash := sha256.Sum256([]byte(newBlock.data + newBlock.prevHash))
	newBlock.hash = fmt.Sprintf("%x", hash)
	// 블록체인에 앞서 생성한 블록을 넣어준다.
	b.blocks = append(b.blocks, newBlock)
}
func (b *blockchain) listBlocks() {
	for _, block := range b.blocks {
		fmt.Printf("Data : %s\n", block.data)
		fmt.Printf("Hash : %s\n", block.hash)
		fmt.Printf("Prev Hash : %s\n", block.prevHash)
	}
}
func main() {
	chain := blockchain{}
	chain.addBlock("Genesis Block")
	chain.addBlock("Second Block")
	chain.addBlock("Third Block")
	chain.listBlocks()
}
