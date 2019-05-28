package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/kbrimm/tic-tac-toe/controllers"
	"github.com/kbrimm/tic-tac-toe/models"
)

func main() {
	fmt.Println("                     ===== Tic Tac Toe =====")
	fmt.Println("\nText-based tic-tac-toe, because the console never goes out of style.")
	fmt.Println("               ==================================")
	fmt.Println("You can always type 'new' to start a new game or 'exit' to quit.")
	reader := bufio.NewReader(os.Stdin)
	for {
		for {
			var game models.Game
			fmt.Println("\nPlayer 1: 'x' or 'o'?")
			input, _ := reader.ReadString('\n')
			input, new, exit := parseInput(input)

			if new {
				fmt.Println("Starting a new game...")
				break
			} else if exit {
				fmt.Println("Goodbye!")
				os.Exit(0)
			}

			markChoice := []rune(input)[0]
			if markChoice == 'x' || markChoice == 'o' {
				game = controllers.StartNewGame(markChoice)
			} else {
				break
			}

			fmt.Printf("%+v\n", game)
			fmt.Print("Where will you make your mark? (ex. A3): ")
			for {
				input, _ = reader.ReadString('\n')
				input, new, exit = parseInput(input)
				if new {
					fmt.Println("Starting a new game...")
					break
				} else if exit {
					fmt.Println("Goodbye!")
					os.Exit(0)
				}
				row, col, err := getCoordinates(input)
				if err != nil {
					fmt.Println(err)
					continue
				}
				game, err = controllers.PlaceMark(game, row, col)
				fmt.Printf("%+v\n", game)
				if err != nil {
					fmt.Println(err)
				}
				winner, win, stalemate := controllers.ScoreGame(game)
				if win {
					fmt.Println(fmt.Sprintf("Congratulations! Player '%c' wins!", winner))
					break
				} else if stalemate {
					fmt.Println("Stalemate!")
					break
				}
				fmt.Print("Where will you make your mark? (ex. A3): ")
			}
		}
	}
}

func parseInput(input string) (string, bool, bool) {
	input = strings.TrimRight(input, "\r\n")
	if input == "new" {
		return "", true, false
	} else if input == "exit" {
		return "", false, true
	} else {
		return strings.ToLower(input), false, false
	}
}

func getCoordinates(str string) (row int, col int, err error) {
	validCoords, _ := regexp.Compile(`^[abc][1-3]$`)
	if validCoords.MatchString(str) {
		var rowInt, colInt int
		rowRune := []rune(str)[0]
		switch rowRune {
		case 'a':
			rowInt = 0
		case 'b':
			rowInt = 1
		case 'c':
			rowInt = 2
		}

		colRune := []rune(str)[1]
		switch colRune {
		case '1':
			colInt = 0
		case '2':
			colInt = 1
		case '3':
			colInt = 2
		}
		return rowInt, colInt, nil
	}
	return 0, 0, errors.New("input error: invalid grid coordinates")
}
