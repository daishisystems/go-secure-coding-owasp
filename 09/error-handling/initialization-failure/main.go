package main

import (
	"fmt"
	"log"
)

func initialize(i int) {
	//This is just to deliberately crash the function.
	if i < 2 {
		fmt.Printf("Var %d - initialized\n", i)
	} else {
		//This was never supposed to happen, so we'll terminate our program.
		log.Fatal("Init failure - Terminating.")
	}
}

func main() {
	i := 1
	for i < 3 {
		initialize(i)
		i++
	}
	fmt.Println("Initialized all variables successfully")
}
