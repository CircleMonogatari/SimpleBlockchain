package Block

import "fmt"

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

////检查交易是否为 coinbase
//func (tx Transaction) IsCoinbase() bool {
//	return len(tx.Vin) == 1 && tx.Vin[0].Txid == -1 && tx.Vin[0].Vout == -1
//}

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
