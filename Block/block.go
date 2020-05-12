package Block

import (
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

var Blockchain []BlockData


func init() {
	fmt.Println("数据创建")
}

//生成新的区块
func GenerateBlock(oldBlock *BlockData, BPM int) (BlockData, error) {
	var block BlockData

	//t := time.Now()
	//
	//block.PrevHash = oldBlock.PrevHash
	//block.Index = oldBlock.Index + 1
	//block.Timestamp = t.Unix()
	//block.BPM = BPM
	//block.Hash = CalculateHash(&block)

	return block, nil
}

func NewBlock(data string, prevBlockHash []byte) *BlockData {
	block := &BlockData{time.Now().Unix(), []byte(data), prevBlockHash, []byte{}, 0}
	pow := NewProofOfWork(block)
	nonce, hash := pow.Run()

	block.Hash = hash[:]
	block.Nonce = nonce

	return block
}