# tic-tac-toe

## Features

* Slick console-based interface for optimum efficiency. Includes helpful, relevant prompts and a skillful ASCII representation of the game board!
* First player gets to choose whether to be 'x' or 'o', and the second player just has to live with that decision!
* Each player can place their respective marks in any valid (non-occupied, on-board) position and will be chastized if they fail to meet that requirement!
* Automagically determines if a game has been won, or if it has fallen into a terminal stalemate!

## Decisions

The biggest decision and concession I made was in opting for a console-based interface. As much as I would have loved to put out a super shiny web interface, the time it would have taken me to come up to speed on a front end would have completely exhausted the requirement for the exercise. A fully-featured console application seemed more appropriate than a middling Javascript hack.

Other choices were somewhat more trivial. These included tech (Go, because I like it), data structures (2D arrays, because they make sense), and deploy methods (Docker, because it's easy).

## How To Run

The most convenient way to run this tic tac toe implementation is through Docker. If you have Docker configured on your system, you can launch straight into this riveting game by building and running a container:

```bash
docker build -t tic-tac-toe .
docker run -it tic-tac-toe
```

This will launch you into a pre-built environment without having to fuss with configuring your GoLang environment.

If you _do_ want to fuss with configuring your GoLang environment, [golang.org](https://golang.org/doc/install) has an excellent guide. If you're running GoLang locally, you can access the application with:

```bash
go get
go run main.go
```

You can also watch all the unit tests pass with `go test ./...`

## Assignment Specifications

Your implementation of the app should meet the following product requirements:

- [x] The game should have some kind of interface for user(s) to interact. (e.g. in the browser or on the command line)
- [x] It should enforce the basic rules and determine the end result of a game.
- [x] The user should be able to start the game
- [x] The user should be able to take alternate turns to play the game
- [x] The logic should be able to evaluate the winner
- [x] The user should be able to stop an ongoing game at any point
- [x] There should be proper messaging when the game changes the state

Details:

* Use whatever language(s) you prefer.
* Include a short README that describes the features you’ve implemented, what decisions you made and why, and how to run your game.
* Treat this exercise as if you were writing production code.
* Upload your game to GitHub (or source control of your choice), or email us your submission.

## Assumptions Made

A player gets a designated token - 'x' or 'o' - which they can place on the board in any one of the nine spots, so long as it is otherwise unoccupied. Play then passes to the next player who places their own, opposing token. Each player's objective is to place three of these tokens in a straight line, which can be horizontal, vertical, or diagonal in orientation.

### Rules

1. Play begins with a blank 3x3 board.
2. Play alternates between two players.
3. Players may only place their designated token in a valid, unoccupied square on the board.
4. A placed token cannot be moved.
5. Play ends when a player has won the game, all squares are occupied, or all possible routes to success are blocked.
