package main

import (
	"fmt"
	"log"
	"sort"
)
var (
	input string
)

func main() {
	var is_num bool
	fmt.Println("Welcome to the alpha-numeric sorter.")
	
	for x := 0; x == 0; {								// Main Menu
		fmt.Println("Do you want to sort Numbers or Words?")
		fmt.Println("     [N]umbers")
		fmt.Println("     [W]ords")
		fmt.Println("     [Q]uit")
		fmt.Printf("Enter selection: ")
		_, err := fmt.Scanln(&input)
		if err != nil {
			log.Fatalf("%s", err)
		}

		switch input {									// Main Menu selection
		case "W": fallthrough
		case "w": fallthrough
		case "words": fallthrough
		case "Words":
			is_num = false
			data_entry(is_num)
		case "N": fallthrough
		case "n": fallthrough
		case "numbers": fallthrough
		case "Numbers":
			is_num = true
			data_entry(is_num)
		case "Q": fallthrough
		case "q": fallthrough
		case "quit": fallthrough
		case "Quit": 
			x = 1
		default:
			fmt.Println("\nIncorrect submission!\n")
		}
	}
}

func data_entry(is_num bool) {
	var (
		str_input, insert_type string
		int_input, sets int
	)
	
	switch is_num {
	case true: insert_type = "Number"
	case false: insert_type = "Word"
	}
	
	fmt.Printf("How many sets of do you want to do: ")	// How many numbers/words to allocate
	_, err := fmt.Scanln(&sets)
	if err != nil {log.Fatalf("%", err)}
	int_slice := make([]int, sets)
	str_slice := make([]string, sets)
	
	for x := 0; x < sets; x++{							// Enter numbers/words
		fmt.Printf("Please insert %v ([-1] to quit): ", insert_type)
		switch is_num {
		case false:										// Words
			_, err := fmt.Scanln(&str_input)
			if err != nil {log.Fatalf("%s", err)}
			str_slice[x] = str_input
		case true:										// Numbers
			_, err := fmt.Scanln(&int_input)
			if err != nil {log.Fatalf("%s", err)}
			int_slice[x] = int_input
		}
		if int_input == -1 || str_input == "-1"{x = sets}
	}
	sort_num(int_slice)
	sort_str(str_slice)
}

func sort_num(int_slice []int) {
	fmt.Printf("Original Numbers: %v\n", int_slice)
	var sorting_slice sort.IntSlice = int_slice
	sort.Sort(sorting_slice)
	fmt.Printf("Sorted Numbers: %v\n", sorting_slice)
}

func sort_str(str_slice []string) {
	fmt.Printf("Original Words: %v\n", str_slice)
	var sorting_slice sort.StringSlice = str_slice
	sort.Sort(sorting_slice)
	fmt.Printf("Sorted Words: %v\n", sorting_slice)
}

