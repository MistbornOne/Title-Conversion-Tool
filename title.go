package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	// Define flags
	inPlace := flag.Bool("i", false, "Edit file in place (instead of writing to stdout)")
	bold := flag.Bool("b", false, "Format as bold (**text**)")
	heading := flag.Int("h", 0, "Format as heading (1-6, e.g. -h 2 for ##)")
	flag.Parse()

	// Handle input source
	var input io.Reader
	var filePath string

	args := flag.Args()
	if len(args) > 0 {
		filePath = args[0]
		file, err := os.Open(filePath)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error opening file: %v\n", err)
			os.Exit(1)
		}
		defer file.Close()
		input = file
	} else {
		// No file argument, read from stdin
		input = os.Stdin
	}

	// Read input and convert to title case
	scanner := bufio.NewScanner(input)
	var outputLines []string

// If you don't want your title to be bolded, use the option below	
	for scanner.Scan() {
		line := scanner.Text()
		titleCaseLine := strings.Title(line)

// If you want the title to be bold every time, then use this option
	//for scanner.Scan() {
		//line := scanner.Text()
		//titleCaseLine := "**" + strings.Title(line) + "**"

		// Apply markdown formatting if requested
		if *bold {
			titleCaseLine = "**" + titleCaseLine + "**"
		}
		
		if *heading > 0 && *heading <= 6 {
			prefix := strings.Repeat("#", *heading) + " "
			titleCaseLine = prefix + titleCaseLine
		}
		
		outputLines = append(outputLines, titleCaseLine)
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "Error reading input: %v\n", err)
		os.Exit(1)
	}

	outputText := strings.Join(outputLines, "\n")

	// Handle output destination
	if *inPlace && filePath != "" {
		err := os.WriteFile(filePath, []byte(outputText), 0644)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error writing to file: %v\n", err)
			os.Exit(1)
		}
	} else {
		fmt.Println(outputText)
	}
}
