package Block

import (
	"crypto/sha256"
	"encoding/hex"
	"time"
)

//计算区块哈希值
func calculateHash(block *BlockData) string {
	record := string(block.Index) + block.Timestamp + string(block.BPM) + block.PrevHash
	h := sha256.New()
	h.Write([]byte(record))
	hashed := h.Sum(nil)
	return hex.EncodeToString(hashed)
}

//生成新的区块
func generateBlock(oldBlock *BlockData, BPM int) (BlockData, error) {
	var block BlockData

	t := time.Now()

	block.PrevHash = oldBlock.PrevHash
	block.Index = oldBlock.Index + 1
	block.Timestamp = t.String()
	block.BPM = BPM
	block.Hash = calculateHash(&block)

	return block, nil
}

//判断区块是否有效
func isBlockValid(newBlock, oldBlock BlockData) bool {
	if newBlock.Index != oldBlock.Index+1 {
		return false
	}
	if newBlock.PrevHash != oldBlock.Hash {
		return false
	}
	if calculateHash(&newBlock) != newBlock.Hash {
		return false
	}
	return true
}

func replaceChain(newBlocks []BlockData) {
	if len(newBlocks) > len(Blockchain) {
		Blockchain = newBlocks
	}
}
