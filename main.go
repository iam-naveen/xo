package main

import (
	"fmt"
)

type Player int

const (
	_ Player = iota
	X
	O
)

type Game struct {
	Players []Player
	Board   [3][3]Player
	move    int
}
func NewGame() *Game {
	game := &Game{
		Players: []Player{X, O},
		Board:   [3][3]Player{},
		move:    1,
	}
	game.PrintBoard()
	return game
}

func (g *Game) PrintBoard() {
	fmt.Print("\033[H\033[2J") // clear screen
	for i := 0; i < 3; i++ {
		if i == 0 {
			fmt.Println("    0   1   2  ")
			fmt.Println("  ┌───┬───┬───┐")
		} else {
			fmt.Println("  ├───┼───┼───┤")
		}
		fmt.Printf("%d ", i)
		for j := 0; j < 3; j++ {
			switch g.Board[i][j] {
			case X:
				fmt.Print("│ X ")
			case O:
				fmt.Print("│ O ")
			default:
				fmt.Print("│   ")
			}
		}
		fmt.Println("│")
	}
	fmt.Println("  └───┴───┴───┘")
}

func (g *Game) MakeMove() bool {
	var x, y int
	for {
		fmt.Scanf("%d %d", &x, &y)
		if g.Board[x][y] == 0 {
			g.Board[x][y], g.move = g.Players[g.move%2], g.move+1
			return true
		}
		g.PrintBoard()
		fmt.Println("Invalid move... try again")
	}
}

func (g *Game) CheckWinner() bool {
	for i := 0; i < 3; i++ {
		if g.Board[i][0] == g.Board[i][1] && g.Board[i][1] == g.Board[i][2] && g.Board[i][0] != 0 {
			fmt.Println("Winner is", g.Board[i][0])
			return true
		}
		if g.Board[0][i] == g.Board[1][i] && g.Board[1][i] == g.Board[2][i] && g.Board[0][i] != 0 {
			fmt.Println("Winner is", g.Board[0][i])
			return true
		}
	}
	if g.Board[0][0] == g.Board[1][1] && g.Board[1][1] == g.Board[2][2] && g.Board[0][0] != 0 {
		fmt.Println("Winner is", g.Board[0][0])
		return true
	}
	if g.Board[0][2] == g.Board[1][1] && g.Board[1][1] == g.Board[2][0] && g.Board[0][2] != 0 {
		fmt.Println("Winner is", g.Board[0][2])
		return true
	}
	return false
}

func main() {
	game := NewGame()
	for game.MakeMove() && !game.CheckWinner() {
		game.PrintBoard()
	}
}
