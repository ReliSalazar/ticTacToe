package ttt

import "fmt"

func NewGame(p1 Player, p2 Player) *Game {
	return &Game{
		p1:      p1,
		p2:      p2,
		current: p1,
		board:   newBoard(),
		round:   1,
	}
}

// Game keeps track of the progress of a tic tac toe game.
type Game struct {
	p1      Player
	p2      Player
	current Player
	board   *board
	round   int
}

// IsOver checks if a game is over.
func (g *Game) isOver() bool {
	return g.board.winner() != "" || g.board.emptyCount() == 0
}

// Start will start a game.
func (g *Game) Start() {
	fmt.Println("___Welcome to Tic Tac Toe___")
	for !g.isOver() {
		g.printInfo()
		i, j, err := g.current.GetMove(g.board)

		if err != nil {
			fmt.Println("your input is invalid, please try again")
			continue
		}

		if i > 2 || j > 2 {
			fmt.Println("your input is invalid, please try again")
			continue
		}

		if g.board[i][j] != "_" {
			fmt.Println("Position is alredy occupied, please try again")
			continue
		}

		g.board[i][j] = g.current.Mark()
		g.switchPlayer()
		g.round++
	}

	fmt.Println(g.board)
	fmt.Println("Game over!")
}

func (g *Game) switchPlayer() {
	if g.current == g.p1 {
		g.current = g.p2
	} else {
		g.current = g.p1
	}
}

func (g *Game) printInfo() {
	fmt.Printf("Turn #%d\n", g.round)
	fmt.Println(g.board)
	fmt.Println("Current player:", g.current.Name())
}
