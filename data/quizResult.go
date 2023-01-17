package data

import "fmt"

type QuizResult struct {
	ID                  int
	quizID              int
	userID              int
	totalCorrectAnswers int
}

type QuizResults []*QuizResult

var quizResultsList = QuizResults{
	&QuizResult{ID: 1, quizID: 1, userID: 1, totalCorrectAnswers: 1},
	&QuizResult{ID: 2, quizID: 1, userID: 1, totalCorrectAnswers: 2},
}

type ErrorQuizResultNotFound struct {
	quizResultID int
}

func (err *ErrorQuizResultNotFound) Error() string {
	return fmt.Sprintf("Could not found answered quiz. Quiz result id : %v", err.quizResultID)
}

func GetAllQuizResults() *QuizResults {
	return &quizResultsList
}

func GetQuizResultsByQuizResultID(quizResultID int) (*QuizResult, error) {

	for _, quizResult := range quizResultsList {
		if quizResult.ID == quizResultID {
			return quizResult, nil
		}
	}

	return nil, &ErrorQuizResultNotFound{quizResultID}
}

func GetMaxQuizResultId() int {
	maxId := 0

	for _, quizResult := range quizResultsList {
		if quizResult.ID > maxId {
			maxId = quizResult.ID
		}
	}

	return maxId
}

func (quizResult *QuizResult) GetQuizID() int {
	return quizResult.quizID
}

func (quizResult *QuizResult) GetUserID() int {
	return quizResult.userID
}

func (quizResult *QuizResult) GetTotalCorrectAnswers() int {
	return quizResult.totalCorrectAnswers
}

func (quizResult *QuizResult) GetAnswers() (*Answers, error) {
	answers := Answers{}
	for _, answer := range answersList {
		if answer.quizResultID == quizResult.ID {
			answers = append(answers, answer)
		}
	}

	if len(answers) == 0 {
		return nil, &ErrorAnswersNotFound{quizResult.ID}
	}

	return &answers, nil
}

func CreateNewQuizResult(quizID int, userID int) QuizResult {

	id := GetMaxQuizResultId() + 1
	newQuizResult := QuizResult{ID: id,
		quizID:              quizID,
		userID:              userID,
		totalCorrectAnswers: 0}
	quizResultsList = append(quizResultsList, &newQuizResult)

	return newQuizResult
}

func (quizResult *QuizResult) UpdateTotalCorrectAnswer() {
	quizResult.totalCorrectAnswers++
}
