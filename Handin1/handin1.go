package main

import(
	"fmt"
	"time"
	"math/rand"
	
) 

// channels between forks and Freuds
var ready = make([]chan bool, 5)
var finished = make([]chan bool, 5)

// channel used to communicate that Freuds are satisfied
var satisfied = make(chan bool)

func freud(index int){
	hunger := 3

	for hunger > 0 {
		think(index)

		left := false
		right := false

		select{ // check if left fork is ready
		case left = <- ready[index]:
		default: left = false
		}

		if !left { // if left fork is not ready, start over by thinking
			continue
		}

		select{
		case right = <- ready[mod(index-1,5)]:
		default: right = false
		}

		if !right { // Avoid deadlock by letting go of left fork, if right fork is not ready
			finished[index] <- true
			continue
		}

		if left && right {
			eat(index, &hunger)
			finished[index] <- true // let go of forks after eating
			finished[mod(index-1,5)] <- true
		}
	}
	satisfied <- true //When Freud has eaten 3 times, let table() know they are done
	

}

func fork(index int){
	for{
		ready[index] <- true

		// When fork is finished being used, loop will start over and fork will be ready again
		<- finished[index] 
	}
}

func eat(index int, hunger *int){ // avoid deadlock by randomizing amount of time each Freud spends eating
	fmt.Println("Freud number", index, "is eating")
	time.Sleep(time.Duration(rand.Intn(2000)) * time.Millisecond)
	*hunger = *hunger - 1
}

func think(index int){ // avoid deadlock by randomizing amount of time each Freud spends thinking
	thinktime := time.Duration(rand.Intn(1000)) * time.Millisecond
	fmt.Println("Freud number", index, "is thinking")
	time.Sleep(thinktime)
}

func table(){
	for i := 0; i < 5; i++ {
		ready[i] = make(chan bool)
		finished[i] = make(chan bool)
		go fork(i)
		go freud(i)
	}
	
	for i := 0; i < 5; i++{ // keep table() running until every Freud is satisfied
		<- satisfied
	}
	
}

func main(){
	table()
	
}

func mod(a, b int) int { // Go language does not support modulus with negative numbers, so we have made our own
    return (a % b + b) % b
}