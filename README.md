# Title Converion Tool
## Change lowercase strings to Title format with a single command in the terminal or a keybinding in nvim!

### Demo in nvim:
https://github.com/user-attachments/assets/aa683ad6-a3bd-4a59-b868-48dc2e2419d6

---

### Why Though? 🧠
As someone who writes a lot in Neovim, I wanted a way to quickly toggle a lower case string to a title format (first letter of each word in the string is capitalized).  This is useful for writing notes in markdown files for Obsidian, but also works for commenting in code files.  

This is my first script written in Go.  I wanted to play around with that language and had fun learning a little bit of it's syntax and styling.

## Installation

**Step 1: Navigate to or create the directory where you want the code to live:**

```Bash
mkdir -p [YOURFOLDERNAME]
```
This should create necessary parent directories.

or

```Bash
cd ~/your/filepath/here
```
This option takes you to the directory that already exists.

**Step 2: Create the file for the Go code:**

```Bash
nano title.go
```
or for nvim users:

```Bash
nvim title.go
```

**Step 3: Paste the following code into your file:**

```Go
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

```

**Step 4: Save and exit the file, then run the following code in the root directory:**

```Bash
go build -o title title.go
```
This tells Go to build a new program called 'title' using the title.go file.


**Step 5: Make the program executable:**

```Bash
chmod +x title
```

**Step 6: Check that the program is working in the terminal by executing this:**

```Bash
echo "hello world" | /Your/File/Path/title
```
You should get a print out of: "Hello World"


**Step 7: Add the keybindings to your keybinding.lua or init.lua file in your nvim config files:**
Note, you will need to replace my filepath with your own filepath where you stored the _program_ (not the title.go file)

```Lua
-- title program

-- For normal mode: convert current line
vim.api.nvim_set_keymap(
	"n",
	"<leader>ti",
	":.!'/Users/ianwatkins/dev/github/MistbornOne/projects/programs/title'<CR>",
	{ noremap = true, silent = true }
)

-- Bold title case on current line
vim.api.nvim_set_keymap(
	"n",
	"<leader>tb",
	":.!sh -c '\"/Users/ianwatkins/dev/github/MistbornOne/projects/programs/title\" -b'<CR>",
	{ noremap = true, silent = true }
)

-- H2 heading title case on current line
vim.api.nvim_set_keymap(
	"n",
	"<leader>th",
	":.!sh -c '\"/Users/ianwatkins/dev/github/MistbornOne/projects/programs/title\" -h=2'<CR>",
	{ noremap = true, silent = true }
)

-- For visual mode: convert selected text
vim.api.nvim_set_keymap(
	"v",
	"<leader>ti",
	":!'/Users/ianwatkins/dev/github/MistbornOne/projects/programs/title'<CR>",
	{ noremap = true, silent = true }
```

The above keybinds can be tweaked to your liking.  I personally like my <leader> to be the space bar.  
Then the keystrokes after that follow a logical pattern:  

In Normal Mode:
"<leader>ti" is for standard title conversion for the line your cursor is on.

"<leader>tb" is for bold title conversion for the line your cursor is on.

"<leader>th" is for heading title conversion for the line your cursor is on.

In Visual Mode:
"<leader>ti" is for standard title conversion for the text you've highlighted.


**Step 8: Test the program in nvim!  Navigate to a file in nvim and in normal mode, place your cursor on the line to be tested then try one or more of the keybindings.

---

## Thank You 🙏🏼

I hope this little program is useful to anyone who uses it.  It was fun to build and I'm happy to share it.  Considering at the time of writing I've been coding for less than two months, I feel pretty proud of this little guy.
