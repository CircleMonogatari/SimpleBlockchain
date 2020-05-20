package main

import (
	"fmt"
	"github.com/Circlemono/simpelBlock/Block"
)

func main() {
	fmt.Println("进入区块链")

	cli := Block.CLI{}
	cli.Run()

}
