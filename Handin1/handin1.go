package main

import(
	"fmt"
	"time"
) 

var channels = make([]chan bool, 10, 10)

func freud(index int){
	var hunger int = 3

	for hunger > 0 {
		think(index)
		
	}
	
	eat(index, &hunger)

}

func fork(index int){

}

func eat(index int, hunger *int){
	
	*hunger = *hunger - 1
	fmt.Println("Freud number", index, "is eating")
}

func think(index int){
	time.Sleep(1000 * time.Millisecond)
	fmt.Println("Freud number", index, "is thinking")
}

func main(){
	fmt.Println("Hello")

	for i := 0; i < 5; i++ {
		go fork(i)
	}

	for i := 0; i < 5; i++ {
		go freud(i)
	}
}