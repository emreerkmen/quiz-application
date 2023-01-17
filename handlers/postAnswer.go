package handlers

import (
	"net/http"
	"quiz-app/quiz-api/data"
	"quiz-app/quiz-api/model"
)

type QuizResult struct {
	QuizResultID int
}

// Answer a quiz
func (quizHandler *QuizHandler) AnswerQuiz(rw http.ResponseWriter, r *http.Request) {
	quizHandler.logger.Debug("Starting answering quiz")

	// Get answer from context
	answer := r.Context().Value(KeyAnswer{}).(*model.Answer)
	quizHandler.logger.Debug("Post", "Answer", answer)

	ID, err := quizHandler.answerModels.AnswerQuiz(answer)
	quizResultID := QuizResult{QuizResultID: ID}

	if err != nil {
		quizHandler.logger.Error("There are erros when answering quiz", "Answer", answer, "error", err)
		http.Error(rw, err.Error(), http.StatusBadRequest)
		return
	}

	err = data.ToJSON(quizResultID, rw)
	if err != nil {
		quizHandler.logger.Error("Unable to serializing quiz result", "error", err)
		return
	}

}
