package Block

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"github.com/gin-gonic/gin"
	"testing"
)

//version
func Test_Version(t *testing.T) {
	bc := NewBlockchain("")
	defer bc.DB.Close()

	fmt.Println(bc.Version())

}

//func Test_Transaction(t *testing.T) {
//	bc := NewBlockchain("")
//	defer bc.DB.Close()
//
//	from := "Ivan"
//	to := "sfr"
//	data := "测试"
//	amount := 10
//
//	tx, err := NewUTXOTransaction(from, to, data, amount, bc)
//	if err != nil {
//		t.Fatal(err)
//	}
//	bc.MineBlock([]*Transaction{tx})
//
//	fmt.Println(bc.Version())
//
//}

//结构体测试
func Test_BlockData_Transactions(t *testing.T) {
	tx := Transaction{
		ID:   []byte("12345"),
		Vin:  []TXInput{},
		Vout: []TXOutput{},
		Data: "ceshishuju",
	}

	fmt.Println(gin.H{"data": tx})
}

//余额测试
func Test_Balance(t *testing.T) {
	bc := NewBlockchain("")
	defer bc.DB.Close()

	fmt.Println(bc.TransactionList())

}

type b struct {
	Age  int
	Name string
}

type a struct {
	Bs   []b
	Name string
}

func (b *a) Serialize() []byte {
	var result bytes.Buffer
	encoder := gob.NewEncoder(&result)
	err := encoder.Encode(*b)
	if err != nil {
		return nil
	}
	return result.Bytes()
}

// Block反序列化
func DeserializeTest(d []byte) *a {
	var aa a

	decoder := gob.NewDecoder(bytes.NewBuffer(d))
	err := decoder.Decode(&aa)
	if err != nil {
		return nil
	}
	return &aa
}

func Test_struct_god(t *testing.T) {
	bb := b{18, "123"}
	aa := a{[]b{bb}, "ce"}

	abyte := aa.Serialize()

	aa2 := DeserializeTest(abyte)

	fmt.Println(aa)
	fmt.Println(aa2)

}

//序列化和反序列化
func Test_Encode_Decode(t *testing.T) {

	tx := Transaction{[]byte("123"), nil, nil, "测试"}
	bd := BlockData{0, []*Transaction{&tx}, nil, nil, 0}
	bdbyte := bd.Serialize()

	enbd := Deserialize(bdbyte)

	fmt.Println(bd.Transactions[0])
	fmt.Println(enbd.Transactions[0])

}
