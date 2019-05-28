# tic-tac-go

Most definitely not a coding exercise for a job interview.

## Gameplay Basics

A player gets a designated token - 'x' or 'o' - which they can place on the board in any one of the nine spots, so long as it is otherwise unoccupied. Play then passes to the next player, who places their own, opposing token. Each player's objective is to place three of these tokens in a straight line, which can be horizontal, vertical, or diagonal in orientation.

### Rules

1. Play begins with a blank 3x3 board.
2. Play alternates between two players.
3. Players may only place their designated token in a valid, unoccupied square on the board.
4. A placed token cannot be moved.
5. Play ends when a player has won the game, all squares are occupied, or all possible routes to success are blocked.

## Assignment Specifications

Your implementation of the app should meet the following product requirements:

- [ ] The game should have some kind of interface for user(s) to interact. (e.g. in the browser or on the command line)
- [ ] It should enforce the basic rules and determine the end result of a game.
- [ ] The user should be able to start the game
- [ ] The user should be able to take alternate turns to play the game
- [ ] The logic should be able to evaluate the winner
- [ ] The user should be able to stop an ongoing game at any point
- [ ] There should be proper messaging when the game changes the state

Details:

- Use whatever language(s) you prefer.
- Include a short README that describes the features you’ve implemented, what decisions you made and why, and how to run your game.
- Treat this exercise as if you were writing production code.
- Upload your game to GitHub (or source control of your choice), or email us your submission.
