package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Simple Shell")
	fmt.Println("---------------------")

	for {
		fmt.Println("Next command")
		fmt.Print(": ")
		text, _ := reader.ReadString('\n')
		// convert CRLF to LF
		text = strings.Replace(text, "\n", "", -1)

		if text == "help" || text == "h" {
			fmt.Println(printHelp())
		}

		if strings.Compare("hi", text) == 0 {
			fmt.Println("hello, Yourself")
		}

		if text == "exit" || text == "e" {
			break
		}

		fmt.Println()
	}

}

func printHelp() string {
	return `
Commands:
help -> Prints possible commands
hi   -> Pring hello message
exit -> Exit application
	`
}
