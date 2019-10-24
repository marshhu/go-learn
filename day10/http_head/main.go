package main

import(
	"fmt"
	"net"
	"net/http"
	"time"
)

var url = []string{
	"http://www.baidu.com",
	"http://www.google.com",
	"http://www.taobao.com",
}

func main(){
	for _,v := range url{
		c := http.Client{
			Transport:&http.Transport{
				Dial:func(network,addr string)(net.Conn,error){
					timeout := time.Second*2
					return net.DialTimeout(network,addr,timeout)
				},
			},
		}
		// c := http.Client{
		// 	Timeout:time.Second*5,
		// }
		
		resp,err := c.Head(v)
		if err != nil{
			fmt.Printf("head %s failed,err:%v\n",v,err)
			continue
		}
		fmt.Printf("head success,status:%v\n",resp.Status)
	}
}