package main

import (
	"fmt"
	"github.com/Circlemono/simpelBlock/Block"
)

func main() {
	fmt.Println("进入区块链")

	bc := Block.NewBlockchain("大萝卜家")
	defer bc.DB.Close()

	cli := Block.CLI{bc}
	cli.Run()
	//fmt.Println("开始添加区块")
	//bc.AddBlock("Send 1 BTC to Ivan")
	//bc.AddBlock("Send 2 more BTC to Ivan")
	//
	//
	//fmt.Println("展示区块")
	//for _, block := range bc.Blocks{
	//	fmt.Printf("Prev.hash: x\n", block.PrevBlockHash)
	//	fmt.Printf("data: x\n", block.Data)
	//	fmt.Printf("Hash: x\n", block.Hash)
	//	fmt.Println()
	//}

}
