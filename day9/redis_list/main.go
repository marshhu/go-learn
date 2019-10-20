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
	_,err = conn.Do("lpush","mylist","ofo","mobike","foo")
	if err != nil   {
		fmt.Println("redis lpush failed",err.Error())
	}
	_,err = conn.Do("rpush","mylist","bluegogo","xiaolan","xiaoming")
	if err != nil{
		fmt.Println("redis rpush failed",err.Error())
	}
	values,err := redis.Strings(conn.Do("lrange","mylist",0,10))
	if err != nil{
		fmt.Println("lrange err",err.Error())
	}
	fmt.Println("mylist is:")
	for _,v := range values{
		fmt.Println(v)
	}

	_, err = conn.Do("lpop", "mylist")
	if err != nil {
		fmt.Println("lpop failed:", err.Error())
	}

	_, err = conn.Do("rpop", "mylist")
	if err != nil {
		fmt.Println("rpop failed", err.Error())
	}
}