package Blockhttp

import (
	"bytes"
	"encoding/gob"
	"encoding/hex"
	"fmt"
	"github.com/CircleMonogatari/SimpleBlockchain/Cli"
	_ "github.com/CircleMonogatari/SimpleBlockchain/docs"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Runserver() {
	r := gin.Default()

	r.Use(Cors())
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.GET("/", Root)

	//license

	//Demo接口
	r.GET("/users", Users)                     //用户列表
	r.GET("/transactionlist", Transactionlist) //区块链数据

	//前端
	r.POST("/entry", Entry)                    //数据录入
	r.GET("/balance", Balance)                 //用户余额
	r.GET("/balancedetailed", BalanceDetailed) //余额明细
	r.POST("/transaction", Transaction)        //茶叶交易

	r.POST("/teadata", TeaData) //茶叶数据

	//服务端
	r.GET("/registerinfo", RegisterInfo) //服务器列表
	r.GET("/register", Register)         //注册服务器
	r.GET("/version", Version)           //当前区块链版本
	r.POST("/blockchain", BlockChain)    //区块链数据

	//license
	l := r.Group("/license")
	{
		l.POST("/entry", licenseEntry)
		l.POST("/send", licenseSend)
		l.GET("/node", licenseNode)
		l.GET("/nodelist", licenseNodeList)
	}

	r.GET("/test", DemoData)

	r.Run() // listen and serve on 0.0.0.0:8080

	fmt.Println("WEB END")
}

//@Summary 申请证书表单数据
//@Description 获取证书所有数据
//@Tags license
//@Param txid query string false "FxjaxE4MlGgnMuuiPmo6lDko00q1Hzcg1Bip+Nf8iQs="
//@Success 200 {object} gin.H {"statuc":"ok"}
//@Failure 400 {object} gin.H {"statuc":"error","msg":"失败原因"}
//@Router /license/nodelist [get]
func licenseNodeList(c *gin.Context) {
	txid := c.DefaultQuery("txid", "")
	if txid == "" {
		c.JSON(http.StatusBadRequest, gin.H{"status": "error"})
		return
	}

	cli := Cli.GetInstance()
	txs := cli.GetNodeList(txid)
	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
		"data":   txs,
	})

}

//@Summary 获取申请的表单
//@Description 获取指定用户创建的申请表单
//@Tags license
//@Param address query string false "sfr"
//@Success 200 {object} gin.H {"statuc":"ok"}
//@Failure 400 {object} gin.H {"statuc":"error","msg":"失败原因"}
//@Router /license/node [get]
func licenseNode(c *gin.Context) {
	address := c.DefaultQuery("address", "")
	if address == "" {
		c.JSON(http.StatusBadRequest, gin.H{"status": "error"})
		return
	}

	cli := Cli.GetInstance()
	tx := cli.GetNodeAll(address)
	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
		"data":   tx,
	})
}

//@Summary 证书交易
//@Description 把指定证书交易给指定对象
//@Tags license
//@Param data formData string false "{json}"
//@Param address formData string false "sfr"
//@Param to formData string false "cs"
//@Param txid formData string false "FxjaxE4MlGgnMuuiPmo6lDko00q1Hzcg1Bip+Nf8iQs="
//@Success 200 {object} gin.H {"statuc":"ok"}
//@Failure 400 {object} gin.H {"statuc":"error","msg":"失败原因"}
//@Router /license/send [post]
func licenseSend(c *gin.Context) {
	cli := Cli.GetInstance()
	address := c.PostForm("address")
	to := c.PostForm("to")
	txid := c.PostForm("txid")
	data := c.PostForm("data")

	err := cli.SendTxid(address, to, data, txid)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "error",
			"msg":    err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
	})
}

//@Summary 申请表单
//@Description 申请表单创建后, 生成第一个交易数据, 并返回该表单的ID
//@Tags license
//@Param data formData string false "{json}"
//@Param address formData string true "Ivan"
//@Success 200 {object} gin.H {"statuc":"ok"}
//@Failure 400 {object} gin.H {"statuc":"error","msg":"失败原因"}
//@Router /license/entry [post]
func licenseEntry(c *gin.Context) {
	cli := Cli.GetInstance()
	address := c.PostForm("address")
	data := c.PostForm("data")

	err := cli.Entry(address, data, 1)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "error",
			"msg":    err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
	})
}

func DemoData(context *gin.Context) {
	a := "qwert"

	context.JSON(http.StatusOK, gin.H{
		"a": a,
		"b": []byte(a),
	})
}

// @Summary 区块链数据
// @Description 返回当前区块链里的所有交易数据
// @Tags Demo
// @Success 200 {object} gin.H "{"data":{}}"
// @Router /transactionlist [get]
func Transactionlist(c *gin.Context) {

	cli := Cli.GetInstance()

	c.JSON(http.StatusOK, gin.H{
		"data": cli.GetTranList(),
	})
}

