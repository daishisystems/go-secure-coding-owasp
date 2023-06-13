package main

import "fmt"

func main() {
	start()
	fmt.Println("Returned normally from start().")
}

func start() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in start()")
		}
	}()
	fmt.Println("Called start()")
	part2(0)
	fmt.Println("Returned normally from part2().")
}

func part2(i int) {
	if i > 0 {
		fmt.Println("Panicking in part2()!")
		panic(fmt.Sprintf("%v", i))
	}
	defer fmt.Println("Defer in part2()")
	fmt.Println("Executing part2()")
	part2(i + 1)
}
