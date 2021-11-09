package utils

type Result struct {
	Message   string `json: message`
	Finnished bool   `json: finnish`
}

const (
	PLAYER   = 1
	COMPUTER = 2
)

var MESSAGES = []string{"Keep playing", "You won", "Game over"}

func PlayRound(buttonsSelected string) Result {
	finnished, winner := determineWinner(buttonsSelected)
	message := MESSAGES[winner]

	var result Result
	result.Message = message
	result.Finnished = finnished
	return result
}

func determineWinner(buttonsSelected string) (bool, int) {
	return false, 0
}
