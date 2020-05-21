package blockhttp

import (
	"fmt"

	"github.com/Circlemono/simpelBlock/Block"
	"github.com/gin-gonic/gin"
)

func Runserver() {
	r := gin.Default()
	r.LoadHTMLGlob("web/*")

	r.GET("/", Root)
	r.GET("/GetLocalHost", GetLocalHost)
	r.GET("/show", ShowTX)
	r.GET("/Init", WebInit)
	r.GET("/version", Version)

	r.Run() // listen and serve on 0.0.0.0:8080

	fmt.Println("WEB END")
}

func Root(c *gin.Context) {
	c.HTML(200, "index.html", nil)
}

func GetLocalHost(c *gin.Context) {
	cli := Block.GetInstance()

	c.JSON(200, gin.H{
		"addres": cli.GetVersion(),
	})
}

func WebInit(c *gin.Context) {

}

func ShowTX(c *gin.Context) {

}

func Version(c *gin.Context) {
	cli := Block.GetInstance()
	version := cli.GetVersion()
	c.JSON(200, gin.H{
		//返回version数据
		"version": version,
	})
}

func tbdata() {
	//resp, err := http.Get(cli.Localhost + "/GetVersion")
	//if err != nil {
	//
	//}
	//defer resp.Body.Close()
	//
	//body, err := ioutil.ReadAll(resp.Body)
	//if err != nil {
	//
	//}
	//
	//fmt.Println(string(body))
	////获取版本数据
}
