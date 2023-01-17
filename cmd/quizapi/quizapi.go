package main

import (
	"context"
	"fmt"
	gohandlers "github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/hashicorp/go-hclog"
	"log"
	"net/http"
	"os"
	"os/signal"
	"quiz-app/quiz-api/data"
	"quiz-app/quiz-api/handlers"
	"quiz-app/quiz-api/model"
	"time"
)

func main() {
	fmt.Println("Quiz app started.")

	logger := hclog.New(&hclog.LoggerOptions{
		Name:  "quiz-app",
		Level: hclog.LevelFromString("DEBUG"),
	})
	validation := data.NewValidation()

	// create quiz models instance
	quizModel := model.NewQuizzesModels(logger)
	quizResultModels := model.NewQuizResultModels(logger)
	answertModels := model.NewAnswerModels(logger)
	makeCuopleOfQuizzes(logger, *quizModel, *quizResultModels, *answertModels)

	// create the handlers
	quizHandler := handlers.NewQuizHandler(logger, validation, quizModel, quizResultModels, answertModels)

	// create a new router and register the handlers
	serverMux := mux.NewRouter()
	versionRouter := serverMux.PathPrefix("/v1").Subrouter()

	// Subrouter for get routers
	getRouter := versionRouter.Methods(http.MethodGet).Subrouter()
	getRouter.HandleFunc("/quizzes", quizHandler.GetAllQuizzes)
	getRouter.HandleFunc("/quizResults", quizHandler.GetAllQuizResults)
	getRouter.HandleFunc("/quizResult/{id:[0-9]+}", quizHandler.GetQuizResult)

	// Subrouter for get routers
	postRouter := versionRouter.Methods(http.MethodPost).Subrouter()
	postRouter.HandleFunc("/answer", quizHandler.AnswerQuiz)
	postRouter.Use(quizHandler.MiddlewareValidateAnswer)

	// CORS
	ch := gohandlers.CORS(gohandlers.AllowedOrigins([]string{"*"}))

	// create a new server
	server := http.Server{
		Addr:         ":9090",                                               // configure the bind address
		Handler:      ch(serverMux),                                         // set the default handler
		ErrorLog:     logger.StandardLogger(&hclog.StandardLoggerOptions{}), // set the logger for the server
		ReadTimeout:  5 * time.Second,                                       // max time to read request from the client
		WriteTimeout: 10 * time.Second,                                      // max time to write response to the client
		IdleTimeout:  120 * time.Second,                                     // max time for connections using TCP Keep-Alive
	}

	// start the server
	go func() {
		logger.Info("Starting server on port 9090")

		err := server.ListenAndServe(); 
		
		if err != http.ErrServerClosed {
			logger.Error("Error starting server", "error", err)
			os.Exit(1)
		}
	}()

	// trap sigterm or interupt and gracefully shutdown the server
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	// Block until a signal is received.
	sig := <-c
	logger.Debug("Got signal:", sig)

	// gracefully shutdown the server, waiting max 30 seconds for current operations to complete
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)

	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		logger.Error("Server forced to shutdown:", err)
	}

	log.Println("Server exiting")
}

func makeCuopleOfQuizzes(logger hclog.Logger, quizModels model.QuizzesModels, quizResultModels model.QuizResultModels, answerModel model.AnswerModels) {
	exampleAnswers := []*model.Answer{
		{QuizID: 1,
			UserID:          1,
			SelectedChoices: []int{3, 0, -1, 1, 0, 2}},
		{QuizID: 1,
			UserID:          1,
			SelectedChoices: []int{3, 1, -1, 1, 2, 2}},
		{QuizID: 1,
			UserID:          1,
			SelectedChoices: []int{3, 1, 2, 0, -1, 2}},
		{QuizID: 2,
			UserID:          1,
			SelectedChoices: []int{3, 1, -1, 1, 3, 2}},
		{QuizID: 2,
			UserID:          1,
			SelectedChoices: []int{3, 0, 2, 1, -1, 2}},
	}

	for _, a := range exampleAnswers {

		quizResultID, err := answerModel.AnswerQuiz(a)
		if err != nil {
			logger.Error("Quiz result error", "quizResultID", quizResultID, "err", err)
		}

		quizResultModels.GetResult(quizResultID)
	}
}
