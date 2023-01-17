package handlers

import (
	"net/http"
	"quiz-app/quiz-api/data"
)

// GetAllQuizzes handles GET requests and returns all current quizzes
func (quizHandler *QuizHandler) GetAllQuizzes(rw http.ResponseWriter, r *http.Request) {
	quizHandler.logger.Debug("Get all records")
	rw.Header().Add("Content-Type", "application/json")

	quizzes, err := quizHandler.quizzesModels.GetAllQuizzes()

	if err != nil {
		http.Error(rw, err.Error(), http.StatusBadRequest)
		return
	}

	err = data.ToJSON(quizzes, rw)
	if err != nil {
		// we should never be here but log the error just incase
		quizHandler.logger.Error("Unable to serializing quizzes", "error", err)
	}
}
