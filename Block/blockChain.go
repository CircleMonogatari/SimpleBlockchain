package Block

import (
	"encoding/hex"
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

//找到相关的交易
func (bc *BlockChain) FindUnspentTransactions(address string) []Transaction {
	var unspentTXs []Transaction
	spentTXOs := make(map[string][]int)
	bci := bc.Iterator()

	for {
		block := bci.Next()

		//遍历交易
		for _, tx := range block.Transactions {
			//获取交易ID
			txID := hex.EncodeToString(tx.ID)
			//编译所有输出
		Outputs:
			for outIdx, out := range tx.Vout {
				if spentTXOs[txID] != nil {
					for _, spentOut := range spentTXOs[txID] {
						if spentOut == outIdx {
							continue Outputs
						}
					}
				}

				if out.CanBeUnlockedWith(address) {
					unspentTXs = append(unspentTXs, *tx)
				}
			}

			if tx.IsCoinbase() == false {
				for _, in := range tx.Vin {
					if in.CanUnlockOutputWith(address) {
						inTxID := hex.EncodeToString(in.Txid)
						spentTXOs[inTxID] = append(spentTXOs[inTxID], in.Vout)
					}
				}
			}
		}
		//跳出循环
		if len(block.PrevBlockHash) == 0 {
			break
		}
	}

	return unspentTXs
}

//查找交易中 至少amount的 UTXO
func (bc *BlockChain) FindSpendableOutputs(address string, amount int) (int, map[string][]int) {
	unspentOutputs := make(map[string][]int) //未交易输出
	unspentTXs := bc.FindUnspentTransactions(address)
	accumulated := 0

Work:
	for _, tx := range unspentTXs {
		txID := hex.EncodeToString(tx.ID)

		for outIdx, out := range tx.Vout {
			if out.CanBeUnlockedWith(address) && accumulated < amount {
				accumulated += out.Value
				unspentOutputs[txID] = append(unspentOutputs[txID], outIdx)

				if accumulated >= amount {
					break Work
				}
			}
		}
	}

	return accumulated, unspentOutputs
}
