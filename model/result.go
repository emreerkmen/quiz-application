package model

import (
	"fmt"
	"github.com/hashicorp/go-hclog"
	"quiz-app/quiz-api/data"
)

type Result struct {
	QuizId              int                  `json:"quizID"`
	QuizName            string               `json:"quizName"`
	UserName            string               `json:"userName"`
	QuestionAndAnswers  []*QuestionAndAnswer `json:"questionAndAnswers"`
	TotalCorrectAnswers int                  `json:"totalCorrectAnswers"`
	Status              string               `json:"status"`
}

type QuestionAndAnswer struct {
	Question       string `json:"question"`
	SelectedAnswer string `json:"selectedAnswer"`
	CorrectAnswer  string `json:"correctAnswer"`
	Result         string `json:"result"`
}

type QuestionAndAnswers []*QuestionAndAnswer

type QuizResultModels struct {
	logger hclog.Logger
}

func NewQuizResultModels(logger hclog.Logger) *QuizResultModels {
	quizResultModels := QuizResultModels{logger: logger}
	return &quizResultModels
}

func (quizResultModesl QuizResultModels) GetAllResults() ([]Result, error) {
	resultModels := []Result{}
	results := data.GetAllQuizResults()

	for _, result := range *results {
		quizResult, err := quizResultModesl.GetResult(result.ID)
		if err != nil {
			return nil, &ErrorGeneric{err}
		}

		resultModels = append(resultModels, *quizResult)
	}

	return resultModels, nil
}

func (quizResultModesl QuizResultModels) GetResult(quizResultID int) (*Result, error) {
	result := Result{}

	quizResult, err := data.GetQuizResultsByQuizResultID(quizResultID)
	if err != nil {
		quizResultModesl.logger.Error("Error", "err", err)
		return nil, &ErrorQuizResult{message: err}
	}

	result.QuizId = quizResult.GetQuizID()

	quiz, err := data.GetQuizByID(result.QuizId)
	if err != nil {
		quizResultModesl.logger.Error("Error", "err", err)
		return nil, &ErrorQuizResult{message: err}
	}

	result.QuizName = quiz.Name

	user, err := data.GetUser(quizResult.GetUserID())
	if err != nil {
		quizResultModesl.logger.Error("Error", "err", err)
		return nil, &ErrorQuizResult{message: err}
	}

	result.UserName = user.GetUserName()

	questions, err := data.GetQuizQuestions(result.QuizId)
	if err != nil {
		quizResultModesl.logger.Error("Error", "err", err)
		return nil, &ErrorQuizResult{message: err}
	}

	answers, err := quizResult.GetAnswers()
	if err != nil {
		quizResultModesl.logger.Error("Error", "err", err)
		return nil, &ErrorQuizResult{message: err}
	}

	result.QuestionAndAnswers = GetQuestionsAndAnswers(answers, questions)
	result.TotalCorrectAnswers = quizResult.GetTotalCorrectAnswers()
	result.Status = CalculateStatus(quizResult.GetTotalCorrectAnswers(), quizResultID)

	return &result, nil
}

func (queAndAns QuestionAndAnswer) String() string {
	return fmt.Sprintf("{%v %v %v %v }", queAndAns.Question, queAndAns.SelectedAnswer, queAndAns.CorrectAnswer, queAndAns.Result)
}

func CalculateStatus(currentTotalCorrectAnswer int, quizResultID int) string {
	quizResults := data.GetAllQuizResults()
	worstQuizResultsAmount := 0.0

	for _, quizResult := range *quizResults {
		if quizResult.GetTotalCorrectAnswers() < currentTotalCorrectAnswer && quizResultID != quizResult.ID {
			worstQuizResultsAmount++
		}
	}

	if worstQuizResultsAmount == 0 {
		return "You scored the most low in all the quizzers"
	}

	value := int(worstQuizResultsAmount / float64((len(*quizResults) - 1)) * 100)

	return fmt.Sprintf("You scored higher than %v%% of all quizzers", value)
}

type ErrorQuizResult struct {
	message error
}

func (err *ErrorQuizResult) Error() string {
	return fmt.Sprintf("Error when show quizzes result: %v", err.message)
}

type ErrorGeneric struct {
	message error
}

func (err *ErrorGeneric) Error() string {
	return fmt.Sprintf("%v", err.message)
}

func GetQuestionsAndAnswers(answers *data.Answers, questions *data.Questions) QuestionAndAnswers {

	questionAndAnswers := QuestionAndAnswers{}
	for index, question := range *questions {
		questionText := question.Question

		choices, err := data.GetQuestionChoices(question.ID)
		if err != nil {
			fmt.Println(err)
		}

		correctChoiceID := (*answers)[index].GetCorrectChoiceID()
		selectedChoiceID := (*answers)[index].GetSelectedChoiceID()

		var selectedAnswer string
		if selectedChoiceID != -1 {
			selectedAnswer = (*choices)[selectedChoiceID].Choice
		}
		correctAnswer := (*choices)[correctChoiceID].Choice

		questionAndAnswer := QuestionAndAnswer{Question: questionText,
			SelectedAnswer: selectedAnswer,
			CorrectAnswer:  correctAnswer}

		if selectedChoiceID == -1 {
			questionAndAnswer.Result = "Empty."
		} else if selectedChoiceID == correctChoiceID {
			questionAndAnswer.Result = "Correct Answer :)"
		} else {
			questionAndAnswer.Result = "Wrong Answer :("
		}

		questionAndAnswers = append(questionAndAnswers, &questionAndAnswer)
	}

	return questionAndAnswers
}
