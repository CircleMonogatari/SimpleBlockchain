package Block

import (
	"fmt"
	"time"
)

type BlockData struct {
	Index     int64
	Timestamp int64
	BPM       int
	Hash      []byte
	PrevHash  []byte
	Data  []byte
}

var Blockchain []BlockData


func init() {
	fmt.Println("数据创建")
}

//生成新的区块
func GenerateBlock(oldBlock *BlockData, BPM int) (BlockData, error) {
	var block BlockData

	t := time.Now()

	block.PrevHash = oldBlock.PrevHash
	block.Index = oldBlock.Index + 1
	block.Timestamp = t.String()
	block.BPM = BPM
	block.Hash = CalculateHash(&block)

	return block, nil
}
