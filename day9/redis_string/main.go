package main

import(
	"fmt"
	"github.com/gomodule/redigo/redis"
)

func main(){
	conn,err := redis.Dial("tcp","localhost:6379",redis.DialPassword("123456"))
	if err != nil{
		fmt.Println("conn redis failed",err)
		return
	}
	fmt.Println("conn success")
	defer conn.Close()
	_,err = conn.Do("Set","test",100)
	if err != nil{
		fmt.Println(err)
	}
	value,err := redis.Int(conn.Do("Get","test"))
	if err != nil{
		fmt.Println("get test failed,",err)
		return
	}
	fmt.Println("test:",value)
}