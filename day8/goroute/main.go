package main

import(
	"fmt"
	"time"
)

func test(){
	var i int
	for{
		fmt.Println(i)
		time.Sleep(time.Second)
		i++
	}
}

func main(){
	go test()
	for{
		fmt.Println("main running")
		time.Sleep(time.Second)
	}
}