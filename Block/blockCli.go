package Block

import (
	"bytes"
	"encoding/gob"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
)

type CLI struct {
	Localhost string
	mode      int
	Servers   []Serverinfo
}

type servertype int

const (
	CentralServer servertype = iota + 1
	TransactionServer
	InteractiveServer
)

type Serverinfo struct {
	ServerType servertype
	Address    string
}

var Cli *CLI

func GetInstance() *CLI {
	if Cli == nil {
		Cli = &CLI{}
	}
	return Cli
}

func (cli *CLI) GetServerVersion() int {
	url := "http://" + cli.Localhost

	resp, err := http.Get(url + "/version")
	if err != nil {
		log.Println(err)
	}
	defer resp.Body.Close()
	//
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
	}

	var mapResult map[string]interface{}
	err = json.Unmarshal(body, &mapResult)
	if err != nil {
		log.Println(err)
	}
	return mapResult["version"].(int)
}

func (cli *CLI) Run() {
	cli.validateArgs()

	mode := flag.Int("mode", 0, "服务器类型: 1 中心服务器; 2 功能服务器; 3 处理服务器")
	if *mode == 0 {
		log.Printf("服务器类型不能为空")

	}

	cli.mode = *mode

	localhost := flag.String("localhost", "", "填入中心服务器IP, 未填写默认为中心服务器")
	cli.Localhost = *localhost

	//getBalanceCmd := flag.NewFlagSet("getbalance", flag.ExitOnError)
	//createBlockchainCmd := flag.NewFlagSet("createblockchain", flag.ExitOnError)
	//sendCmd := flag.NewFlagSet("send", flag.ExitOnError)
	//printChainCmd := flag.NewFlagSet("printchain", flag.ExitOnError)
	//
	//getBalanceAddress := getBalanceCmd.String("address", "", "The address to get balance for")
	//createBlockchainAddress := createBlockchainCmd.String("address", "", "The address to send genesis block reward to")
	//sendFrom := sendCmd.String("from", "", "Source wallet address")
	//sendTo := sendCmd.String("to", "", "Destination wallet address")
	//sendAmount := sendCmd.Int("amount", 0, "Amount to send")
	//
	//switch os.Args[1] {
	//case "getbalance":
	//	err := getBalanceCmd.Parse(os.Args[2:])
	//	if err != nil {
	//		log.Panic(err)
	//	}
	//case "createblockchain":
	//	err := createBlockchainCmd.Parse(os.Args[2:])
	//	if err != nil {
	//		log.Panic(err)
	//	}
	//case "send":
	//	err := sendCmd.Parse(os.Args[2:])
	//	if err != nil {
	//		log.Panic(err)
	//	}
	//case "printchain":
	//	err := printChainCmd.Parse(os.Args[2:])
	//	if err != nil {
	//		log.Panic(err)
	//	}
	//default:
	//	cli.printUsage()
	//	os.Exit(1)
	//}
	//
	//if getBalanceCmd.Parsed() {
	//	if *getBalanceAddress == "" {
	//		getBalanceCmd.Usage()
	//		os.Exit(1)
	//	}
	//	cli.getBalance(*getBalanceAddress)
	//}
	//
	//if createBlockchainCmd.Parsed() {
	//	if *createBlockchainAddress == "" {
	//		createBlockchainCmd.Usage()
	//		os.Exit(1)
	//	}
	//	cli.createBlockchain(*createBlockchainAddress)
	//}
	//
	//if printChainCmd.Parsed() {
	//	cli.printChain()
	//}
	//
	//if sendCmd.Parsed() {
	//	if *sendFrom == "" || *sendTo == "" || *sendAmount <= 0 {
	//		sendCmd.Usage()
	//		os.Exit(1)
	//	}
	//
	//	cli.send(*sendFrom, *sendTo, *sendAmount)
	//}
}

//同步数据
func (cli *CLI) Syncdata() {

	if cli.GetServerVersion() != cli.GetVersion() {
		log.Printf("本机区块链版本低于集群版本, 正在同步")
		blockdata := cli.GetBlockChain()
		log.Printf("下载完毕! 共 %d 字节\n", len(blockdata))
		cli.SetBlockChain(blockdata)
		log.Println("同步完毕")
	}
	log.Println("版本一致")
}

func (cli *CLI) GetBalance(address string) []TXOutput {
	cli.Syncdata()
	bc := NewBlockchain(address)
	defer bc.DB.Close()

	UTXOs := bc.FindUTXO(address)

	return UTXOs
}

func (cli *CLI) printUsage() {
	fmt.Println("Usage:")
	fmt.Println("  mode CentralServer is 0 ; ")
	fmt.Println("  createblockchain -address ADDRESS - Create a blockchain and send genesis block reward to ADDRESS")
	fmt.Println("  printchain - Print all the blocks of the blockchain")
	fmt.Println("  send -from FROM -to TO -amount AMOUNT - Send AMOUNT of coins from FROM address to TO")
}

func (cli *CLI) validateArgs() {
	if len(os.Args) < 2 {
		cli.printUsage()
		os.Exit(1)
	}
}

func (cli *CLI) printChain() {
	bc := NewBlockchain("")
	defer bc.DB.Close()

	bci := bc.Iterator()

	for {
		block := bci.Next()

		fmt.Printf("Prev hash: %x\n", block.PrevBlockHash)
		fmt.Printf("Hash: %x\n", block.Hash)
		pow := NewProofOfWork(block)
		fmt.Printf("PoW: %s\n", strconv.FormatBool(pow.Validate()))
		fmt.Println()

		if len(block.PrevBlockHash) == 0 {
			break
		}
	}
}

//录入
func (cli *CLI) Entry(address, data string, amount int) error {
	cli.Syncdata()
	bc := NewBlockchain(address)
	defer bc.DB.Close()

	cbtx := NewCoinbaseTX(address, data, amount)

	bc.MineBlock([]*Transaction{cbtx})
	return nil
}

//交易
func (cli *CLI) Send(from, to string, amount int) error {
	cli.Syncdata()
	bc := NewBlockchain(from)
	defer bc.DB.Close()

	tx, err := NewUTXOTransaction(from, to, amount, bc)
	if err != nil {
		return err
	}
	bc.MineBlock([]*Transaction{tx})
	return nil
}

func (cli *CLI) GetVersion() int {

	bc := NewBlockchain("")
	defer bc.DB.Close()

	return bc.Version()
}

func (cli *CLI) Users() []string {

	bc := NewBlockchain("")
	defer bc.DB.Close()

	return bc.Users()

}

func (cli *CLI) GetLocalHost() string {
	return cli.Localhost
}

//获取区块链数据并ENcode
func (cli *CLI) GetBlockChain() []byte {
	cli.Syncdata()
	blockchain := NewBlockchain("")
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
	var blocks []BlockByte
	blockchain := NewBlockchain("")

	decoder := gob.NewDecoder(bytes.NewBuffer(d))
	err := decoder.Decode(&blocks)
	if err != nil {
		log.Fatal(err)
		return err
	}
	blockchain.SetBlockAll(blocks)
	return nil
}

//请求服务器版本
func (cli *CLI) ServerVersion() (int, error) {
	//获取版本
	resp, err := http.Get(cli.Localhost + "/version")
	if err != nil {
		fmt.Println(err)
		return 0, err
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))

	return 0, nil
}
