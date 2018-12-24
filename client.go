package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	//客户端主动连接服务器
	conn, err := net.Dial("tcp", "127.0.0.1:8000")

	if err != nil {
		log.Fatal(err) //log.Fatal() 会产生panic
		return
	}

	defer conn.Close()        //关闭
	buf := make([]byte, 1024) //缓冲区
	for {
		fmt.Printf("请输入发送的内容:")
		fmt.Scan(&buf)
		fmt.Printf("发送的内容: %s\n", string(buf))

		//发送数据
		conn.Write(buf)

		//阻塞等待服务器回复的数据
		n, err := conn.Read(buf) //n代码接收数据的长度
		if err != nil {
			fmt.Println(err)
			return
		}

		//切片截取， 只截取有效数据
		result := buf[:n]
		fmt.Printf("接收到数据[%d]:%s\n", n, string(result))
	}
}
