# quiz-app-go
This is a simple quiz app that developed with golang.

It uses;
1. Gorilla Mux for Rest Api
2. Cobra for Cli

# App Start

1. Quiz Api

Local Environment
```bash
//in project root directory
$cd cmd/quizapi
$go run .
```

Docker
```bash
//in project root directory
$docker build -t quiz-api:multistage -f Dockerfile.multistage .
$docker run -p 9090:9090 quiz-api:multistage
```

2. Quiz Cli
```bash
//in project root directory
$cd cmd/quizcli
$go run . --help
```

# End Points and Samples

## GET "/v1/quizzes"

Lists available quizzes

response:
```bash
[
    {
        "ID": 1,
        "name": "Animal Quiz",
        "description": "This is a animal quiz. You can find questions about animals.",
        "questions": [
            {
                "question": "In which country can you find the smallest jellyfish?",
                "choices": [
                    "South Africa",
                    "China",
                    "Australia",
                    "Iceland"
                ]
            },
            {
                "question": "What animals are Holstein Friesians?",
                "choices": [
                    "Cats",
                    "Dogs",
                    "Horses",
                    "Cows"
                ]
            }
        ]
    },
    {
        "ID": 2,
        "name": "Food Quiz",
        "description": "This quiz has questions that specific to food. Pls make the quiz with a full stomach :).",
        "questions": [
            {
                "question": "Which one of the following puddings is not sweet?",
                "choices": [
                    "Sticky Toffee pudding",
                    "Christmas pudding",
                    "Yorkshire pudding",
                    "Bread and butter pudding"
                ]
            },
            {
                "question": "Which one of the following ingredients wouldn’t you find in a full English breakfast?",
                "choices": [
                    "Baked Beans",
                    "Mushroom",
                    "Broccoli",
                    "Bacon"
                ]
            }
        ]
    }
]
```

## GET "/v1/quizResults"

Lists all quiz results

response:
```bash
[
    {
        "quizID": 1,
        "quizName": "Animal Quiz",
        "userName": "Emre",
        "questionAndAnswers": [
            {
                "question": "In which country can you find the smallest jellyfish?",
                "selectedAnswer": "",
                "correctAnswer": "Australia",
                "result": "Empty."
            },
            {
                "question": "What animals are Holstein Friesians?",
                "selectedAnswer": "Dogs",
                "correctAnswer": "Cows",
                "result": "Wrong Answer :("
            },
            {
                "question": "Which famous horse race is held in Melbourne, Australia?",
                "selectedAnswer": "Melbourne Cup",
                "correctAnswer": "Melbourne Cup",
                "result": "Correct Answer :)"
            },
            {
                "question": "Which famous dolphin can you find in Clearwater Aquarium, Florida?",
                "selectedAnswer": "Flipper",
                "correctAnswer": "Winter",
                "result": "Wrong Answer :("
            },
            {
                "question": "Where can you find the largest coral reef in the world?",
                "selectedAnswer": "Mexico",
                "correctAnswer": "Australia",
                "result": "Wrong Answer :("
            },
            {
                "question": "Which country has the most sheep in the world?",
                "selectedAnswer": "Wales",
                "correctAnswer": "China",
                "result": "Wrong Answer :("
            }
        ],
        "totalCorrectAnswers": 1,
        "status": "You scored higher than 25% of all quizzers"
    },
    {
        "quizID": 1,
        "quizName": "Animal Quiz",
        "userName": "Emre",
        "questionAndAnswers": [
            {
                "question": "In which country can you find the smallest jellyfish?",
                "selectedAnswer": "South Africa",
                "correctAnswer": "South Africa",
                "result": "Correct Answer :)"
            },
            {
                "question": "What animals are Holstein Friesians?",
                "selectedAnswer": "Horses",
                "correctAnswer": "Horses",
                "result": "Correct Answer :)"
            },
            {
                "question": "Which famous horse race is held in Melbourne, Australia?",
                "selectedAnswer": "Melbourne Cup",
                "correctAnswer": "The Grand National",
                "result": "Wrong Answer :("
            },
            {
                "question": "Which famous dolphin can you find in Clearwater Aquarium, Florida?",
                "selectedAnswer": "Flipper",
                "correctAnswer": "Fungie",
                "result": "Wrong Answer :("
            },
            {
                "question": "Where can you find the largest coral reef in the world?",
                "selectedAnswer": "Mexico",
                "correctAnswer": "Belize",
                "result": "Wrong Answer :("
            },
            {
                "question": "Which country has the most sheep in the world?",
                "selectedAnswer": "Wales",
                "correctAnswer": "India",
                "result": "Wrong Answer :("
            }
        ],
        "totalCorrectAnswers": 2,
        "status": "You scored higher than 57% of all quizzers"
    }
]    
```

## GET "/v1/quizResult/{id}"

lists the result with the given quiz result id

response:
```bash
{
    "quizID": 1,
    "quizName": "Animal Quiz",
    "userName": "Emre",
    "questionAndAnswers": [
        {
            "question": "In which country can you find the smallest jellyfish?",
            "selectedAnswer": "South Africa",
            "correctAnswer": "Australia",
            "result": "Wrong Answer :("
        },
        {
            "question": "What animals are Holstein Friesians?",
            "selectedAnswer": "Cats",
            "correctAnswer": "Cows",
            "result": "Wrong Answer :("
        },
        {
            "question": "Which famous horse race is held in Melbourne, Australia?",
            "selectedAnswer": "The Grand National",
            "correctAnswer": "Melbourne Cup",
            "result": "Wrong Answer :("
        },
        {
            "question": "Which famous dolphin can you find in Clearwater Aquarium, Florida?",
            "selectedAnswer": "Flipper",
            "correctAnswer": "Winter",
            "result": "Wrong Answer :("
        },
        {
            "question": "Where can you find the largest coral reef in the world?",
            "selectedAnswer": "Australia",
            "correctAnswer": "Australia",
            "result": "Correct Answer :)"
        },
        {
            "question": "Which country has the most sheep in the world?",
            "selectedAnswer": "China",
            "correctAnswer": "China",
            "result": "Correct Answer :)"
        }
    ],
    "totalCorrectAnswers": 2,
    "status": "You scored higher than 57% of all quizzers"
}
```

## Post "/v1/answer"

Answer a quiz.
1. Each selectedChoice corresponds to the questions in the quiz in order
2. "-1" means Empty choice

request;
```bash
{
    "quizID": 1,
    "userId": 1,
    "selectedChoices": [
        -1,
        1,
        0,
        2,
        0,
        3
    ]
}
```

response:
```bash
{
    "QuizResultID": 8
}
```
