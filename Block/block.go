package Block

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"
	"fmt"
	"time"
)

type BlockData struct {
	Timestamp     int64
	Transactions  []*Transaction
	PrevBlockHash []byte
	Hash          []byte
	Nonce         int
}

func init() {
	fmt.Println("数据创建")
}

//生成新的区块
func NewBlock(Transactions []*Transaction, prevBlockHash []byte) *BlockData {
	block := &BlockData{time.Now().Unix(), Transactions, prevBlockHash, []byte{}, 0}
	pow := NewProofOfWork(block)
	nonce, hash := pow.Run()

	block.Hash = hash[:]
	block.Nonce = nonce

	return block
}

//生成创世区块
func NewGenesisBlock(coinbase *Transaction) *BlockData {
	return NewBlock([]*Transaction{coinbase}, []byte{})
}

//Block数据系列化
func (b *BlockData) Serialize() []byte {
	var result bytes.Buffer
	encoder := gob.NewEncoder(&result)
	err := encoder.Encode(b)
	if err != nil {
		return nil
	}
	return result.Bytes()
}

// Block反序列化
func Deserialize(d []byte) *BlockData {
	var block BlockData

	decoder := gob.NewDecoder(bytes.NewBuffer(d))
	err := decoder.Decode(&block)
	if err != nil {
		return nil
	}
	return &block
}

func (b *BlockData) HashTransactions() []byte {
	var txHashes [][]byte
	var txHash [32]byte

	for _, tx := range b.Transactions {
		txHashes = append(txHashes, tx.ID)
	}

	txHash = sha256.Sum256(bytes.Join(txHashes, []byte{}))

	return txHash[:]
}
