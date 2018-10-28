package main

import (
	"fmt"
	"math/rand"
	"time"
)

const maxPoints = 2048
const cols = 4
const rows = 4
const up = 'w'
const down = 'w'
const left = 'w'
const right = 'w'

type model struct {
	Score int
	Board [cols][rows]int
}

func (m *model) getScore() int {
	return m.Score
}

func (m *model) getBoard() [cols][rows]int {
	return m.Board
}

func (m *model) printBoard() {
	for i := 0; i < rows; i++ {
		for k := 0; k < cols; k++ {
			fmt.Printf("[%d]", m.Board[i][k])
		}
		fmt.Printf("\n")
	}
}

func (m *model) createBoard() {
	rndTile1X, rndTile1Y := getStartTile()
	rndTile2X, rndTile2Y := getStartTile()
	m.Board[rndTile1X][rndTile1Y] = 2
	m.Board[rndTile2X][rndTile2Y] = 2
}

func getStartTile() (int, int) {
	return rand.Intn(cols), rand.Intn(cols)
}

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	m := model{}
	m.createBoard()
	m.printBoard()

}
