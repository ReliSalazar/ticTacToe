package main

import "./ttt"

func main() {
	hp := ttt.NewHumanPlayer("Reli", "X")
	cp := ttt.NewComputerPlayer("Lain", "O")
	g := ttt.NewGame(hp, cp)
	g.Start()
}
