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
	_, err = conn.Do("hset", "myhash", "bike1", "mobike")
	if err != nil {
		fmt.Println("hset failed", err.Error())
	}
	value,err := redis.String(conn.Do("hget", "myhash", "bike1"))
	if err != nil{
		fmt.Println("hget failed,",err)
		return
	}
	fmt.Println(value)

	_, err = conn.Do("hmset", "myhash", "bike2", "bluegogo", "bike3", "xiaoming", "bike4", "xiaolan")
	if err != nil {
		fmt.Println("hmset error", err.Error())
	}

	list, err := redis.Strings(conn.Do("hmget", "myhash", "bike1", "bike2", "bike3", "bike4"))
	fmt.Println("hmget result:")
	for _,v := range list{
		fmt.Println(v)
	}

	isExist, err := conn.Do("hexists", "myhash", "tmpnum")
	if err != nil {
		fmt.Println("hexist failed", err.Error())
	} else {
		fmt.Println("exist or not:", isExist)
	}

	resKeys, err := redis.Strings(conn.Do("hkeys", "myhash"))
	if err != nil {
		fmt.Println("hkeys failed", err.Error())
	}
	fmt.Println("hkeys result:")
	for _,v := range resKeys{
		fmt.Println(v)
	}

	resValues, err := redis.Strings(conn.Do("hvals", "myhash"))
	if err != nil {
		fmt.Println("hvals failed", err.Error())
	}
	fmt.Println("hvals result:")
	for _,v := range resValues{
		fmt.Println(v)
	}

	_, err = conn.Do("HDEL", "myhash", "tmpnum")
	if err != nil {
		fmt.Println("hdel failed", err.Error())
	}

	result, err := redis.Strings(conn.Do("hgetall", "myhash"))
	if err != nil {
		fmt.Println("hgetall failed", err.Error())
	}
	fmt.Println("hgetall result:")
	for _,v := range result{
		fmt.Println(v)
	}
}