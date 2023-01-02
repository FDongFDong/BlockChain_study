package blockchain

import (
	"coin/exam39/utils"
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
	ID        string   `json:"id"`
	Timestamp int      `json:"timestamp"`
	TxIns     []*TxIn  `json:"txins"`
	TxOuts    []*TxOut `json:"txouts"`
}

func (t *Tx) getId() {

	t.ID = utils.Hash(t)
}

type TxIn struct {
	TxID  string `json:"txid"`
	Index int    `json:"index"`
	Owner string `json:"owner"`
}

type TxOut struct {
	Owner  string `json:"owner"`
	Amount int    `json:"amount"`
}

// output이 쓰였는지 안쓰였는지 알 수 있게 해준다.
type UTxOut struct {
	TxID   string
	Index  int
	Amount int
}

func isOnMempool(uTxOut *UTxOut) bool {
	exists := false
Outer:
	for _, tx := range Mempool.Txs {
		for _, input := range tx.TxIns {
			if input.TxID == uTxOut.TxID && input.Index == uTxOut.Index {
				exists = true
				break Outer

			}

		}
	}
	return exists
}

func makeCoinbaseTx(address string) *Tx {
	txIns := []*TxIn{
		{"", -1, "COINBASE"},
	}
	txOuts := []*TxOut{
		{address, minerReward},
	}
	tx := Tx{
		ID:        "",
		Timestamp: int(time.Now().Unix()),
		TxIns:     txIns,
		TxOuts:    txOuts,
	}
	tx.getId()
	return &tx
}

func makeTx(from, to string, amount int) (*Tx, error) {
	if BalanceByAddress(from, Blockchain()) < amount {
		return nil, errors.New("not enoguh ")
	}
	var txOuts []*TxOut
	var txIns []*TxIn
	total := 0
	// 사용되지 않은 UTXO Output만 가져온다.
	uTxOuts := UTxOutsByAddress(from, Blockchain())
	// 더 많거나 같은 금액을 가지고 있으면
	for _, uTxOut := range uTxOuts {
		if total >= amount {
			break
		}
		txIn := &TxIn{uTxOut.TxID, uTxOut.Index, from}
		txIns = append(txIns, txIn)
		total += uTxOut.Amount
	}
	if change := total - amount; change != 0 {
		changeTxOut := &TxOut{from, change}
		txOuts = append(txOuts, changeTxOut)
	}
	txOut := &TxOut{to, amount}
	txOuts = append(txOuts, txOut)
	tx := &Tx{
		ID:        "",
		Timestamp: int(time.Now().Unix()),
		TxIns:     txIns,
		TxOuts:    txOuts,
	}
	tx.getId()
	return tx, nil

}

func (m *mempool) AddTx(to string, amount int) error {
	tx, err := makeTx("fdongfdong", to, amount)
	if err != nil {
		return err
	}
	m.Txs = append(m.Txs, tx)
	return nil
}

// 승인할 트랜잭션 가져오기
func (m *mempool) TxToConfirm() []*Tx {
	coinbase := makeCoinbaseTx("fdongfdong")
	txs := m.Txs
	txs = append(txs, coinbase)
	// Mempool을 비워주고 Transaction을 반환한다.
	m.Txs = nil
	return txs
}
