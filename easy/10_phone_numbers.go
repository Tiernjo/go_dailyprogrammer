package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	var terminal = bufio.NewReader(os.Stdin)			// Read from terminal
	var numbers []string = []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}
	var valid_entry int = 0
	var valid_number bool = false
	
	fmt.Println("Welcome to the phonenumber checker.")
	fmt.Printf("Please enter number: ")					// Enter phonenumber
	input, err := terminal.ReadString('\n')
	if err != nil {log.Fatalf("%s", err)}
	
	checking_phone := make([]string, len(input))
	
	for i, v := range input {							// Convert string to []string
		checking_phone[i] = string(v)
	}
	switch checking_phone[0] {							// Check against first charcter intered
	case "(":											// First character as "("
		switch len(checking_phone){
		case 14:										// (999)999-9999 format
			fmt.Println("bloop")
			for i, v := range checking_phone {
				for _, k := range numbers {
					if v == k && i != 4 && i != 8 {valid_entry +=1 }
				}
				if v == ")" && i == 4 {valid_entry += 1}	// end parenthesis
				if v == "-" && i == 8 {valid_entry += 1}	// only dash
			}
			if valid_entry == 12 {valid_number = true}
		case 15:										// (999) 999-9999 format
			fmt.Println("bloop")
			for i, v := range checking_phone {
				for _, k := range numbers {
					if v == k && i != 4 && i != 5 && i != 9{valid_entry +=1 }
				}
				if v == ")" && i == 4 {valid_entry += 1}	// end parenthesis
				if v == " " && i == 5 {valid_entry += 1}	// only blank
				if v == "-" && i == 9 {valid_entry += 1}	// only dash
			}
			if valid_entry == 13 {valid_number = true}
		default: valid_number = false
	}
	default:											// first character isnt "("
		switch len(checking_phone) {
		case 11:										// 9999999999 format
			for _, v := range checking_phone {
				for _, k := range numbers {
					if v == k {valid_entry += 1}
				}
			}
			if valid_entry == 10 {valid_number = true}
		case 13:										// 999-999-9999 and 999.999.9999 format
			for i, v := range checking_phone {
				for _, k := range numbers {
					if v == k && i != 3 && i != 7 {valid_entry += 1}
				}
				if i == 3 && v == "-" {valid_entry += 1}	// first dash
				if i == 3 && v == "." {valid_entry += 1}	// first dot
				if i == 7 && v == "-" {valid_entry += 1}	// second dash
				if i == 7 && v == "." {valid_entry += 1}	// second dot
				if valid_entry == 12{valid_number = true}
			}
		}
	}
	
	fmt.Println(valid_entry)
	if valid_number == true {fmt.Println("Correct Phonenumber Entry!")} else {fmt.Println("Incorrect Phonenumber Entry!")}	
}
