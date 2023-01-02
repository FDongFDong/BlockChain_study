package blockchain

import (
	"coin/exam38/utils"
	"errors"
	"time"
)

const (
	// 채굴 시 보상으로 주는 코인의 갯수
	minerReward int = 50
)

type mempool struct {
	Txs []*Tx
}

// 비어있는 mempool 생성
var Mempool *mempool = &mempool{}

type Tx struct {
	Id        string   `json:"id"`
	Timestamp int      `json:"timestamp"`
	TxIns     []*TxIn  `json:"txins"`
	TxOuts    []*TxOut `json:"txouts"`
}

func (t *Tx) getId() {
	// 트랜잭션에 아이디 값을 만들어준다.
	t.Id = utils.Hash(t)
}

type TxIn struct {
	Owner  string `json:"owner"`
	Amount int    `json:"amount"`
}

type TxOut struct {
	Owner  string `json:"owner"`
	Amount int    `json:"amount"`
}

func makeCoinbaseTx(address string) *Tx {
	txIns := []*TxIn{
		{"COINBASE", minerReward},
	}
	txOuts := []*TxOut{
		{address, minerReward},
	}
	tx := Tx{
		Id:        "",
		Timestamp: int(time.Now().Unix()),
		TxIns:     txIns,
		TxOuts:    txOuts,
	}
	tx.getId()
	return &tx
}

// 유효한 트랜잭션을 생성하려면, 유저가 input에 돈이 들어있다는걸 보여주면 된다.
// Transaction Input을 가져와서 Transaction Output을 만들면 해당 Transaction은 유효해진다.
func makeTx(from, to string, amount int) (*Tx, error) {
	// from의 잔고가 보내고자 하는 금액 보다 적으면 에러 출력
	if Blockchain().BalanceByAddress(from) < amount {
		return nil, errors.New("not enough money")
	}
	// 새로운 트랜잭션을 만들기 위해 txIns, txOuts을 생성한다.
	var txIns []*TxIn
	var txOuts []*TxOut

	total := 0
	// 이전 TxOut으로 TxInput을 만들기 위해 요청한 사용자의 TxOut List를 가져온다.
	oldTxOuts := Blockchain().TxOutsByAddress(from)
	// total값에 TxOut을 모아서 amount 값이 되도록 한다.
	for _, txOut := range oldTxOuts {
		if total > amount {
			break
		}
		// 보내려는 값이 일치할 때 까지
		txIn := &TxIn{txOut.Owner, txOut.Amount}
		txIns = append(txIns, txIn)
		total += txIn.Amount
	}
	change := total - amount
	// TxOut에게 줄 거스름돈이 있을 경우
	if change != 0 {
		changeTxOut := &TxOut{from, change}
		txOuts = append(txOuts, changeTxOut)
	}
	txOut := &TxOut{to, amount}
	txOuts = append(txOuts, txOut)
	// 새로운 트랜잭션을 만들어준다.
	tx := &Tx{
		Id:        "",
		Timestamp: int(time.Now().Unix()),
		TxIns:     txIns,
		TxOuts:    txOuts,
	}
	tx.getId()
	return tx, nil
}

// 트랜잭션을 Mempool에 추가한다. 트랜잭션을 생성하지는 않는다.
// 어떠한 이유로 트랜잭션을 블록에 추가할 수 없으면 error을 리턴해준다.
func (m *mempool) AddTx(to string, amount int) error {
	tx, err := makeTx("fdongfdong", to, amount)
	if err != nil {
		return err
	}
	// 트랜잭션이 정상적으로 만들어졌다면 Mempool에 추가해준다.
	m.Txs = append(m.Txs, tx)
	return nil
}
