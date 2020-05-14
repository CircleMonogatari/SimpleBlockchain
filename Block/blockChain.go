package Block

import (
	"fmt"
	"github.com/boltdb/bolt"
)

const dbFile = "blockchain.db"
const blocksBucket = "blocks"
const genesisCoinbaseData = "genesis Coinbase Data"

type BlockChain struct {
	tip []byte   //储存区块链的tip
	DB  *bolt.DB //数据库句柄
}

//获取第一条链
func NewBlockchain(address string) *BlockChain {
	var tip []byte
	db, err := bolt.Open(dbFile, 0600, nil)
	if err != nil {
		return nil
	}

	err = db.Update(func(tx *bolt.Tx) error {
		//获取db句柄
		b := tx.Bucket([]byte(blocksBucket))

		//如果数据为空, 添加创世块
		if b == nil {
			cbtx := NewCoinbaseTX(address, genesisCoinbaseData)
			genesis := NewGenesisBlock(cbtx)
			b, err := tx.CreateBucket([]byte(blocksBucket))
			if err != nil {
				fmt.Println("NewBlockChain read error")
				return nil
			}
			err = b.Put(genesis.Hash, genesis.Serialize())
			err = b.Put([]byte("1"), genesis.Hash)
			tip = genesis.Hash
		} else {
			tip = b.Get([]byte("1"))
		}
		return nil
	})

	bc := BlockChain{tip, db}

	return &bc
}

//添加区块
func (bc *BlockChain) AddBlock(Transactions []*Transaction) {
	var lastHash []byte
	err := bc.DB.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blocksBucket)) //获取对应数据库
		lastHash = b.Get([]byte("1"))
		return nil
	})
	if err != nil {
		fmt.Println("get lastblock error ")
	}

	newBlock := NewBlock(Transactions, lastHash)

	err = bc.DB.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blocksBucket))
		err := b.Put(newBlock.Hash, newBlock.Serialize())
		if err != nil {
			fmt.Println("AddBlock Put Db Error")
			return err
		}
		err = b.Put([]byte("1"), newBlock.Hash)
		if err != nil {
			fmt.Println("AddBlock Put Db Error")
			return err
		}
		bc.tip = newBlock.Hash
		return nil
	})
}
