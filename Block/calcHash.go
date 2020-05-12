package Block

import (
	"bytes"
	"crypto/sha256"
	"encoding/binary"
	"encoding/hex"
	"log"
	"math"
	"math/big"
)

const targetBits = 24

type ProofOfWork struct{
	BlockData *BlockData
	target *big.Int
}


func NewProofOfWork(b * BlockData) *ProofOfWork {
	target := big.NewInt(1)

	// 0x10000000000000000000000000000000000000000000000000000000000
	// 向左便宜 256 - targetBits 位
	target.Lsh(target, uint(256-targetBits))

	return &ProofOfWork{b, target}
}

// IntToHex converts an int64 to a byte array
func IntToHex(num int64) []byte {
	buff := new(bytes.Buffer)
	err := binary.Write(buff, binary.BigEndian, num)
	if err != nil {
		log.Panic(err)
	}

	return buff.Bytes()
}

//生成
func (pow *ProofOfWork)PrepareData(nonce int) []byte {
	data := bytes.Join([][]byte{
		pow.BlockData.PrevBlockHash,
		pow.BlockData.Data,
		IntToHex(pow.BlockData.Timestamp),
		IntToHex(int64(targetBits)),
		IntToHex(int64(nonce)),
	},
	[]byte{},
	)

	return data
}

func (pow *ProofOfWork)run() (int, []byte) {
	var hashInt big.Int
	var hash [32]byte
	nonce := 0

	for nonce < math.MaxInt64{
		data := pow.PrepareData(nonce)
		hash = sha256.Sum256(data)
		hashInt.SetBytes(hash[:])

		if hashInt.Cmp(pow.target) == -1 {
			break;
		}else {
			nonce++
		}
	}
	return nonce, hash[:]
}


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
