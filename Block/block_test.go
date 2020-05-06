package Block

import (
	"fmt"
	"testing"
)

func TestCalculateHash(t *testing.T) {

	var b BlockData
	b.Index = 0
	b.Timestamp = "2020-04-24 22:41:51.5694503 +0800 CST m=+0.005996401"

	b.BPM = 0
	b.PrevHash = ""

	ret := CalculateHash(&b)
	if ret != "a9b0c0344ea6cb3537538b8769d09f7caa53bc826b71883ef139cef5d6071755" {
		fmt.Println("calcHash is ", ret)
		t.Error("calculateHash Error")
	}
}

func TestGenerateBlock(t *testing.T) {

	var b BlockData
	b.Index = 0
	b.Timestamp = "2020-04-24 22:41:51.5694503 +0800 CST m=+0.005996401"

	b.BPM = 0
	b.PrevHash = ""

	NewBlock, err := GenerateBlock(&b, 4096)
	if err != nil {
		t.Error("create Block Error")
	}

	if NewBlock.BPM != 4096 {
		t.Error("NewBlock check Error")
	}

	if NewBlock.Hash != CalculateHash(&NewBlock) {
		t.Error("NewBlock Create Hash Error")
	}

	if NewBlock.Index != b.Index+1 {
		t.Error("NewBlock Index Error")
	}

	if NewBlock.PrevHash != b.Hash {
		t.Error("NewBlock PrevBlock Hash Error")
	}
}

func TestIsBlockValid(t *testing.T)  {
	var b BlockData
	b.Index = 0
	b.Timestamp = "2020-04-24 22:41:51.5694503 +0800 CST m=+0.005996401"

	b.BPM = 0
	b.PrevHash = ""

	NewBlock, err := GenerateBlock(&b, 4096)
	if err != nil {
		t.Error("create Block Error")
	}

	if IsBlockValid(NewBlock, b) != true{
		t.Error("test IsBlockValid Error")
	}

	NewBlock.Index = 5

	if IsBlockValid(NewBlock, b) != false {
		t.Error("test IsBlockValid Error")
	}
}
