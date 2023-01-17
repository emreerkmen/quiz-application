package quizcli

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"quiz-app/quiz-api/data"

	//"quiz-app/quiz-api/data"
	"strconv"
)

type Quiz struct {
	ID          int       `json:"ID"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Questions   Questions `json:"questions"`
}

type Question struct {
	Question string   `json:"question"`
	Choices  []string `json:"choices"`
}

type Quizzes []Quiz
type Questions []Question
type Results []Result

type Result struct {
	QuizId              int                 `json:"quizID"`
	QuizName            string              `json:"quizName"`
	UserName            string              `json:"userName"`
	QuestionAndAnswers  []QuestionAndAnswer `json:"questionAndAnswers"`
	TotalCorrectAnswers int                 `json:"totalCorrectAnswers"`
	Status              int                 `json:"status"`
}

type QuestionAndAnswer struct {
	Question       string `json:"question"`
	SelectedAnswer string `json:"selectedAnswer"`
	CorrectAnswer  string `json:"correctAnswer"`
	Result         string `json:"result"`
}

type Answer struct {
	QuizID          int   `json:"quizID" validate:"required"`
	UserID          int   `json:"userID" validate:"required"`
	SelectedChoices []int `json:"selectedChoices" validate:"required,dive,min=-1"`
}

type QuestionAndAnswers []QuestionAndAnswer

type QuizResult struct {
	QuizResultID int
}

func GetQuizzes(beautify bool) (result string) {

	domain := "http://localhost:9090/v1/"
	uri := "quizzes"
	url := domain + uri

	resp, getErr := http.Get(url)
	if getErr != nil {
		log.Fatal(getErr)
	}

	quizzes := Quizzes{}

	return GetBeautifyJson(quizzes, resp.Body, beautify)
}

func GetQuizResults(beautify bool) (result string) {

	domain := "http://localhost:9090/v1/"
	uri := "quizResults"
	url := domain + uri

	resp, getErr := http.Get(url)
	if getErr != nil {
		log.Fatal(getErr)
	}

	if resp.StatusCode != 200 {
		body, err := io.ReadAll(resp.Body)
		// b, err := ioutil.ReadAll(resp.Body)  Go.1.15 and earlier
		if err != nil {
			log.Fatalln(err)
		}
		return string(body)
	}

	quizResults := Results{}

	return GetBeautifyJson(quizResults, resp.Body, beautify)
}

func GetQuizResult(quizResultID string, beautify bool) (result string) {

	domain := "http://localhost:9090/v1/"
	uri := "quizResult/" + quizResultID
	url := domain + uri

	resp, getErr := http.Get(url)
	if getErr != nil {
		log.Fatal(getErr)
	}

	if resp.StatusCode != 200 {
		body, err := io.ReadAll(resp.Body)
		// b, err := ioutil.ReadAll(resp.Body)  Go.1.15 and earlier
		if err != nil {
			log.Fatalln(err)
		}
		return string(body)
	}

	if resp.StatusCode != 200 {
		body, err := io.ReadAll(resp.Body)
		// b, err := ioutil.ReadAll(resp.Body)  Go.1.15 and earlier
		if err != nil {
			log.Fatalln(err)
		}
		return string(body)
	}

	quizResult := Result{}

	return GetBeautifyJson(quizResult, resp.Body, beautify)
}

func PostAnswer(quizID string, userID string, selectedAnswers string, beautify bool) (result string) {

	intQuizID, err := strconv.Atoi(quizID)
	if err != nil {
		log.Fatal(err)
	}

	intUserID, err := strconv.Atoi(userID)
	if err != nil {
		log.Fatal(err)
	}

	intSelectedAnswers := getIntArrayAnswers(selectedAnswers)

	domain := "http://localhost:9090/v1/"
	uri := "answer"
	url := domain + uri

	answer := Answer{QuizID: intQuizID, UserID: intUserID, SelectedChoices: intSelectedAnswers}
	jsonAnswer, err := json.Marshal(answer)
	if err != nil {
		log.Fatal(err)
	}

	requestBody := bytes.NewBuffer(jsonAnswer)
	resp, err := http.Post(url, "application/json", requestBody)
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		body, err := io.ReadAll(resp.Body)
		// b, err := ioutil.ReadAll(resp.Body)  Go.1.15 and earlier
		if err != nil {
			log.Fatalln(err)
		}
		return string(body)
	}

	quizResultID := QuizResult{}

	return GetBeautifyJson(quizResultID, resp.Body, beautify)
}

func getIntArrayAnswers(selectedAnswers string) []int {
	intSelectedAnswers := []int{}

	var ans string
	for _, char := range selectedAnswers + "," {

		if char == ',' {

			intAns, err := strconv.Atoi(ans)
			if err != nil {
				log.Fatal(err)
			}

			intSelectedAnswers = append(intSelectedAnswers, intAns)
			ans = ""
			continue
		}
		ans = ans + string(char)
	}
	return intSelectedAnswers
}

// Response body to string with beautify json
func GetBeautifyJson(i interface{}, r io.Reader, beautify bool) string {

	err := data.FromJSON(&i, r)
	if err != nil {
		log.Fatal("Deserializing quizzes", "error", err)
		return "Deserializing quizzes error"
	}

	if !beautify {
		return fmt.Sprintf("%v", i)
	}

	quizzesBea, err := json.MarshalIndent(i, "", " ")
	if err != nil {
		return ""
	}

	return string(quizzesBea)
}
