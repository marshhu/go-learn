package main

import "fmt"

func send(ch chan int,exitChan chan bool){
	for i := 0; i < 10;i++{
		ch <- i
	}
	close(ch)
	exitChan <- true
}

func recv(ch chan int,exitChan chan bool){
	for{
		v,ok := <- ch
		if !ok{
			break
		}
		fmt.Println(v)
	}
	exitChan <- true
}

func main(){
	ch := make(chan int,10)
	exitChan := make(chan bool,2)
	go send(ch,exitChan)
	go recv(ch,exitChan)

	for i := 0;i < 2;i++{
		<- exitChan
	}
}