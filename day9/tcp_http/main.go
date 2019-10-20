package main

import(
	"fmt"
	"io"
	"net"
)

func main(){
	conn,err := net.Dial("tcp","www.baidu.com:80")
	if err!=nil{
		fmt.Println("Error dialing",err.Error())
		return
	}
	defer conn.Close()
	msg := "GET / HTTP/1.1\r\n"
	msg +="Host: www.baidu.com\r\n"
	msg +="Connection: close\r\n"
	// msg +="User-Agent: Mozilla/5.0 (iPhone; CPU iPhone OS 11_0 like Mac OS X) AppleWebKit/604.1.38 (KHTML, like Gecko) Version/11.0 Mobile/15A372 Safari/604.1"
	msg +="\r\n\r\n"

	n,err := io.WriteString(conn,msg)
	if err != nil{
		fmt.Println("write string failed,",err)
		return
	}
    fmt.Println("send to baidu.com bytes:",n)
	buf := make([]byte,4096)
	for{
		count,err := conn.Read(buf)
		fmt.Println("count:",count,"err:",err)
		if err != nil{
			break
		}
		 fmt.Println(string(buf[0:count]))
	}

}