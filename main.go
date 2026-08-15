package main

import (
	"fmt"
	"bufio"
	"os"
)

func main()  {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		input := scanner.Text()
		wordList := cleanInput(input)
		fmt.Printf("Your command was: %s\n", wordList[0])
	}
}
