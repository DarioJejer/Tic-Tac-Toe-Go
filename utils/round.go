package utils

type Result struct {
	Message   string
	Finnished bool
}

func PlayRound(buttonsSelected string) Result {
	var result Result
	result.Message = "Keep playing!"
	result.Finnished = false
	return result
}
