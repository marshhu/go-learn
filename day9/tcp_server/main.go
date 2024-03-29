package main

import (
	"fmt"
	"net"
)

func main(){
	fmt.Println("start server...")
	listen,err := net.Listen("tcp",":9080")
	if err != nil{
		fmt.Println("listen failed,err:",err)
		return
	}
	fmt.Println("server running..")
	for{
		conn,err := listen.Accept()
		if err!=nil{
			fmt.Println("accept failed,err:",err)
			continue
		}
		go process(conn)
	}
}

func process(conn net.Conn){
	defer conn.Close()
	for{
		buf := make([]byte,512)
		n,err := conn.Read(buf)
		if err != nil{
			fmt.Println("read err:",err)
			return
		}
		fmt.Println(string(buf[0:n]))
	}
}