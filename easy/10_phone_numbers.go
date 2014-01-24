package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	var terminal = bufio.NewReader(os.Stdin)
	var numbers []string = []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}
	var valid_entry int = 0
	var valid_number bool = false
	
	fmt.Println("Welcome to the phonenumber checker.")
	fmt.Printf("Please enter number: ")
	input, err := terminal.ReadString('\n')
	if err != nil {log.Fatalf("%s", err)}
	
	checking_phone := make([]string, len(input))
	
	for i, v := range input {
		checking_phone[i] = string(v)
	}
	switch checking_phone[0] {
	case "(":
		switch len(checking_phone){
		case 14:
			fmt.Println("bloop")
			for i, v := range checking_phone {
				for _, k := range numbers {
					if v == k && i != 4 && i != 8 {valid_entry +=1 }
				}
				if v == ")" && i == 4 {valid_entry += 1}
				if v == "-" && i == 8 {valid_entry += 1}
			}
			if valid_entry == 12 {valid_number = true}
		case 15:
			fmt.Println("bloop")
			for i, v := range checking_phone {
				for _, k := range numbers {
					if v == k && i != 4 && i != 5 && i != 9{valid_entry +=1 }
				}
				if v == ")" && i == 4 {valid_entry += 1}
				if v == " " && i == 5 {valid_entry += 1}
				if v == "-" && i == 9 {valid_entry += 1}
			}
			if valid_entry == 13 {valid_number = true}
		default: valid_number = false
	}
	default:
		switch len(checking_phone) {
		case 11:
			for _, v := range checking_phone {
				for _, k := range numbers {
					if v == k {valid_entry += 1}
				}
			}
			if valid_entry == 10 {valid_number = true}
		case 13:
			for i, v := range checking_phone {
				for _, k := range numbers {
					if v == k && i != 3 && i != 7 {valid_entry += 1}
				}
				if i == 3 && v == "-" {valid_entry += 1}
				if i == 3 && v == "." {valid_entry += 1}
				if i == 7 && v == "-" {valid_entry += 1}
				if i == 7 && v == "." {valid_entry += 1}
				if valid_entry == 12{valid_number = true}
			}
		}
	}
	
	fmt.Println(valid_entry)
	if valid_number == true {fmt.Println("Correct Phonenumber Entry!")} else {fmt.Println("Incorrect Phonenumber Entry!")}	
}
