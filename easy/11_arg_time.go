package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"
)

func usage() {
	fmt.Fprintf(os.Stderr, "usage: %s %s %s\n", os.Args[0], 
		os.Args[1], os.Args[2])	// define how to input arguments
	flag.PrintDefaults()		// Print defaults into arguments
	os.Exit(2)					// Exit on error
}

func main() {
	flag.Usage = usage			// assign usage
	flag.Parse()				// parse flags
	
	args := flag.Args()			// set args into []string
	if len(args) != 3 {			// always take 3 arguments
		fmt.Println("Need three arguments!")
		os.Exit(1)
	}
	
	run(args)
}

func run(args []string) {
	int64_day, err := strconv.ParseInt(args[0], 10, 0) 		// Convert string to int
	if err != nil {log.Fatalln(err)}
	int64_month, err := strconv.ParseInt(args[1], 10, 64)	// ^^^
	if err != nil {log.Fatalln(err)}
	int64_year, err := strconv.ParseInt(args[2], 10, 0)		// ^^^
	if err != nil {log.Fatalln(err)}
	
	var str_month time.Month								// Convert numerical month to string month
	int_month := int(int64_month)
	str_month = time.Month(int_month)
	int_day := int(int64_day)								// convert int64 to int
	int_year := int(int64_year)								// ^^^
	
	time := time.Date(int_year, str_month, int_day, 0, 0, 0, 0, time.UTC)	// Create time out of data

	fmt.Printf("Day:     %v\nMonth:   %v\nYear:    %v\nWeekday: %v\n", int_day, str_month, int_year, time.Weekday())
}
