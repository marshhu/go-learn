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
	_,err = conn.Do("sadd", "myset", "mobike", "foo", "ofo", "bluegogo")
	if err != nil{
		fmt.Println("sadd failed:",err)
	}
	value, err := redis.Strings(conn.Do("smembers", "myset"))
	if err != nil{
		fmt.Println("get myset failed,",err)
		return
	}
	for _,v := range value{
		fmt.Println(v)
	}

	isMember, err := conn.Do("sismember", "myset", "foo")
	if err != nil {
		fmt.Println("sismember get failed", err.Error())
	} else {
		fmt.Println("foo is or not myset's member:", isMember)
	}

	_, err = conn.Do("sadd", "dbset", "foo", "ofo", "xiaolan")
	if err != nil {
		fmt.Println("set add failed", err.Error())
	}
	inner, _ := redis.Strings(conn.Do("sinter", "myset", "dbset"))
	fmt.Println("myset and dbset inner result:")
	for _,v := range inner{
		fmt.Println(v)
	}

	union, err := redis.Strings(conn.Do("sunion", "myset", "dbset"))
	fmt.Println("myset and dbset union result:")
	for _,v := range union{
		fmt.Println(v)
	}

	diff, err := redis.Strings(conn.Do("sdiff", "dbset", "myset"))
	fmt.Println("myset and dbset diff result:")
	for _,v := range diff{
		fmt.Println(v)
	}
}