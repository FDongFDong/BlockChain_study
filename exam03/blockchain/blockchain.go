package blockchain

import "sync"

type block struct {
	data     string
	hash     string
	prevHash string
}
type blockchain struct {
	blocks []block
}

// 블록체인을 공유하고 초기화하는 작업
// Singleton pattern
// blockchain의 단 하나의 instance만을 공유하는 방법

// 공유하지 않으므로 private 선언
var b *blockchain

// 위 변수의 instance를 대신해서 드러내는 function 생성
// -> 다른 패키지에서 우리의 blockchain이 어떤식으로 생성될지 제어 가능하다는 의미
func GetBlockchain() *blockchain {
	// b가 초기화되었는지 확인 후 처음이자 마지막으로 초기화 진행
	// nil이면 아직 생성되지 않아 새로 생성한다는 의미
	// 두번째로 부르는 순간부터 생성, 초기화는 진행되지 않는다.
	if b == nil {
		// "=" 를 통해 b를 새로 생성하지 않고 대입해준다.
		b = &blockchain{}
	}
	return b
}
