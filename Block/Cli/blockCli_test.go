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
