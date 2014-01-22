/* http://www.reddit.com/r/dailyprogrammer/comments/pnhyn/2122012_challenge_5_easy/
 * 
 * Your challenge for today is to create a program which is password protected, 
 * and wont open unless the correct user and password is given. 
 * 
 * For extra credit, have the user and password in a seperate .txt file.
 * 
 * for even more extra credit, break into your own program :)
 * 
 */
package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

func main() {
	var username, password string
	fmt.Printf("What is your username: ")
	fmt.Scanln(&username)
	fmt.Printf("What is your password: ")
	fmt.Scanln(&password)
	
	userpass := load_password()
	
	switch {
	case username == userpass[0]:
		fallthrough
	case password == userpass[1]:
		load_file()
	default:
		fmt.Println("Fuck off")
	}
}

func load_password() []string{
	var (
		contents string
		is_pass bool = false
	)
	userpass := make([]string, 2, 2)
	
	file, err := ioutil.ReadFile("example/userpass.txt")
	if err != nil {
		log.Fatalf("When trying to fetch userpass.txt: %s", err)
	}
	contents = string(file)
	
	for _, v := range contents {
		if string(v) == " " {
			is_pass = true
			continue
		}
		
		if is_pass == false {
			userpass[0] += string(v)
		} else {
			userpass[1] += string(v)
		}
	}
	return userpass
}

func load_file() {
	file, err := ioutil.ReadFile("example/lorem_ipsum.txt")
	if err != nil {
		log.Fatalf("When reading unlocked file: %s", err)
	}
	
	content := string(file)
	fmt.Println()
	fmt.Println(content)
}
