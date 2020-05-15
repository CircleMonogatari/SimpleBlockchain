package Block

import (
	"flag"
	"fmt"
	"os"
	"strconv"
)

type CLI struct {
	BC *BlockChain
}

func (cli *CLI) Run() {

	addBlockCmd := flag.NewFlagSet("addblock", flag.ExitOnError)
	printChainCmd := flag.NewFlagSet("printchain", flag.ExitOnError)

	addBlockData := addBlockCmd.String("data", "", "Block data")

	switch os.Args[1] {
	case "addblock":
		err := addBlockCmd.Parse(os.Args[2:])
		if err != nil {
			fmt.Println("addblock Parse error")
		}
	case "printchain":
		err := printChainCmd.Parse(os.Args[2:])
		if err != nil {
			fmt.Println("printchain Parse error")
		}
	default:
		cli.printchain()
		os.Exit(1)
	}

	if addBlockCmd.Parsed() {
		if *addBlockData == "" {
			os.Exit(1)
		}
		cli.addblock(*addBlockData)
	}
	if printChainCmd.Parsed() {
		cli.printchain()
	}

}

func (cli *CLI) addblock(data string) {

	//cli.BC.AddBlock(data)
	fmt.Println("Success")
}

func (cli *CLI) printchain() {
	It := cli.BC.Itrrator()

	for {
		block := It.Next()

		fmt.Printf("Prev.Hash: %x\n", block.PrevBlockHash)
		//fmt.Printf("Data: %s\n", block.Data)
		fmt.Printf("Hash: %x\n", block.Hash)
		pow := NewProofOfWork(block)

		fmt.Printf("pow: %s\n", strconv.FormatBool(pow.Validate()))
		fmt.Println()

		if len(block.PrevBlockHash) == 0 {
			break
		}
	}
}
