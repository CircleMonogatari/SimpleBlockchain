package Block

import (
	"fmt"
	"testing"
)

func TestAdd(t *testing.T) {

	var b BlockData
	b.Index = 0
	b.Timestamp = "2020-04-24 22:41:51.5694503 +0800 CST m=+0.005996401"

	b.BPM = 0
	b.PrevHash = ""

	ret := calculateHash(&b)
	if ret != "a9b0c0344ea6cb3537538b8769d09f7caa53bc826b71883ef139cef5d6071755" {
		fmt.Println("calcHash is ", ret)
		t.Error("calculateHash Error")
	}

	// sum := Add(1,2)
	// if sum == 3 {
	// 	t.Log("the result is ok")
	// } else {
	// 	t.Fatal("the result is wrong")
	// }
}
