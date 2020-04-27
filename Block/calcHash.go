package Block

import (
	"crypto/sha256"
	"encoding/hex"
)

//计算区块哈希值
func CalculateHash(block *BlockData) string {
	record := string(block.Index) + block.Timestamp + string(block.BPM) + block.PrevHash
	h := sha256.New()
	h.Write([]byte(record))
	hashed := h.Sum(nil)
	return hex.EncodeToString(hashed)
}

//判断区块是否有效
func IsBlockValid(newBlock, oldBlock BlockData) bool {
	if newBlock.Index != oldBlock.Index+1 {
		return false
	}
	if newBlock.PrevHash != oldBlock.Hash {
		return false
	}
	if CalculateHash(&newBlock) != newBlock.Hash {
		return false
	}
	return true
}

func ReplaceChain(newBlocks []BlockData) {
	if len(newBlocks) > len(Blockchain) {
		Blockchain = newBlocks
	}
}
