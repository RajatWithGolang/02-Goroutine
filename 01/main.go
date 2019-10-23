package main

import "fmt"

func hello(){
	fmt.Println("Hello")
}

func main(){
	fmt.Println("Main Started")
	fmt.Println("Calling Hello")
	go hello()   // Main launches hello goroutine but doesn't wait for that
	fmt.Println("Hello Printed")
	fmt.Println("Main Terminated")

}