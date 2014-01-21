// http://www.reddit.com/r/dailyprogrammer/comments/pih8x/easy_challenge_1/
package main

import (
	"fmt"
)

func main() {
	var (
		name, username string
		age int
	)
	
	fmt.Print("What is your name: ")
	fmt.Scanln(&name)
	
	fmt.Print("How old are you: ")
	fmt.Scanln(&age)
	
	fmt.Print("What is your reddit username: ")
	fmt.Scanln(&username)
	
	fmt.Println("Your name is ", name, ", you are ", age, " years old, and your username is ", username)
}
