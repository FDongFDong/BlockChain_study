# BlockChain_study
Go 언어로 블록체인 스터디
## Genesis Block 만들어보기
[exam01](https://github.com/FDongFDong/BlockChain_study/tree/main/exam01)
 
## Genesis Block, Second Block, ... 만들어보기
[exam02]()

## Refactoring, Singleton 
[exam03]()
[exam04]()

- 기존 함수에 문제가 있다.
  - 하나의 함수에서 블록을 생성하고
  - 블록을 해시하고
  - 새로운 블록을 추가한다.
  
```go
  func (b *blockchain) addBlock(data string) {
  newBlock := block{data, "", b.getLastHash()}
  hash := sha256.Sum256([]byte(newBlock.data + newBlock.prevHash))
  newBlock.hash = fmt.Sprintf("%x", hash)
  b.blocks = append(b.blocks, newBlock)
  }
```
