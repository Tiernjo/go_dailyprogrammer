/* http://www.reddit.com/r/dailyprogrammer/comments/pm6oj/2122012_challenge_4_easy/
 * 
 * You're challenge for today is to create a random password generator!
 * For extra credit, allow the user to specify the amount of passwords to generate.
 * For even more extra credit, allow the user to specify the length of the strings he wants to generate!
 * 
 */
package main

import (
	"fmt"
	"strings"
	"math/rand"
	"time"
	"os"
)

// List of password characters
var password_slice = []string {"0", "1", "2", "3", "4", "5", "6", "7", "8", "9",
	"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z",
	"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z",
	"!", "@", "#", "$", "%", "^", "&", "*", "~", "<", ">"}

func main() {
	var (
		password_num, password_length int
		random_character int
		list_counter int = 0
	)
	
	// print all usable characters
	fmt.Println(strings.Join(password_slice, ""))
	
	// get how many passwords to make
	fmt.Println("Welcome to the friendly password creator.")
	fmt.Printf("How many passwords do you want to make: ")
	fmt.Scanln(&password_num)
	// slice of all passwords
	combine_password := make([]string, password_num, password_num)
	
	// seed random number generator
	rand.Seed(time.Now().UTC().UnixNano())
	
	// find length of all passwords
	fmt.Println("How many characters in the password: ")
	fmt.Scanln(&password_length)
	password_output := make([]string, password_length, password_length)
	
	// run number of passwords times
	for i := 0; i < password_num; i++ {
		//	run length of password
		for i, _ := range password_output {
			// get a random character out of password characters
			// set each random character into password
			random_character = rand.Intn(len(password_slice))
			password_output[i] = password_slice[random_character]
		}
		
		// slice of passwords
		combine_password[list_counter] = strings.Join(password_output, "")
		list_counter++
	}
	
	print(combine_password)
}


func print(passwords []string) {
	file, err := os.Create("example/passwords.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	
	// create password file, newline after each one
	file.WriteString(strings.Join(passwords, "\n"))
	fmt.Println("Done!")
}
