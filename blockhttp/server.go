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
	r.GET("/Init", WebInit)


	//前端

	r.POST("/balancedetailed", BalanceDetailed)		//余额明细
	r.GET("/balance", Balance)			//用户余额
	r.POST("/transaction", Transaction)	//茶叶交易
	r.POST("/entry", Entry)				//数据录入
	r.POST("/teadata", TeaData)			//茶叶数据


	//服务端
	r.GET("/registerinfo", RegisterInfo) //服务器列表

	r.GET("/version", Version)			//当前区块链版本
	r.POST("/BlockChain", BlockChain)	//区块链数据




	r.Run() // listen and serve on 0.0.0.0:8080

	fmt.Println("WEB END")
}

// @Summary 茶叶数据
// @Description 获取指定地址的茶叶数据
// @Tags 前端
// @Param address formData string true "ASHASDSABDKJQWFKJBASFKAF"
// @Success 200 {object} gin.H {"statuc":"ok", "data":""}
// @Failure 400 {object} gin.H {"statuc":"error", "message":"失败原因"}
// @Router /teadata [POST]
func TeaData(context *gin.Context) {

}


// @Summary 数据录入
// @Description 生产茶叶数据， 用于交易
// @Tags 前端
// @Param address formData string true "Ivan"
// @Success 200 {object} gin.H {"statuc":"ok", "data":""}
// @Failure 400 {object} gin.H {"statuc":"error", "message":"失败原因"}
// @Router /entry [POST]
func Entry(c *gin.Context) {

}


// @Summary 茶叶交易
// @Description 用于两个不同地址之间的数据交易
// @Tags 前端
// @Param address formData string true "Ivan"
// @Success 200 {object} gin.H {"statuc":"ok", "data":"bytesdata"}
// @Failure 400 {object} gin.H {"statuc":"error", "message":"失败原因"}
// @Router /transaction [POST]
func Transaction(c *gin.Context) {

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

// @Summary 余额明细
// @Description 返回指定地址的交易明细
// @Tags 前端
// @Param address query string true "Ivan"
// @Success 200 {object} gin.H
// @Router /balancedetailed [get]
func BalanceDetailed(c *gin.Context) {
	address := c.DefaultQuery("address", "")
	if address == "" {
		c.JSON(http.StatusBadRequest, gin.H{"status": "error"})
	}
	c.JSON(http.StatusBadRequest, gin.H{
		"status": "ok",
	})
}

// @Summary 余额
// @Description 返回指定用户的余额信息
// @Tags 前端
// @Param address query string true "Ivan"
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