// @Summary 注册服务器到中心服务器中
// @Description 注册当前服务器信息到中心服务器中
// @Tags 服务端
// @Param mode formData int true "5000"
// @Success 200 {object} gin.H "{"statuc":"ok"}"
// @Router /register [get]
func Register(c *gin.Context) {
	clientaddr := c.ClientIP()
	cli := Cli.GetInstance()

	modestr := c.PostForm("mode")

	mode, err := strconv.Atoi(modestr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "error",
			"data":   err,
		})
		return
	}

	//注册服务器
	cli.Register(Cli.Servertype(mode), clientaddr)

	//数据转为二进制
	var result bytes.Buffer
	encoder := gob.NewEncoder(&result)
	err = encoder.Encode(cli.GetServerList())
	if err != nil {
		log.Println(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"status":   "ok",
		"data":     cli.GetServerList(),
		"databyte": result.Bytes(),
	})
}

// @Summary 用户数据
// @Description 获取所有区块链中的用户地址(在实际的区块链中该地址是保密的, 当前为demo演示接口)
// @Tags Demo
// @Success 200 {object} gin.H "{"data":["sadhaj","Pedro","Ivan"],"statuc":"ok"}"
// @Router /users [get]
func Users(context *gin.Context) {
	cli := Cli.GetInstance()

	data := cli.Users()
	log.Println(data)

	context.JSON(http.StatusOK, gin.H{
		"statuc": "ok",
		"data":   data,
	})
}

// @Summary 茶叶数据
// @Description 获取指定地址的茶叶数据
// @Tags 前端
// @Param address formData string true "ASHASDSABDKJQWFKJBASFKAF"
// @Success 200 {object} gin.H {"statuc":"ok", "data":""}
// @Failure 400 {object} gin.H {"statuc":"error", "message":"失败原因"}
// @Router /teadata [post]
func TeaData(context *gin.Context) {

}

// @Summary 数据录入
// @Description 生产茶叶数据， 用于交易
// @Tags 前端
// @Param address formData string true "Ivan"
// @Param amount formData int true "5000"
// @Param data formData string true "json"
// @Success 200 {object} gin.H {"statuc":"ok", "data":""}
// @Failure 400 {object} gin.H {"statuc":"error", "message":"失败原因"}
// @Router /entry [post]
func Entry(c *gin.Context) {
	cli := Cli.GetInstance()
	address := c.PostForm("address")
	amountstr := c.PostForm("amount")
	data := c.PostForm("data")

	amount, err := strconv.Atoi(amountstr)

	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "error",
			"data":   err,
		})
		return
	}

	cli.Entry(address, data, amount)
	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
	})
}

// @Summary 茶叶交易
// @Description 用于两个不同地址之间的数据交易
// @Tags 前端
// @Param from formData string true "Ivan"
// @Param to formData string true "Ble"
// @Param amount formData int true "300"
// @Param data formData string false "{}"
// @Success 200 {object} gin.H {"statuc":"ok"}
// @Failure 400 {object} gin.H {"statuc":"error", "data":"失败原因"}
// @Router /transaction [post]
func Transaction(c *gin.Context) {
	cli := Cli.GetInstance()
	from := c.PostForm("from")
	to := c.PostForm("to")
	data := c.PostForm("data")
	amountstr := c.PostForm("amount")

	amount, err := strconv.Atoi(amountstr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "error",
			"data":   err.Error(),
		})
		return
	}

	err = cli.Send(from, to, data, amount)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "error",
			"data":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
	})

}

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Next()
	}
}

// @Summary 当前区块链数据
// @Description 用于同步本地区块链数据
// @Tags 服务端
// @Success 200 {object} gin.H "{"statuc":"ok", "databyte":"bytesdata"}"
// @Router /blockchain [post]
func BlockChain(c *gin.Context) {
	cli := Cli.GetInstance()
	log.Printf("发送完毕! 共 %d 字节\n", len(cli.GetBlockChain()))

	log.Println("databyte:")
	log.Println(hex.EncodeToString(cli.GetBlockChain()))

	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
		"data":   hex.EncodeToString(cli.GetBlockChain()),
	})

}

func Root(c *gin.Context) {
	c.HTML(200, "index.html", nil)
}

// @Summary 返回注册信息
// @Description 返回指当前所有的注册信息
// @Tags 服务端
// @Success 200 {object} gin.H "{"statuc":"ok", "addres": []}"
// @Router /registerinfo [get]
func RegisterInfo(c *gin.Context) {
	cli := Cli.GetInstance()

	c.JSON(200, gin.H{
		"status": "ok",
		"addres": cli.GetServerList(),
	})
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
		return
	}

	cli := Cli.GetInstance()

	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
		"data":   cli.GetBalanceDetails(address),
	})
}

// @Summary 余额
// @Description 返回指定用户的余额信息
// @Tags 前端
// @Param address query string true "Ivan"
// @Success 200 {object} gin.H
// @Router /balance [get]
func Balance(c *gin.Context) {
	address := c.DefaultQuery("address", "")
	if address == "" {
		c.JSON(http.StatusBadRequest, gin.H{"status": "error"})
		return
	}

	cli := Cli.GetInstance()
	balance := cli.GetBalance(address)

	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
		"data":   balance,
	})
}

// @Summary 区块链版本
// @Description 返回当前区块链长度
// @Tags 服务端
// @Success 200 {object} gin.H "{"version":20}"
// @Router /version [get]
func Version(c *gin.Context) {
	cli := Cli.GetInstance()
	version := cli.GetVersion()
	c.JSON(200, gin.H{
		"version": version,
	})
}
