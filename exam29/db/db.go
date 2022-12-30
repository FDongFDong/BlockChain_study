package db

import (
	"coin/exam29/utils"
	"fmt"

	"github.com/boltdb/bolt"
)

const (
	dbname       = "blockchain.db"
	dataBucket   = "data"
	blocksBucket = "blocks"
	checkpoint   = "checkpoint"
)

var db *bolt.DB

// DB initialize, Singleton pattern형식
func DB() *bolt.DB {
	if db == nil {
		// init db
		// path는 DB의 이름, 파일이 없으면 자동으로 생성해준다,
		dbPointer, err := bolt.Open(dbname, 0600, nil)
		utils.HandleErr(err)
		db = dbPointer
		// bucket이 존재하지 않으면 생성시켜주는 Transaction, 두개의 bucket을 만들어준다.
		// bucket는 table 같은거다
		err = db.Update(func(tx *bolt.Tx) error {
			_, err := tx.CreateBucketIfNotExists([]byte(dataBucket))
			utils.HandleErr(err)
			_, err = tx.CreateBucketIfNotExists([]byte(blocksBucket))

			return err
		})
		utils.HandleErr(err)
	}
	return db
}

func SaveBlock(hash string, data []byte) {
	fmt.Printf("Saving Block %s\nData: %b\n", hash, data)
	err := DB().Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(blocksBucket))
		err := bucket.Put([]byte(hash), data)
		return err
	})
	utils.HandleErr(err)
}
func SaveBlockchain(data []byte) {
	err := DB().Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(dataBucket))
		err := bucket.Put([]byte(checkpoint), data)
		return err
	})
	utils.HandleErr(err)
}

// checkpoint가 있는지 없는지 리턴하는 함수
func Checkpoint() []byte {
	var data []byte
	// View : Read Only
	DB().View(func(tx *bolt.Tx) error {
		// bucket을 가져온다
		bucket := tx.Bucket([]byte(dataBucket))
		data = bucket.Get([]byte(checkpoint))
		return nil
	})
	return data
}

func Block(hash string) []byte {
	var data []byte
	DB().View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(blocksBucket))
		data = bucket.Get([]byte(hash))
		return nil
	})
	return data
}
