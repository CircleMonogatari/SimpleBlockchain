package Cli

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

//发送本机信息到中心服务器
func (cli *CLI) SendAddress() {

	url := "http://" + cli.Localhost

	resp, err := http.Get(url + "/register")
	if err != nil {
		log.Println(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
	}

	var mapResult map[string]interface{}
	err = json.Unmarshal(body, &mapResult)
	if err != nil {
		log.Println(err)
	}

}

//同步数据
func (cli *CLI) Syncdata() {

	if cli.GetServerVersion() != cli.GetVersion() {
		log.Printf("本机区块链版本低于集群版本, 正在同步")
		blockdata := cli.GetServerBlockChain()
		log.Printf("下载完毕! 共 %d 字节\n", len(blockdata))
		cli.SetBlockChain(blockdata)
		log.Println("同步完毕")
	}
	log.Println("版本一致")
}

//获取服务器区块链
func (cli *CLI) GetServerBlockChain() []byte {
	url := "http://" + cli.Localhost

	resp, err := http.Get(url + "/BlockChain")
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
	if mapResult["status"].(string) != "ok" {
		log.Println("Get Server Blockchain data error")
		return nil
	}

	return []byte(mapResult["databyte"].(string))
}

//获取中心服务器区块链版本
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

//
////请求服务器版本
//func (cli *CLI) ServerVersion() (int, error) {
//	//获取版本
//	resp, err := http.Get(cli.Localhost + "/version")
//	if err != nil {
//		log.Println(err)
//		return 0, err
//	}
//	defer resp.Body.Close()
//	body, _ := ioutil.ReadAll(resp.Body)
//	log.Println(string(body))
//
//	return 0, nil
//}
