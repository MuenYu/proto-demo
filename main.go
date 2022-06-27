package main

import (
	"os"
	"os/signal"
)

func main() {
	sign := make(chan os.Signal)
	signal.Notify(sign, os.Interrupt)
	go server()
	go client()
	<-sign
}

func server() {

}

func client() {

}
