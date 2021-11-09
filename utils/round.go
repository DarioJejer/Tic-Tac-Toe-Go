package utils

import (
	"strconv"
)

type Result struct {
	Message   string `json:"message"`
	Finnished bool   `json:"finnished"`
}

const (
	PLAYER   = 1
	COMPUTER = 2
)

var MESSAGES = []string{"Keep playing", "You won", "Game over"}

// the board string will be formated as follow

// board = "001020100"

// 0 0 1
// 0 2 0
// 1 0 0

// where 0 is an empty position
// 1 is a player move
// 0 is a computer move

func PlayRound(board string) Result {
	finnished, winner := checkWinner(board)
	message := MESSAGES[winner]

	var result Result
	result.Message = message
	result.Finnished = finnished
	return result
}

func checkWinner(board string) (bool, int) {

	for i := 0; i < 3; i++ {
		// check rows
		if board[3*i] != '0' && (board[3*i] == board[3*i+1]) && (board[3*i+1] == board[3*i+2]) {
			return determineWinner(board[3*i])
		}
		// check columns
		if board[i] != '0' && (board[i] == board[i+3]) && (board[i+3] == board[i+6]) {
			return determineWinner(board[i])
		}
	}
	// check diagonals
	if board[0] != '0' && (board[0] == board[4]) && (board[4] == board[8]) {
		return determineWinner(board[0])
	}
	if board[2] != '0' && (board[2] == board[4]) && (board[4] == board[6]) {
		return determineWinner(board[2])
	}
	return false, 0
}

func determineWinner(button byte) (bool, int) {
	winner, _ := strconv.Atoi(string(button))
	return true, winner
}
