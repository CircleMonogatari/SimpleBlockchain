package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"strconv"
	"strings"
	"time"
)

var tag string

const HAND_SHAKE_MSG = "打洞消息"

func main()  {
	ip := net.ParseIP("39.106.149.207")
	tag = os.Args[1]
	srcAddr := &net.UDPAddr{IP:net.IPv4zero, Port:9982}
	dstAddr := &net.UDPAddr{IP: ip, Port:9981}

	conn, err := net.DialUDP("udp", srcAddr, dstAddr)
	if err != nil {
		fmt.Println(err)
	}	

	defer conn.Close()


	conn.Write([]byte("client say Hello"))

	data := make([]byte, 1024)
	n, err := conn.Read(data)
	if err != nil {
		fmt.Printf("error during read: %s", err)
	}
	conn.Close()

	fmt.Printf("read %s from <%s>\n", data[:n], conn.RemoteAddr())

	//解析远端IP
	peer := parseAddr(string(data[:n]))
	//开始打洞
	bidirectionHole(srcAddr, &peer)

}

func bidirectionHole(srcAddr *net.UDPAddr, anotherAddr *net.UDPAddr){
	conn, err := net.DialUDP("udp", srcAddr, anotherAddr)
	if err != nil {
		fmt.Println(err)
	}
	defer conn.Close()

	//向对方发送数据 打洞
	//if _, err = conn.Write([]byte(HAND_SHAKE_MSG));err!=nil{
	//	log.Printf("send handshake:", err)
	//}

	go func() {
		for  {
			time.Sleep(10 * time.Second)
			if _, err = conn.Write([]byte("from["+ tag + "]")); err != nil{
				log.Println("send msg fail ", err)
			}
		}
	}()

	for  {
		data := make([]byte, 1024)
		n, _, err := conn.ReadFromUDP(data)
		if err != nil {
			log.Printf("error during read: %s\n", err)
		}else {
			log.Printf("收到数据:%s\n", data[:n])
		}

	}


}

func parseAddr(addr string)net.UDPAddr{
	t := strings.Split(addr, ":")
	port, _ := strconv.Atoi(t[1])
	return net.UDPAddr{
		IP : net.ParseIP(t[0]),
		Port: port,
	}
}
