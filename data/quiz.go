package data

import "fmt"

type quiz struct {
	ID          int
	Name        string
	Description string
}

type quizzes []*quiz

var quizzesList = quizzes{
	&quiz{ID: 1,
		Name:        "Animal Quiz",
		Description: "This is a animal quiz. You can find questions about animals."},
	&quiz{ID: 2,
		Name:        "Food Quiz",
		Description: "This quiz has questions that specific to food. Pls make the quiz with a full stomach :)."},
}

func (quiz quiz) String() string {
	return fmt.Sprintf("{%v %v}", quiz.ID, quiz.Name)
}

type ErrorQuizNotFound struct {
	id int
}

func (err *ErrorQuizNotFound) Error() string {
	return fmt.Sprintf("Quiz id: %v could not found.", err.id)
}

func GetAllQuizzes() quizzes {
	return quizzesList
}

func GetQuiz(id int) (*quiz, error) {

	for _, quiz := range quizzesList {
		if quiz.ID == id {
			return quiz, nil
		}
	}

	return nil, &ErrorQuizNotFound{id}
}

func GetQuizByID(quizID int) (*quiz, error) {

	for _, quiz := range quizzesList {
		if quiz.ID == quizID {
			return quiz, nil
		}
	}

	return nil, &ErrorQuizNotFound{quizID}
}

func GetQuizQuestions(quizId int) (*Questions, error) {

	quizQuestions := Questions{}

	for _, question := range questionsList {
		if question.quizID == quizId {
			quizQuestions = append(quizQuestions, question)
		}
	}

	if len(quizQuestions) == 0 {
		return nil, &ErrorQuestionNotFound{quizId}
	}

	return &quizQuestions, nil
}

func GetQuestionByQuizID(quizID int) (*Questions, error) {

	questions := Questions{}
	for _, question := range questionsList {
		if question.quizID == quizID {
			questions = append(questions, question)
		}
	}

	if len(questions) != 0 {
		return &questions, nil
	}

	return nil, &ErrorQuestionsNotFound{quizID}
}
