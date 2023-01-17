package model

import (
	"quiz-app/quiz-api/data"

	"github.com/hashicorp/go-hclog"
)

type Quiz struct {
	ID          int         `json:"ID"`
	Name        string      `json:"name"`
	Description string      `json:"description"`
	Questions   []*Question `json:"questions"`
}

type Question struct {
	Question string    `json:"question"`
	Choices  []*string `json:"choices"`
}

type Quizzes []*Quiz
type Questions []*Question

type QuizzesModels struct {
	loggger hclog.Logger
}

func NewQuizzesModels(logger hclog.Logger) *QuizzesModels {
	quizzesModel := QuizzesModels{loggger: logger}
	return &quizzesModel
}

func (quizzesModels QuizzesModels) GetAllQuizzes() ([]Quiz, error) {
	quizzesModel := []Quiz{}
	quizzes := data.GetAllQuizzes()

	for _, quiz := range quizzes {
		quizModel, err := quizzesModels.GetQuiz(quiz.ID)

		if err != nil {
			return nil, &ErrorGeneric{err}
		}
		quizzesModel = append(quizzesModel, *quizModel)
	}

	return quizzesModel, nil
}

func (quizzesModels QuizzesModels) GetQuiz(quizId int) (*Quiz, error) {
	quizModel := Quiz{}
	questionsModel := []*Question{}

	quiz, err := data.GetQuiz(quizId)
	if err != nil {
		return nil, &ErrorGeneric{err}
	}

	quizModel.ID = quiz.ID
	quizModel.Name = quiz.Name
	quizModel.Description = quiz.Description

	questions, err := data.GetQuizQuestions(quizId)
	if err != nil {
		return nil, &ErrorGeneric{err}
	}

	for _, question := range *questions {

		choices, err := GetChoicesStringArrays(question.ID)
		if err != nil {
			return nil, &ErrorGeneric{err}
		}

		questionModel := Question{
			Question: question.Question,
			Choices:  choices,
		}
		questionsModel = append(questionsModel, &questionModel)
	}

	quizModel.Questions = questionsModel

	return &quizModel, nil
}

func GetChoicesStringArrays(questionId int) ([]*string, error) {
	questionChoices := []*string{}

	choices, err := data.GetQuestionChoices(questionId)
	if err != nil {
		return nil, &ErrorGeneric{err}
	}

	for _, choice := range *choices {
		questionChoices = append(questionChoices, &choice.Choice)
	}

	return questionChoices, nil
}

func (quiz *Quiz) GetQuestions() []*Question {
	return quiz.Questions
}
