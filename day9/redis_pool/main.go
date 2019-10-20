package main

import(
	"fmt"
	"time"
	"github.com/gomodule/redigo/redis"
)
var pool *redis.Pool

func init(){
	pool = &redis.Pool{
		MaxIdle:16,
		MaxActive:0,
		IdleTimeout:300,
		Dial:func()(redis.Conn,error){
			//1.打开连接
			c,err := redis.Dial("tcp","localhost:6379")
			if err != nil{
				fmt.Println(err)
				return nil,err
			}
			//2.访问认证
			if _,err = c.Do("AUTH","123456");err != nil{
				c.Close()
				return nil,err
			}
			return  c,nil
		},
		TestOnBorrow: func(conn redis.Conn,t time.Time) error{
			if time.Since(t) < time.Minute{
				return nil
			}
			_,err := conn.Do("PING")
			return err
		},
	}
}

func main(){
	c := pool.Get()
	defer c.Close()
	_,err := c.Do("Set","test",100)
	if err != nil{
		fmt.Println(err)
	}
	value,err := redis.Int(c.Do("Get","test"))
	if err != nil{
		fmt.Println("get test failed,",err)
		return
	}
	fmt.Println("test:",value)
}