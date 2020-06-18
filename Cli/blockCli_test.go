package Cli

import (
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
	ts := cli.GetNodeList("EEMUALiPqx/zHjgInJgxt0Jb9XIO4S2CnoPAfia7SXY=")

	fmt.Println(ts)
}
