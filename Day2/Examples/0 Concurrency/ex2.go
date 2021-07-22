package main

import "fmt"

func main() {

	messages := make(chan string)

	go func() {
		messages <- "ping"
		messages <- "pin2"
	}()

	msg := <-messages
	fmt.Println(msg)
	//msg = <-messages
	//fmt.Println(msg)
	//fmt.Println("test")
}