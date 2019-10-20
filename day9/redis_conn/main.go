package main

import(
	"fmt"
	"github.com/gomodule/redigo/redis"
)

func main(){
	conn,err := redis.Dial("tcp","localhost:6379",redis.DialPassword("123456")))
	if err != nil{
		fmt.Println("conn redis failed",err)
		return
	}
	fmt.Println("conn success")
	defer conn.Close()
}