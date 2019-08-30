package main

import (
	"bufio"
	"fmt"
	"os"
)

func menu() {
	fmt.Printf(`
Welcome to Frosty Mumulala Snowsports rental!
1. List all reservations
2. Show reservation
3. Create new reservation
4. Update reservation
5. Delete reservation
`)
}

func readString(prompt string) string {
	fmt.Printf("%s", prompt)

	reader := bufio.NewReader(os.Stdin)

	var input string
	fmt.Fscan(reader, &input)
	return input
}

func app() {
	sysInput := ""
	for sysInput != "exit" {
		menu()
		sysInput = readString("Your Input: ")
		if sysInput == "1" {
			fmt.Println("You chose 1")
		} else if sysInput == "2" {
			fmt.Println("You chose 2")
		} else {
			sysInput = "exit"
		}
	}
}

func main() {
	app()
}
