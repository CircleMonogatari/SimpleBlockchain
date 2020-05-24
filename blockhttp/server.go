package blockhttp

import (
	"fmt"
	_ "github.com/CircleMonogatari/SimpleBlockchain/docs"
	"log"
	"net/http"

	"github.com/CircleMonogatari/SimpleBlockchain/Block"
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
	r.GET("/show", ShowTX)
	r.GET("/balance", Balance)
	r.GET("/Init", WebInit)
	r.GET("/version", Version)


	//web端

	//服务端
	r.GET("/registerinfo", RegisterInfo)
	r.POST("/BlockChain", BlockChain)



	r.Run() // listen and serve on 0.0.0.0:8080

	fmt.Println("WEB END")
}

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Next()
	}
}

// @Summary 当前区块链数据
// @Description 用于同步本地区块链数据
// @Tags 服务器组
// @Success 200 {object} gin.H {"statuc":"ok", "data":"bytesdata"}
// @Router /BlockChain [POST]
func BlockChain(c *gin.Context) {
	cli := Block.GetInstance()
	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
		"data":   string(cli.GetBlockChain()),
	})
}

func Root(c *gin.Context) {
	c.HTML(200, "index.html", nil)
}



// @Summary 返回注册信息
// @Description 返回指当前所有的注册信息
// @Tags 服务端
// @Success 200 {object} gin.H
// @Router /registerinfo [get]
func RegisterInfo(c *gin.Context) {
	cli := Block.GetInstance()

	c.JSON(200, gin.H{
		"addres": cli.GetVersion(),
	})
}

func WebInit(c *gin.Context) {

}

type Response struct {
	status string `json:"status" example:"ok" format:"string"`
	data   interface{}
}

// @Summary 余额
// @Description 返回指定用户的余额信息
// @Tags 前端
// @Param address query string false "Ivan"
// @Success 200 {object} Response
// @Router /balance [get]
func Balance(c *gin.Context) {
	address := c.DefaultQuery("address", "")
	if address == "" {
		c.JSON(http.StatusBadRequest, gin.H{"status": "error"})
	}

	cli := Block.GetInstance()
	balance := cli.GetBalance(address)

	c.JSON(http.StatusBadRequest, Response{
		status: "ok",
		data:   balance,
	})
}

// @Summary 区块链版本
// @Description 返回当前区块链长度
// @Tags 前端
// @Success 200 {object} gin.H
// @Router /show [get]
func ShowTX(c *gin.Context) {
	address := c.DefaultQuery("address", "")
	if address == "" {
		c.JSON(http.StatusBadRequest, gin.H{"status": "error"})
	}
	c.JSON(http.StatusBadRequest, gin.H{
		"status": "ok",
	})
}

func Syncdata() {
	cli := Block.GetInstance()
	if cli.GetServerVersion() != cli.GetVersion() {
		log.Printf("本机区块链版本低于集群版本, 正在同步")
		blockdata := cli.GetBlockChain()
		log.Printf("下载完毕! 共 %d 字节\n", len(blockdata))
		cli.SetBlockChain(blockdata)
		log.Println("同步完毕")
	}
	log.Println("版本一致")
}

// @Summary 区块链版本
// @Description 返回当前区块链长度
// @Tags 前端
// @Success 200 {object} gin.H
// @Router /version [get]
func Version(c *gin.Context) {
	cli := Block.GetInstance()
	version := cli.GetVersion()
	c.JSON(200, gin.H{
		"version": version,
	})
}
