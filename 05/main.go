package main

import(
	"fmt"
	"time"
)

func main(){
	fmt.Println("main started")
	 
	go func(){
		fmt.Println("Hello World")
	}()
	time.Sleep(1* time.Millisecond)
	fmt.Println("main Terminated")
}