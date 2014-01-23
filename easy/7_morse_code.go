// http://www.reddit.com/r/dailyprogrammer/comments/pr2xr/2152012_challenge_7_easy/
package main

import (
	"fmt"
	"bufio"
	"os"
	"log"
	"strings"
)

// Create conversion table
var (
	morse_code = map[string]string{
		"a": ".-",
		"b": "-...",
		"c": "-.-.",
		"d": "-..",
		"e": ".",
		"f": "..-.",
		"g": "--.",
		"h": "....",
		"i": "..",
		"j": ".---",
		"k": "-.-",
		"l": ".-..",
		"m": "--",
		"n": "-.",
		"o": "---",
		"p": ".--.",
		"q": "--.-",
		"r": ".-.",
		"s": "...",
		"t": "-",
		"u": "..-",
		"v": "...-",
		"w": ".--",
		"x": "-..-",
		"y": "-.--",
		"z": "--..",
		"1": ".----",
		"2": "..---",
		"3": "...--",
		"4": "....-",
		"5": ".....",
		"6": "-....",
		"7": "--...",
		"8": "---..",
		"9": "----.",
		"0": "-----",
		" ": "/",
	}
)

func main() {
	var input string
	var to_morse bool
	fmt.Println("Welcome to the morse code translator.")
	
	for x := 0; x == 0;{
		fmt.Println("Start with English or Morse Code?")		// Menu options
		fmt.Println("     [E]nglish")
		fmt.Println("     [M]orse Code")
		fmt.Println("     [Q]uit")
		fmt.Printf("Enter text: ")
		_, err := fmt.Scanln(&input)					// Menu selection
		if err != nil {
			log.Fatalf("%s", err)
		}

		
		switch input {
		case "E":
			to_morse = true
			input = text_entry()
			convert(to_morse, input)
		case "M":
			to_morse = false
			input = text_entry()
			convert(to_morse, input)
		case "Q":
			fmt.Print("\a")
			x = 1
		}
	}
}

func text_entry() string {
	terminal := bufio.NewReader(os.Stdin)					// Create input reader of standard input type
	
	fmt.Printf("Enter text: ")
	input, err := terminal.ReadString('\n')					// Read everything typed to console
	if err != nil {
		log.Fatalf("%s", err)
	}
	
	return input
}

func convert(to_morse bool, code string) {
	var (
		now_string string
		code_slice []string
	)
	converted := make([]string, len(code))
	
	if to_morse == false {code_slice = morse_format(code)} else {code_slice = english_format(code)}

	for i, v := range code_slice {
		now_string = string(v)
		for key, value := range morse_code {
			
			switch to_morse {					// convert english <=> morse code
			case true:
				if now_string == key {converted[i] = value}	// english to morse code
			case false:	
				if now_string == value {converted[i] = key}	// morse code to english
			}
		}	
	}
	
	fmt.Println(strings.Join(converted, ""))
}

func morse_format(code string) []string{
	morse_slice := make([]string, len(code))
	morse_counter := 0
	temp_morse := ""
	
	for _, value := range code {
		now_string := string(value)
		morse_action := ""
		
		if now_string == " " || now_string == "/"{			// check if end of letter or word
			switch now_string {					// Set action for end of letter or word
			case " ": morse_action = " "
			case "/": morse_action = "/"
			default:  morse_action = ""
			}
			temp_morse = ""						// Clean out temp_morse
			morse_counter += 1
			morse_slice[morse_counter] = morse_action		// write in space or /
			morse_counter += 1
			continue
		}

		
		temp_morse += now_string					// fill out temp_morse
		morse_slice[morse_counter] = temp_morse				// fill morse_slice with temp_morse
	}
	return morse_slice
}

func english_format(code string) []string {
	english_slice := make([]string, len(code))
	
	for i, value := range code {
		english_slice[i] = string(value)				// Convert string to []string
	}
	
	return english_slice
}
