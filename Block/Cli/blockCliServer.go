package Cli

import (
	"bytes"
	"encoding/gob"
	"github.com/CircleMonogatari/SimpleBlockchain/Block"
	"log"
	"strconv"
)

func (cli *CLI) printChain() {
	bc := Block.NewBlockchain("")
	defer bc.DB.Close()

	bci := bc.Iterator()

	for {
		block := bci.Next()

		log.Printf("Prev hash: %x\n", block.PrevBlockHash)
		log.Printf("Hash: %x\n", block.Hash)
		pow := Block.NewProofOfWork(block)
		log.Printf("PoW: %s\n", strconv.FormatBool(pow.Validate()))
		log.Println()

		if len(block.PrevBlockHash) == 0 {
			break
		}
	}
}

//获取余额
func (cli *CLI) GetBalance(address string) []Block.TXOutput {
	cli.Syncdata()
	bc := Block.NewBlockchain(address)
	defer bc.DB.Close()

	UTXOs := bc.FindUTXO(address)

	return UTXOs
}

//录入
func (cli *CLI) Entry(address, data string, amount int) error {
	cli.Syncdata()
	bc := Block.NewBlockchain(address)
	defer bc.DB.Close()

	cbtx := Block.NewCoinbaseTX(address, data, amount)

	bc.MineBlock([]*Block.Transaction{cbtx})
	return nil
}

//交易
func (cli *CLI) Send(from, to string, amount int) error {
	cli.Syncdata()
	bc := Block.NewBlockchain(from)
	defer bc.DB.Close()

	tx, err := Block.NewUTXOTransaction(from, to, amount, bc)
	if err != nil {
		return err
	}
	bc.MineBlock([]*Block.Transaction{tx})
	return nil
}

func (cli *CLI) GetVersion() int {

	bc := Block.NewBlockchain("")
	defer bc.DB.Close()

	return bc.Version()
}

func (cli *CLI) Users() []string {
	bc := Block.NewBlockchain("")
	defer bc.DB.Close()

	return bc.Users()
}

func (cli *CLI) GetServerList() []Serverinfo {

	return cli.Servers
}

func (cli *CLI) GetLocalHost() string {
	return cli.Localhost
}

//获取区块链数据并ENcode
func (cli *CLI) GetBlockChain() []byte {
	cli.Syncdata()
	blockchain := Block.NewBlockchain("")
	blocks := blockchain.GetBlockAll()

	//序列化
	var result bytes.Buffer
	encoder := gob.NewEncoder(&result)
	err := encoder.Encode(blocks)
	if err != nil {
		return nil
	}
	return result.Bytes()

}

//同步数据到DB
func (cli *CLI) SetBlockChain(d []byte) error {
	cli.Syncdata()
	var blocks []Block.BlockByte
	blockchain := Block.NewBlockchain("")

	decoder := gob.NewDecoder(bytes.NewBuffer(d))
	err := decoder.Decode(&blocks)
	if err != nil {
		log.Fatal(err)
		return err
	}
	blockchain.SetBlockAll(blocks)
	return nil
}
