package Cli

import (
	"github.com/CircleMonogatari/SimpleBlockchain/config"
	"log"
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

	b := config.BaseInfo{}
	if b.LoadFile("./config.yaml") != nil {
		log.Fatal("config.yaml read error")
	}

	if b.Mode == 0 {
		log.Fatal("服务器类型不能为空")
	}
	cli.mode = b.Mode
	cli.Localhost = b.Localhost

	//注册信息到中心服务器
	if cli.mode != 1 {
		cli.SendAddress()
		cli.Syncdata()
	}

}
