package Block

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"
	"encoding/hex"
	"fmt"
	"log"
)

//交易输入
type TXInput struct {
	Txid      []byte //交易ID
	Vout      int    //输出索引
	ScriptSig string
}

const subsidy = 1

//交易输出
type TXOutput struct {
	Value        int
	ScriptPubKey string
}

//交易体
type Transaction struct {
	ID   []byte
	Vin  []TXInput
	Vout []TXOutput
}

//检查交易是否为 coinbase
func (tx Transaction) IsCoinbase() bool {
	return len(tx.Vin) == 1 && len(tx.Vin[0].Txid) == 0 && tx.Vin[0].Vout == -1
}

//把交易结构数据打包成hash填入ID
func (tx *Transaction) SetID() {
	var encode bytes.Buffer
	var hash [32]byte

	enc := gob.NewEncoder(&encode)
	err := enc.Encode(tx)
	if err != nil {
		log.Panic(err)
	}
	hash = sha256.Sum256(encode.Bytes())
	tx.ID = hash[:]
}

//创建 coinbase 交易
func NewCoinbaseTX(to, data string) *Transaction {
	if data == "" {
		data = fmt.Sprintf("Reward to '%s'", to)
	}

	txin := TXInput{[]byte{}, -1, data}
	txout := TXOutput{subsidy, to}

	tx := Transaction{nil, []TXInput{txin}, []TXOutput{txout}}
	//tx.SetID()

	return &tx
}

//货币交易
func NewUTXOTransaction(from, to string, amount int, bc *BlockChain) *Transaction {
	var inputs []TXInput
	var outputs []TXOutput

	acc, validOutputs := bc.FindSpendableOutputs(from, amount)
	if acc < amount {
		log.Panic("Error: Not enough funds")
	}

	for txid, outs := range validOutputs {
		txID, err := hex.DecodeString(txid)
		if err != nil {
			log.Panic(err)
		}

		for _, out := range outs {
			input := TXInput{txID, out, from}
			inputs = append(inputs, input)
		}
	}

	outputs = append(outputs, TXOutput{amount, to})

	if acc > amount {
		outputs = append(outputs, TXOutput{acc - amount, from})
	}

	tx := Transaction{nil, inputs, outputs}
	tx.SetID()

	return &tx
}

func (in *TXInput) CanUnlockOutputWith(unlockingData string) bool {
	return in.ScriptSig == unlockingData
}

func (out *TXOutput) CanBeUnlockedWith(unlockingData string) bool {
	return out.ScriptPubKey == unlockingData
}