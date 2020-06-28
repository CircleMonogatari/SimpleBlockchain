package Block

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"
	"encoding/hex"
	"fmt"
	"log"
)

//交易体
type Transaction struct {
	ID   []byte     //ID
	Vin  []TXInput  //输入
	Vout []TXOutput //输出
	Data string     //交易数据
}

//交易输入
type TXInput struct {
	Txid      []byte //交易ID, 一个交易输入引用了之前一笔交易的一个输出, ID 表明是之前哪笔交易
	Vout      int    //输出索引
	ScriptSig string
}

//交易输出
type TXOutput struct {
	Value        int
	ScriptPubKey string
}

//检查交易是否为 coinbase
func (tx Transaction) IsCoinbase() bool {
	return len(tx.Vin) == 1 && tx.Vin[0].Vout == -1
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
func NewCoinbaseTX(to, data string, subsidy int) *Transaction {
	if data == "" {
		data = fmt.Sprintf("Reward to '%s'", to)
	}

	txin := TXInput{[]byte{}, -1, genesisCoinbaseData}
	txout := TXOutput{subsidy, to}

	tx := Transaction{nil, []TXInput{txin}, []TXOutput{txout}, data}
	tx.SetID()
	//tx.SetID()

	return &tx
}

//货币交易
func NewUTXOTransaction(from, to, data string, amount int, bc *BlockChain) (*Transaction, error) {
	var inputs []TXInput
	var outputs []TXOutput

	//找到符合条件的输出
	acc, validOutputs := bc.FindSpendableOutputs(from, amount)
	if acc < amount {
		log.Println("Error: Not enough funds ", acc)
		return nil, fmt.Errorf("余额不足")
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

	tx := Transaction{nil, inputs, outputs, data}
	tx.SetID()

	return &tx, nil
}

//通过Txid指定交易
func NewUTxIdTransaction(from, to, data string, txid []byte, bc *BlockChain) (*Transaction, error) {
	var inputs []TXInput
	var outputs []TXOutput

	////查找指定txid交易
	//if bc.find

	//查找交易是否被使用
	if bc.FindIsSpendableOutputs(txid, 0) != true {
		return nil, fmt.Errorf("该交易已经被使用")
	}

	input := TXInput{txid, 0, from}
	inputs = append(inputs, input)
	outputs = append(outputs, TXOutput{1, to})

	tx := Transaction{nil, inputs, outputs, data}
	tx.SetID()

	return &tx, nil
}

//验证输入持有者
func (in *TXInput) CanUnlockOutputWith(unlockingData string) bool {
	return in.ScriptSig == unlockingData
}

//验证输出持有者
func (out *TXOutput) CanBeUnlockedWith(unlockingData string) bool {
	return out.ScriptPubKey == unlockingData
}
