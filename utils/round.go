package utils

import (
	"log"
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

func PlayRound(buttonsSelected string) Result {
	finnished, winner := checkWinner(buttonsSelected)
	message := MESSAGES[winner]

	var result Result
	result.Message = message
	result.Finnished = finnished
	return result
}

func checkWinner(buttonsSelected string) (bool, int) {

	for i := 0; i < 3; i++ {
		// check rows
		log.Println(buttonsSelected[3*i], buttonsSelected[3*i+1], buttonsSelected[3*i+2])
		if buttonsSelected[3*i] != '0' && (buttonsSelected[3*i] == buttonsSelected[3*i+1]) && (buttonsSelected[3*i+1] == buttonsSelected[3*i+2]) {
			return determineWinner(buttonsSelected[3*i])
		}
		// check columns
		if buttonsSelected[i] != '0' && (buttonsSelected[i] == buttonsSelected[i+3]) && (buttonsSelected[i+3] == buttonsSelected[i+6]) {
			return determineWinner(buttonsSelected[i])
		}
		log.Println(i)
	}
	// check diagonals
	if buttonsSelected[0] != '0' && (buttonsSelected[0] == buttonsSelected[4]) && (buttonsSelected[4] == buttonsSelected[8]) {
		return determineWinner(buttonsSelected[0])
	}
	if buttonsSelected[2] != '0' && (buttonsSelected[2] == buttonsSelected[4]) && (buttonsSelected[4] == buttonsSelected[6]) {
		return determineWinner(buttonsSelected[2])
	}
	return false, 0
}

func determineWinner(buttonSelected byte) (bool, int) {
	winner, _ := strconv.Atoi(string(buttonSelected))
	return true, winner
}
