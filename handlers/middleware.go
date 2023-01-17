package handlers

import (
	"context"
	"net/http"
	"quiz-app/quiz-api/data"
	"quiz-app/quiz-api/model"
)

// MiddlewareValidateAnswer validates the answer in the request and calls next if ok
func (quizHandler *QuizHandler) MiddlewareValidateAnswer(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		rw.Header().Add("Content-Type", "application/json")
		quizHandler.logger.Debug("Starting middleware")

		// Declare a new Answer struct.
		answer := model.Answer{}

		err := data.FromJSON(&answer, r.Body)
		if err != nil {
			quizHandler.logger.Error("Deserializing answer", "error", err)

			rw.WriteHeader(http.StatusBadRequest)
			data.ToJSON(&GenericError{Message: err.Error()}, rw)
			return
		}

		// validate the answer
		errs := quizHandler.validation.Validate(&answer)
		if len(errs) != 0 {
			quizHandler.logger.Error("Validating answer", "error", errs)

			// return the validation messages as an array
			rw.WriteHeader(http.StatusUnprocessableEntity)
			data.ToJSON(&ValidationError{Messages: errs.Errors()}, rw)
			return
		}

		// add the answer to the context
		ctx := context.WithValue(r.Context(), KeyAnswer{}, &answer)
		r = r.WithContext(ctx)

		// Call the next handler
		next.ServeHTTP(rw, r)
	})
}
