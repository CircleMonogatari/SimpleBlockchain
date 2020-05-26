package Cli

import (
	"flag"
	"fmt"
	"log"
	"os"
)

type CLI struct {
	Localhost string
	mode      int
	Servers   []Serverinfo
}

type Servertype int

const (
	CentralServer Servertype = iota + 1
	TransactionServer
	InteractiveServer
)

type Serverinfo struct {
	ServerType Servertype
	Address    string
}

var Cli *CLI

func GetInstance() *CLI {
	if Cli == nil {
		Cli = &CLI{}
	}
	return Cli
}

func (cli *CLI) Register(mode Servertype, address string) {
	cli.Servers = append(cli.Servers, Serverinfo{mode, address})
}

func (cli *CLI) Run() {
	cli.validateArgs()

	mode := flag.Int("mode", 0, "服务器类型: 1 中心服务器; 2 功能服务器; 3 节点服务器")
	localhost := flag.String("localhost", "121.37.236.234:8080", "填入中心服务器IP")

	flag.Parse()

	if *mode == 0 {
		log.Printf("服务器类型不能为空")

	}
	cli.mode = *mode
	cli.Localhost = *localhost

	//注册信息到中心服务器
	if cli.mode != 1 {
		cli.SendAddress()
		cli.Syncdata()
	}

}

func (cli *CLI) printUsage() {
	fmt.Println("Usage:")
	fmt.Println("  mode CentralServer is 0 ; ")
}

func (cli *CLI) validateArgs() {
	if len(os.Args) < 2 {
		cli.printUsage()
		os.Exit(1)
	}
}
