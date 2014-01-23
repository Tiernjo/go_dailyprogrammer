package main

import(
	"fmt"
)

func main() {
	fmt.Scanln()
	for i := 99; i > 0; i-- {
		fmt.Print(i, " bottles of beer on the wall. ")
		fmt.Print(i, " bottles of beer. ")
		fmt.Print("You take one down, you pass it around. ")
		fmt.Print((i - 1), " bottles of beer on the wall. ")
	}
}
