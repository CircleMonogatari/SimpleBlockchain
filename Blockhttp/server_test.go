package Blockhttp

import (
	"encoding/json"
	"github.com/CircleMonogatari/SimpleBlockchain/Cli"
	"io/ioutil"
	"net/http"
	"testing"

	"fmt"
)

func Test_SendServer(t *testing.T) {
	cli := Cli.GetInstance()
	cli.Localhost = "121.37.236.234:8080"
	cli.SendAddress()
}

func Test_findByPk(t *testing.T) {
	fmt.Println("测试版本 ")
	url := "http://" + "121.37.236.234:8080"

	resp, err := http.Get(url + "/version")
	if err != nil {

		t.Fatal(err)
	}
	defer resp.Body.Close()
	//
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal("ReadAll error")
	}

	fmt.Println(body)
	fmt.Println(string(body))

	var mapResult map[string]interface{}
	err = json.Unmarshal(body, &mapResult)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(mapResult["version"])

	//var hahaha map[string]interface{}
	//json.NewDecoder(resp.Body).Decode(&hahaha)
	//fmt.Println(hahaha)
	//fmt.Println("End")
}
