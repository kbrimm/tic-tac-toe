package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/kbrimm/tic-tac-toe/controllers"
	"github.com/kbrimm/tic-tac-toe/models"
)

func main() {
	fmt.Println("===== Tic Tac Toe =====")
	fmt.Println("This interface is fairly naive. It assumes you will supply")
	fmt.Println("reasonable input, and doesn't do too much error checking.")
	fmt.Println("            ==================================            ")
	fmt.Println("You can always type 'new' to start a new game or 'exit' to quit.")
	reader := bufio.NewReader(os.Stdin)
	for {
		for {
			var game models.Game
			fmt.Println("Player 1: 'x' or 'o'?")
			input, _ := reader.ReadString('\n')
			response, new, exit := readInput(input)
			if new {
				fmt.Println("Starting a new game...")
				break
			} else if exit {
				fmt.Println("Goodbye!")
				os.Exit(0)
			}
			if response == 'x' || response == 'o' {
				game = controllers.StartNewGame(response)
			}
			fmt.Printf("%+v\n", game)
			game, _ = controllers.PlaceMark(game, 1, 1)
			fmt.Printf("%+v\n", game)
		}
	}
}

func readInput(input string) (rune, bool, bool) {
	input = strings.TrimRight(input, "\r\n")
	if input == "new" {
		return ' ', true, false
	} else if input == "exit" {
		return ' ', false, true
	} else {
		return []rune(strings.ToLower(input))[0], false, false
	}
}

func getFirstChar(str string) rune {
	return []rune(str)[0]
}
