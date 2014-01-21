// http://www.reddit.com/r/dailyprogrammer/comments/pjbj8/easy_challenge_2/
package main

import (
	"fmt"
)

func main() {
	
	var input string
	fmt.Println("Welcome to your friendly neighborhood calculator.")
	for x := 0; x == 0; {
		fmt.Println("What kind of calculator do you need?")
		fmt.Println("     [P]hysics")
		fmt.Println("     [Q]uit")
		fmt.Scanln(&input)
		
		switch {
		case input == "P":
			physics()
		case input == "Q":
			x = 1		
		}
	}
}

func physics() {
	var ( 
		input string
		force, acceleration, momentum float64
	)
		
	for x := 0; x == 0; {
		fmt.Println("Which do you need to find?")
		fmt.Println("     [F]orce")
		fmt.Println("     [A]cceleration")
		fmt.Println("     [M]omentum")
		fmt.Println("     [Q]uit")
		fmt.Scanln(&input)
		
		switch {
		case input == "F":
			fmt.Println("What is the Accleration: ")
			fmt.Scanln(&acceleration)
			fmt.Println("What is the Momentum: ")
			fmt.Scanln(&momentum)
			
			force = momentum * acceleration
			fmt.Println("The Force is: ", force)
		case input == "A":
			fmt.Println("What is the Force: ")
			fmt.Scanln(&force)
			fmt.Println("What is the Momentum: ")
			fmt.Scanln(&momentum)
			
			acceleration = force/momentum
			fmt.Println("The Acceleration is: ", acceleration)
		case input == "M":
			fmt.Println("What is the Force: ")
			fmt.Scanln(&force)
			fmt.Println("What is the Accleration: ")
			fmt.Scanln(&acceleration)
			
			momentum = force/acceleration
			fmt.Println("The Momentum is: ", momentum)
		case input == "Q":
			x = 1
		default:
		}
	}
}
