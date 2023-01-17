package handlers

import (
	"github.com/hashicorp/go-hclog"
	"quiz-app/quiz-api/data"
	"quiz-app/quiz-api/model"
)

// KeyAnswer is a key used for the Answer object in the context
type KeyAnswer struct{}

// Quiz handler for geting and anwering Quizzes
type QuizHandler struct {
	logger           hclog.Logger
	validation       *data.Validation
	quizzesModels    *model.QuizzesModels
	quizResultModels *model.QuizResultModels
	answerModels     *model.AnswerModels
}

// NewQuizHandler returns a new quiz handler with the given logger
func NewQuizHandler(logger hclog.Logger, validation *data.Validation, quizzesModels *model.QuizzesModels, quizResultModels *model.QuizResultModels, answerModel *model.AnswerModels) *QuizHandler {
	return &QuizHandler{logger, validation, quizzesModels, quizResultModels, answerModel}
}

// GenericError is a generic error message returned by a server
type GenericError struct {
	Message string `json:"message"`
}

// ValidationError is a collection of validation error messages
type ValidationError struct {
	Messages []string `json:"messages"`
}
