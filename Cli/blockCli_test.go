package Cli

import (
	"encoding/base64"
	"fmt"
	"testing"
)

func Test_SendServer(t *testing.T) {
	str := []byte("1234567890")
	//bstr, _ := hex.DecodeString(str)
	//fmt.Println(bstr)
	//enstr := hex.EncodeToString(str[:])
	fmt.Println(str)

	fmt.Println(string(str))
}

func TestCLI_GetTranList(t *testing.T) {

	cli := GetInstance()

	cli.Entry("sfr", "ceshi", 1)

	cli.Send("sfr", "cs", "cshi2", 1)

	ts := cli.GetTranList()
	fmt.Println(string(ts[2].ID))
}

func TestCLI_GetNodeList(t *testing.T) {
	cli := GetInstance()

	//生成交易
	cli.Entry("sfr", "ceshi", 1)

	//获取数据
	txs := cli.GetNodeAll("sfr")
	if txs == nil {
		t.Fatal("GetNodeAll error")
	}

	txId := txs[0].ID

	id := base64.StdEncoding.EncodeToString(txId)

	//发起交易
	err := cli.SendTxid("sfr", "cs", "申请回报", id)
	if err != nil {
		t.Fatal("SendTxid error")
	}

	//交易溯源
	fmt.Println(cli.GetNodeList("sfr"))
}

//
//func TestCLI_GetNodeList2(t *testing.T) {
//	cli := GetInstance()
//	cli.GetNodeList()
//}
