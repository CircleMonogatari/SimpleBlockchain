package Block

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"time"
)

type BlockData struct {
	Timestamp     int64
	Data          []byte
	PrevBlockHash []byte
	Hash          []byte
	Nonce         int
}

func init() {
	fmt.Println("数据创建")
}

//生成新的区块
func NewBlock(data string, prevBlockHash []byte) *BlockData {
	block := &BlockData{time.Now().Unix(), []byte(data), prevBlockHash, []byte{}, 0}
	pow := NewProofOfWork(block)
	nonce, hash := pow.Run()

	block.Hash = hash[:]
	block.Nonce = nonce

	return block
}

//生成创世区块
func NewGenesisBlock() *BlockData {
	return NewBlock("Genesis Block", []byte{})
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
