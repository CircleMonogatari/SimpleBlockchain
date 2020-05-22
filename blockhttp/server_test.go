package blockhttp

import (
	"encoding/json"
	"net/http"
	"testing"

	"fmt"
)

func Test_findByPk(t *testing.T) {
	fmt.Println("测试版本 ")
	resp, err := http.Get("http://121.37.236.234:8080/version")
	if err != nil {
		t.Fatal("http get error")
	}
	defer resp.Body.Close()
	//
	//body, err := ioutil.ReadAll(resp.Body)
	//if err != nil {
	//	t.Fatal("ReadAll error")
	//}
	//
	//fmt.Println(body)
	//fmt.Println(string(body))
	//
	//
	var hahaha map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&hahaha)
	fmt.Println(hahaha)
	fmt.Println("End")
}
