package Cli

import (
	"encoding/hex"
	"fmt"
	"testing"
)

func Test_SendServer(t *testing.T) {
	str := []byte("123")
	//bstr, _ := hex.DecodeString(str)
	//fmt.Println(bstr)
	enstr := hex.EncodeToString(str[:])
	fmt.Println(enstr)

	fmt.Println(string([]byte(enstr)))
}
