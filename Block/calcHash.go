package Block

import (
	"bytes"
	"crypto/sha256"
	"encoding/binary"
	"log"
	"math"
	"math/big"
)

const targetBits = 24

type ProofOfWork struct {
	BlockData *BlockData
	target    *big.Int
}

func NewProofOfWork(b *BlockData) *ProofOfWork {
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
func (pow *ProofOfWork) PrepareData(nonce int) []byte {
	data := bytes.Join([][]byte{
		pow.BlockData.PrevBlockHash,
		pow.BlockData.HashTransactions(),
		IntToHex(pow.BlockData.Timestamp),
		IntToHex(int64(targetBits)),
		IntToHex(int64(nonce)),
	},
		[]byte{},
	)

	return data
}

//计算哈希值
func (pow *ProofOfWork) Run() (int, []byte) {
	var hashInt big.Int
	var hash [32]byte
	nonce := 0

	for nonce < math.MaxInt64 {
		data := pow.PrepareData(nonce)
		hash = sha256.Sum256(data)
		hashInt.SetBytes(hash[:])

		if hashInt.Cmp(pow.target) == -1 {
			break
		} else {
			nonce++
		}
	}
	return nonce, hash[:]
}

//工作量证明验证

func (pow *ProofOfWork) Validate() bool {
	var hashInt big.Int

	data := pow.PrepareData(pow.BlockData.Nonce)
	hash := sha256.Sum256(data)
	hashInt.SetBytes(hash[:])

	isValid := hashInt.Cmp(pow.target) == -1

	return isValid
}
