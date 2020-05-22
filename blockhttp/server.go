package blockhttp

import (
	"fmt"
	_ "github.com/Circlemono/simpelBlock/docs"
	"net/http"

	"github.com/Circlemono/simpelBlock/Block"
	"github.com/gin-gonic/gin"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Runserver() {
	r := gin.Default()
	r.LoadHTMLGlob("web/*")
	//a := docs.SwaggerInfo
	//if a != nil {
	//
	//}

	r.Use(Cors())
	//url := ginSwagger.URL("http://localhost:8080/swagger/doc.json") // The url pointing to API definition
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.GET("/", Root)
	r.GET("/GetLocalHost", GetLocalHost)
	r.GET("/show", ShowTX)
	r.GET("/Init", WebInit)
	r.GET("/version", Version)

	r.Run() // listen and serve on 0.0.0.0:8080

	fmt.Println("WEB END")
}

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Next()
	}
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
	address := c.DefaultQuery("address", "")
	if address == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "ok",
		})
	}

}

// @Summary 区块链版本
// @Description 返回当前区块链长度
// @Success 200 {objectx} gin.H
// @Router /version [get]
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
