package entity

type Question struct {
	ID              uint
	Question        string
	PossibleAnswers []PossibleAnswer
	CorrectAnswer   string
	Difficulty      string
	CategoryID      uint
}

type PossibleAnswer struct {
	ID      uint
	Content string
	Choise  PossibleAnswerChoise
}

type Answer struct {
	ID         uint
	PlayerID   uint
	QuestionID uint
}


type PossibleAnswerChoise uint8

func (p PossibleAnswerChoise) IsValid () bool {
	if p >= PossibleAnswerA && p <= PossibleAnswerD {
		return true
	}
	return false
}

const (
	PossibleAnswerA PossibleAnswerChoise = iota + 1
	PossibleAnswerB
	PossibleAnswerC
	PossibleAnswerD
)
