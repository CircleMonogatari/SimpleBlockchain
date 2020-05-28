package Block

import (
	"fmt"
	"testing"
)

//version
func Test_Version(t *testing.T) {
	bc := NewBlockchain("")
	defer bc.DB.Close()

	fmt.Println(bc.Version())

}

func Test_Transaction(t *testing.T) {
	bc := NewBlockchain("")
	defer bc.DB.Close()

	from := "Ivan"
	to := "sfr"
	data := "测试"
	amount := 50

	tx, err := NewUTXOTransaction(from, to, data, amount, bc)
	if err != nil {
		t.Fatal(err)
	}
	bc.MineBlock([]*Transaction{tx})

	fmt.Println(bc.Version())

}

//余额测试
func Test_Balance(t *testing.T) {
	bc := NewBlockchain("")
	defer bc.DB.Close()

	fmt.Println(bc.TransactionList())

}
