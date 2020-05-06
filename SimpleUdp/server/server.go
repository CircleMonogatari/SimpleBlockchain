package main

import (
	"fmt"
	"log"
	"net"
	"time"
)

func main()  {
	lister, err := net.ListenUDP("udp", &net.UDPAddr{IP:net.IPv4zero, Port: 9981})
	if err != nil {
		fmt.Println(err)
		return
	}

	log.Printf("本地地址:<%s>\n", lister.LocalAddr().String())
	peers := make([]net.UDPAddr, 0, 2)
	data := make([]byte, 1024)
	for  {
		n, remoteAddr, err := lister.ReadFromUDP(data)
		if err != nil {
			fmt.Printf("error during read:%s", err)
		}
		log.Printf("<%s> %s\n", remoteAddr.String(), data[:n])
		peers = append(peers, *remoteAddr)
		if len(peers) == 2 {
			log.Printf("进行UDP打洞%s <---> %s 连接\n", peers[0].String(), peers[1].String())
			lister.WriteToUDP([]byte(peers[1].String()), &peers[0])
			lister.WriteToUDP([]byte(peers[0].String()), &peers[1])
			log.Println("服务器退出Ing...")
			time.Sleep(time.Second * 800)
			log.Println("服务器退出")
			return
		}
	}

}