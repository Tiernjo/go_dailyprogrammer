package main

import (
	"fmt"
	"strings"
	"io/ioutil"
	"os"
	"log"
)

var (
	cipher = map[string]string{
		"A": "W",
		"B": "V",
		"C": "U",
		"D": "T",
		"E": "S",
		"F": "F",
		"G": "Q",
		"H": "P",
		"I": "O",
		"J": "N",
		"K": "M",
		"L": "L",
		"M": "K",
		"N": "J",
		"O": "I",
		"P": "H",
		"Q": "G",
		"R": "F",
		"S": "E",
		"T": "D",
		"U": "C",
		"V": "B",
		"W": "A",
		"X": "Z",
		"Y": "Y",
		"Z": "X",
		"a": "w",
		"b": "v",
		"c": "u",
		"d": "t",
		"e": "s",
		"f": "r",
		"g": "q",
		"h": "p",
		"i": "o",
		"j": "n",
		"k": "m",
		"l": "l",
		"m": "k",
		"n": "j",
		"o": "i",
		"p": "h",
		"q": "g",
		"r": "f",
		"s": "e",
		"t": "d",
		"u": "c",
		"v": "b",
		"w": "a",
		"x": "z",
		"y": "y",
		"z": "x",
	}
	input string
)

func main() {
	
	fmt.Println("Welcome to the alphabetical caesar cipher.")
	run()
}

func run() {
	var encrypted []string
	
	// ask how to load string
	for x := 0; x == 0; {
		fmt.Println("What method do you want to use?")
		fmt.Println("     [T]ype in input")
		fmt.Println("     [L]oad file")
		fmt.Println("     [Q]uit")
		fmt.Scan(&input)
		
		switch input {
		case "T":
			encrypted = load_string()
			fmt.Println(strings.Join(encrypted, ""))
		case "L":
			encrypted = load_file()
			fmt.Println(strings.Join(encrypted, ""))
		case "Q":
			x = 1
		}
	}
}

// load string
func load_string() []string {
	fmt.Printf("Enter text: ")
	fmt.Scan(&input)
	
	return encrypt(input)
}

func load_file() []string{
	
	var (
		path, contents string
		output []string
		to_encrypt, to_write bool
	)
	
	for crypt_type := 0; crypt_type == 0; {			// set if encrypt or decrypt
		fmt.Println("Do you want to do?")
		fmt.Println("     [E]ncrypt")
		fmt.Println("     [D]ecrypt")
		fmt.Println("     [Q]uit")
		fmt.Scanln(&input)
		
		switch input {
		case "E":
			to_encrypt = true
			crypt_type = 1
		case "D":
			to_encrypt = false
			crypt_type = 1
		case "Q":
			return nil
		}
	}
	
	for if_write := 0; if_write == 0; {				// set if write results to file
		fmt.Println("Do you want to save results?")
		fmt.Println("     [Y]es")
		fmt.Println("     [N]o")
		fmt.Println("     [Q]uit")
		fmt.Scanln(&input)
		
		switch input {
		case "Y":
			to_write = true
			if_write = 1
		case "N":
			to_write = false
			if_write = 1
		case "Q":
			return nil
		}
	}
	
	fmt.Printf("File location: ")					// set path to file
	fmt.Scan(&path)
	fmt.Println(path)
	
	bs, err := ioutil.ReadFile(path)				// Read file
	if err != nil {
		log.Fatalf("When ioutil.ReadFile(path) is called: %s", err)
	}
	
	contents = string(bs)
	
	switch to_encrypt {								// Encrypt/Decrypt
	case true:	
		output = encrypt(contents)
	case false:
		output = decrypt(contents)
		
	}
	switch to_write {								// Write to file if selected 
	case true:
		fmt.Printf("Enter new file name: ")
		fmt.Scanln(&input)
		file, err := os.Create(input)
		if err != nil {
			log.Fatalf("When trying to create a file: %s", err)
		}
		defer file.Close()
		
		file.WriteString(strings.Join(output, ""))
	case false:		
	}
	
	
	return output
}

func encrypt(input string) []string{
	// return encrypted string of input string length/capacity
	output := make([]string, len(input), len(input))
	
	// iterate through input and get the values out
	for i, v := range input {
		now_string := string(v)
		fmt.Println(now_string)
		// iterate through cipher set and encrypt string
		for key, shift := range cipher {
			if now_string == key {
				output[i] = shift
			} else if cipher[now_string] == ""{
				output[i] = now_string
			}
		}
	}
	
	return output
}

func decrypt(input string) []string {
	output := make([]string, len(input), len(input))
	
	for i, v := range input {
		now_string := string(v)
		for key, shift := range cipher {
			if now_string == shift {
				output[i] = key
			} else if cipher[now_string] == "" {
				output[i] = now_string
			}
		}
	}
	
	return output
}
