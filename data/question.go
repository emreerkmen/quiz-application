package data

import "fmt"

type Question struct {
	ID              int
	Question        string
	correctChoiceID int
	quizID          int
}

type Questions []*Question

var questionsList = Questions{
	&Question{ID: 1,
		Question:        "In which country can you find the smallest jellyfish?",
		correctChoiceID: 2,
		quizID:          1},
	&Question{ID: 2,
		Question:        "What animals are Holstein Friesians?",
		correctChoiceID: 3,
		quizID:          1},
	&Question{ID: 3,
		Question:        "Which famous horse race is held in Melbourne, Australia?",
		correctChoiceID: 3,
		quizID:          1},
	&Question{ID: 4,
		Question:        "Which famous dolphin can you find in Clearwater Aquarium, Florida?",
		correctChoiceID: 1,
		quizID:          1},
	&Question{ID: 5,
		Question:        "Where can you find the largest coral reef in the world?",
		correctChoiceID: 0,
		quizID:          1},
	&Question{ID: 6,
		Question:        "Which country has the most sheep in the world?",
		correctChoiceID: 0,
		quizID:          1},
	&Question{ID: 7,
		Question:        "Which one of the following puddings is not sweet?",
		correctChoiceID: 0,
		quizID:          2},
	&Question{ID: 8,
		Question:        "Which one of the following ingredients wouldn’t you find in a full English breakfast?",
		correctChoiceID: 2,
		quizID:          2},
	&Question{ID: 9,
		Question:        "Which pasta dish is made with layers of sauce and flat pasta sheets?",
		correctChoiceID: 0,
		quizID:          2},
	&Question{ID: 10,
		Question:        "If you are ordering a Chicken tikka masala, what type of takeaway are you getting food from?",
		correctChoiceID: 2,
		quizID:          2},
	&Question{ID: 11,
		Question:        "Which one of the following apple types is green?",
		correctChoiceID: 3,
		quizID:          2},
	&Question{ID: 12,
		Question:        "What is the main ingredient of the Spanish dish, Paella?",
		correctChoiceID: 2,
		quizID:          2},
}

type ErrorQuestionNotFound struct {
	quizId int
}

func (err *ErrorQuestionNotFound) Error() string {
	return fmt.Sprintf("Could not found any question. Quiz id : %v", err.quizId)
}

type ErrorQuestionsNotFound struct {
	quizId int
}

func (err *ErrorQuestionsNotFound) Error() string {
	return fmt.Sprintf("Could not found questions for quiz. Quiz id : %v", err.quizId)
}

func (question Question) String() string {
	return fmt.Sprintf("{%v %v %v %v}", question.ID, question.Question, question.correctChoiceID, question.quizID)
}

func GetAllQuestions() Questions {
	return questionsList
}

func GetQuestionChoices(questionId int) (*choices, error) {

	questionChoices := choices{}

	for _, choice := range choicesList {
		if choice.questionID == questionId {
			questionChoices = append(questionChoices, choice)
		}
	}

	if len(questionChoices) == 0 {
		return nil, &ErrorChoiceNotFound{questionId}
	}

	return &questionChoices, nil
}

func (question *Question) GetCorrectAnswer() int {
	return question.correctChoiceID
}
