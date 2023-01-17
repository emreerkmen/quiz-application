package data

import "fmt"

type Answer struct {
	ID               int
	quizResultID     int
	questionID       int
	correctChoiceID  int
	selectedChoiceID int
	result           int //0=empty, 1=true, 2=false
}

type Answers []*Answer

var answersList = Answers{
	&Answer{ID: 1, quizResultID: 1,
		questionID: 1, correctChoiceID: 2, selectedChoiceID: -1},
	&Answer{ID: 2, quizResultID: 1,
		questionID: 2, correctChoiceID: 3, selectedChoiceID: 1},
	&Answer{ID: 3, quizResultID: 1,
		questionID: 3, correctChoiceID: 3, selectedChoiceID: 3},
	&Answer{ID: 4, quizResultID: 1,
		questionID: 4, correctChoiceID: 1, selectedChoiceID: 0},
	&Answer{ID: 5, quizResultID: 1,
		questionID: 5, correctChoiceID: 0, selectedChoiceID: 2},
	&Answer{ID: 6, quizResultID: 1,
		questionID: 6, correctChoiceID: 0, selectedChoiceID: 1},
	&Answer{ID: 7, quizResultID: 2,
		questionID: 7, correctChoiceID: 0, selectedChoiceID: 0},
	&Answer{ID: 8, quizResultID: 2,
		questionID: 8, correctChoiceID: 2, selectedChoiceID: 2},
	&Answer{ID: 9, quizResultID: 2,
		questionID: 9, correctChoiceID: 0, selectedChoiceID: 3},
	&Answer{ID: 10, quizResultID: 2,
		questionID: 10, correctChoiceID: 2, selectedChoiceID: 0},
	&Answer{ID: 11, quizResultID: 2,
		questionID: 11, correctChoiceID: 3, selectedChoiceID: 2},
	&Answer{ID: 12, quizResultID: 2,
		questionID: 12, correctChoiceID: 2, selectedChoiceID: 1},
}

func (answer Answer) String() string {
	return fmt.Sprintf("{%v %v %v %v %v %v}", answer.ID, answer.quizResultID, answer.questionID, answer.correctChoiceID, answer.selectedChoiceID, answer.result)
}

type ErrorAnswerNotFound struct {
	answerId int
}

func (err *ErrorAnswerNotFound) Error() string {
	return fmt.Sprintf("Could not found answer. Answer id : %v", err.answerId)
}

type ErrorAnswersNotFound struct {
	quizResultID int
}

func (err *ErrorAnswersNotFound) Error() string {
	return fmt.Sprintf("Could not found answers for quiz result. Quiz result id : %v", err.quizResultID)
}

func GetAllAnswers() Answers {
	return answersList
}

func GetAnswer(answerId int) (*Answer, error) {

	for _, answer := range answersList {
		if answer.ID == answerId {
			return answer, nil
		}
	}

	return nil, &ErrorAnswerNotFound{answerId: answerId}
}

func GetMaxAnswersId() int {
	maxId := 0

	for _, answer := range answersList {
		if answer.ID > maxId {
			maxId = answer.ID
		}
	}

	return maxId
}

func CreateNewAnswer(quizResultID int, question *Question, selectedChoiceID int) int {
	id := GetMaxAnswersId() + 1
	quizResult, err := GetQuizResultsByQuizResultID(quizResultID)

	if err != nil {
		fmt.Println(err)
	}

	result := 0
	if selectedChoiceID == question.correctChoiceID {
		result = 1
		quizResult.UpdateTotalCorrectAnswer()
	} else if selectedChoiceID != question.correctChoiceID && selectedChoiceID != -1 {
		result = 2
	}

	newAnswer := Answer{ID: id,
		quizResultID:     quizResultID,
		questionID:       question.ID,
		correctChoiceID:  question.correctChoiceID,
		selectedChoiceID: selectedChoiceID,
		result:           result}
	answersList = append(answersList, &newAnswer)

	return id
}

func (answer *Answer) GetSelectedChoiceID() int {
	return answer.selectedChoiceID
}

func (answer *Answer) GetCorrectChoiceID() int {
	return answer.correctChoiceID
}

func (answer *Answer) GetResult() int {
	return answer.result
}
