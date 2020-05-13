package Block

import (
	"flag"
	"fmt"
	"os"
)

type CLI struct {
	bc *BlockChain
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

		os.Exit(1)
	}

}

func (cli *CLI) addblock(data string) {
	cli.bc.AddBlock(data)
	fmt.Println("Success")
}

func (cli *CLI) printchain() {
	It := cli.bc.Itrrator()
}
