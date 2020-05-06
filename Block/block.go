package Block

import (
	"fmt"
	"time"
)

type BlockData struct {
	Index     int
	Timestamp string
	BPM       int
	Hash      string
	PrevHash  string
}

var Blockchain []BlockData
const targetBits int = 24

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
