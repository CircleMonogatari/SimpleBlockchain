package main

import (
	"fmt"
	"github.com/CircleMonogatari/SimpleBlockchain/Block/Cli"
	"github.com/CircleMonogatari/SimpleBlockchain/blockhttp"
)

func main() {
	fmt.Println("进入区块链")

	//命令行处理
	cli := Cli.GetInstance()
	cli.Run()
	//启动服务器
	blockhttp.Runserver()
}
