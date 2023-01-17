package handlers

import (
	"net/http"
	"quiz-app/quiz-api/data"
	"strconv"

	"github.com/gorilla/mux"
)

// GetAllQuizResults handles GET requests and returns all current quiz results
func (quizHandler *QuizHandler) GetAllQuizResults(rw http.ResponseWriter, r *http.Request) {
	quizHandler.logger.Debug("Get all records")
	rw.Header().Add("Content-Type", "application/json")

	quizResults, err := quizHandler.quizResultModels.GetAllResults()

	if err != nil {
		http.Error(rw, err.Error(), http.StatusBadRequest)
		return
	}

	err = data.ToJSON(quizResults, rw)
	if err != nil {
		// we should never be here but log the error just incase
		quizHandler.logger.Error("Unable to serializing quiz results", "error", err)
	}
}

// GetQuizResult handles GET requests and return quiz result with id
func (quizHandler *QuizHandler) GetQuizResult(rw http.ResponseWriter, r *http.Request) {
	quizHandler.logger.Debug("Get the recors")
	rw.Header().Add("Content-Type", "application/json")

	quizResultId := quizHandler.getQuizResultID(r)

	quizResult, err := quizHandler.quizResultModels.GetResult(quizResultId)

	if err != nil {
		http.Error(rw, err.Error(), http.StatusBadRequest)
		return
	}

	err = data.ToJSON(quizResult, rw)
	if err != nil {
		// we should never be here but log the error just incase
		quizHandler.logger.Error("Unable to serializing quiz resuls", "error", err)
	}
}

func (quizHandler *QuizHandler) getQuizResultID(r *http.Request) int {
	// parse the quiz result id from the url
	vars := mux.Vars(r)

	// convert the id into an integer and return
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		// should never happen
		quizHandler.logger.Error("Unable to serializing quizzes", "error", err)
	}

	return id
}
