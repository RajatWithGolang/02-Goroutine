package main

import (
	"fmt"
	"time"
)

func hello(){
	time.Sleep(2 * time.Second)
	
	fmt.Println("Hello")
}

func main(){
	fmt.Println("Main Started")
	fmt.Println("Calling Hello")
	go hello()
	fmt.Println("Main is Sleeping")
	time.Sleep(1 * time.Second)
	fmt.Println("Main has awaken Terminated")
	
}