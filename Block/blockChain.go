package Block

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"github.com/boltdb/bolt"
	"log"
)

const dbFile = "blockchain.db"
const blocksBucket = "blocks"
const genesisCoinbaseData = "genesis Coinbase Data"
const genesisCoinbaseUser = "Ivan"

type BlockChain struct {
	tip []byte   //储存区块链的tip
	DB  *bolt.DB //数据库句柄
}

//获取第一条链
func NewBlockchain(to string) *BlockChain {
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
			touser := to
			if touser == "" {
				touser = genesisCoinbaseUser
			}

			cbtx := NewCoinbaseTX(touser, "测试创世块交易", 20)
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
	if err != nil {
		return nil
	}

	bc := BlockChain{tip, db}

	return &bc
}

//获取当前区块链长度  当做版本
func (bc *BlockChain) Version() int {
	num := 0
	it := bc.Iterator()
	for {
		block := it.Next()
		num++
		if len(block.PrevBlockHash) == 0 {
			break
		}
	}
	return num
}

type BlockByte struct {
	Key   []byte
	Value []byte
}

//获取DB中所有的Key和value
func (bc *BlockChain) GetBlockAll() []BlockByte {
	var blocks []BlockByte

	it := bc.Iterator()

	blocks = append(blocks, BlockByte{[]byte("1"), it.currentHash})
	for {
		block := it.Next()
		value := bc.GetValue(block.Hash)

		blocks = append(blocks, BlockByte{block.Hash, value})
		if len(block.PrevBlockHash) == 0 {
			break
		}
	}
	return blocks
}

func (bc *BlockChain) SetBlockAll(bs []BlockByte) {

	for _, block := range bs {
		err := bc.SetValue(block.Key, block.Value)
		if err != nil {
			log.Print(err)
		}
	}
}

func (bc *BlockChain) GetValue(key []byte) []byte {
	var value []byte
	err := bc.DB.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blocksBucket))
		value = b.Get(key)
		return nil
	})
	if err != nil {
		return nil
	}
	return value
}

func (bc *BlockChain) SetValue(key, value []byte) error {

	err := bc.DB.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blocksBucket))
		err := b.Put(key, value)
		return err
	})
	if err != nil {
		return err
	}
	return nil
}

//查找所有相关交易输出
func (bc *BlockChain) FindUTXO(address string) []TXOutput {
	var UTXOs []TXOutput
	unspentTransactions := bc.FindUnspentTransactions(address)

	for _, tx := range unspentTransactions {
		for _, out := range tx.Vout {
			if out.CanBeUnlockedWith(address) {
				UTXOs = append(UTXOs, out)
			}
		}
	}

	return UTXOs
}

func (bc *BlockChain) MineBlock(transactions []*Transaction) {
	var lastHash []byte

	err := bc.DB.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blocksBucket))
		lastHash = b.Get([]byte("1"))
		return nil
	})
	if err != nil {
		log.Panic(err)
	}

	newBlock := NewBlock(transactions, lastHash)
	err = bc.DB.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blocksBucket))
		err := b.Put(newBlock.Hash, newBlock.Serialize())
		if err != nil {
			log.Panic(err)
		}

		err = b.Put([]byte("1"), newBlock.Hash)
		if err != nil {
			log.Panic(err)
		}

		bc.tip = newBlock.Hash

		return nil
	})

}

