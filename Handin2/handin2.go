package main

import (
	"fmt"
	"math/rand"

) 

var channel = make(chan data)
var done = make(chan bool)



type data struct{
	seq int
	ack int
	message string
}



func client(){
	x := rand.Intn(10)
	fmt.Println("Client sends seq",x)
	channel <- data{x, 0, " "}
	v:= <-channel
	fmt.Println("Client receives ack",v.ack)
	if v.ack == x+1 {
		channel <- data{v.ack, v.seq+1, " "}
		fmt.Println("Client sends seq",v.ack, "and ack", v.seq+1)
		message := RandomString(10)
		fmt.Println("Client sends message", message, "and seq", v.ack)
		channel <- data{v.ack, 0, message}
	} else{
		fmt.Println("Handshake went wrong")
		done <- true
		
	}
	//Client receives final ack
	<-channel
	
}

func server(){
	y:=rand.Intn(10)
	v := <- channel 
	fmt.Println("Server receives seq", v.seq)
	fmt.Println("Server sends ack", v.seq+1, "and seq", y)
	channel <- data{y, v.seq+1, " "}

	w:=<- channel
	if w.ack == y+1 {
		z := <- channel 
		fmt.Println("Message has been received", z.message)
		fmt.Println("Server sends ack", z.seq)
		channel <- data{z.seq, z.seq, " "}


	} else {
		fmt.Println("Handshake went wrong")
	}

	done <- true
}

func main(){
	go client()
	go server()

	<- done
}

//Random string generator taken from: https://stackoverflow.com/questions/22892120/how-to-generate-a-random-string-of-a-fixed-length-in-go
var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func RandomString(n int) string {
    b := make([]rune, n)
    for i := range b {
        b[i] = letterRunes[rand.Intn(len(letterRunes))]
    }
    return string(b)
}