//返回当前用户
func (bc *BlockChain) Users() []string {
	usermap := make(map[string]int)

	it := bc.Iterator()
	for {
		block := it.Next()

		for _, tr := range block.Transactions {
			for _, out := range tr.Vout {
				usermap[out.ScriptPubKey] = 1
			}
		}

		if len(block.PrevBlockHash) == 0 {
			break
		}
	}
	var users []string

	for key := range usermap {
		users = append(users, key)
	}
	return users
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

//交易溯源
func (bc *BlockChain) Traceability(address string) []Transaction {
	var Transactions []Transaction

	it := bc.Iterator()
	for {
		block := it.Next()

	optis:
		for _, tx := range block.Transactions {
			for _, in := range tx.Vin {
				if in.ScriptSig == address {
					Transactions = append(Transactions, *tx)
					continue optis
				}
			}

			for _, out := range tx.Vout {
				if out.ScriptPubKey == address {
					Transactions = append(Transactions, *tx)
					continue optis
				}
			}
		}

		if len(block.PrevBlockHash) == 0 {
			break
		}
	}

	return Transactions
}

//交易列表
func (bc *BlockChain) TransactionList() []Transaction {
	var Transactions []Transaction

	it := bc.Iterator()
	for {
		block := it.Next()

		for _, tx := range block.Transactions {
			Transactions = append(Transactions, *tx)
		}

		if len(block.PrevBlockHash) == 0 {
			break
		}
	}

	return Transactions
}

//查询当前余额
func (bc *BlockChain) Balance(address string) []TXOutput {
	var txoutputs []TXOutput
	unspentTXs := bc.FindUnspentTransactions(address)

	for _, tx := range unspentTXs {
		for _, out := range tx.Vout {
			if out.CanBeUnlockedWith(address) {
				txoutputs = append(txoutputs, out)
			}
		}
	}
	return txoutputs
}

//查找所有相关交易
//func (bc *BlockChain) FindUTXOList(address string) []Transaction {
//	var txs []Transaction
//	var tmpint []TXInput
//	unspentTransactions := bc.FindUnspentTransactions(address)
//
//	txs = append(txs, unspentTransactions...)
//
//	tmptx := bc.FindTX(tmpint)
//
//	txs = append(txs, tmptx...)
//
//	return txs
//}
//
//func (bc *BlockChain) FindTX(tmpint []TXInput) []Transaction {
//	if tmpint == nil {
//		return nil
//	}
//
//	var TxTmp []Transaction
//
//	for _, in := range tmpint {
//		if len(in.Txid) == 0 {
//			continue
//		}
//
//		tx, err := bc.FindTransaction(in.Txid)
//		if err == nil {
//			TxTmp = append(TxTmp, tx)
//		}
//	}
//
//
//
//	for _,tx := range {
//
//	}
//
//	return TxTmp
//}

func (bc *BlockChain) FindTransaction(txid []byte) (*Transaction, error) {
	bci := bc.Iterator()

	for {
		block := bci.Next()

		//遍历交易
		for _, tx := range block.Transactions {
			//获取交易ID

			if bytes.Compare(tx.ID, txid) == 0 {
				return tx, nil
			}
		}
		//跳出循环
		if len(block.PrevBlockHash) == 0 {
			break
		}
	}
	return nil, fmt.Errorf("未找到交易")
}

//寻找交易是否被引用
func (bc *BlockChain) FindTransactionNext(txid []byte) (*Transaction, error) {
	bci := bc.Iterator()

	for {
		block := bci.Next()

		//遍历交易
		for _, tx := range block.Transactions {
			for _, in := range tx.Vin {

				if bytes.Compare(in.Txid, txid) == 0 {
					return tx, nil
				}
			}
		}
		//跳出循环
		if len(block.PrevBlockHash) == 0 {
			break
		}
	}
	return nil, fmt.Errorf("未找到交易")
}

func (bc *BlockChain) FindTransactionList(txid []byte) ([]Transaction, error) {

	var transactionlistNext []Transaction
	var transactionlist []Transaction

	tx, err := bc.FindTransaction(txid)
	if err != nil {
		return nil, err
	}
	transactionlist = append(transactionlist, *tx)

	//往前找
	for {
		tx, err = bc.FindTransaction(tx.Vin[0].Txid)
		if err != nil {
			break
		}
		transactionlist = append(transactionlist, *tx)
	}

	//往后找
	for {
		tx, err = bc.FindTransactionNext(txid)
		if err != nil {
			break
		}
		transactionlistNext = append(transactionlistNext, *tx)
	}

	transactionlist = append(transactionlistNext, transactionlist...)

	return transactionlist, nil
}
