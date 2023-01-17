package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"quiz-app/quiz-api/pkg/quizcli"
)

var postanswerCmd = &cobra.Command{
	Use:     "post-answer",
	Aliases: []string{"pa"},
	Short:   "Make a quiz",
	Args:    cobra.ExactArgs(3),
	Example: `post-answer {quizID} {userID} [selectedChoice,selectedChoice,selectedChoice,..] 
post-asnwer 1 1 1,1,1`,
	Run: func(cmd *cobra.Command, args []string) {

		quizID := args[0]
		userID := args[1]
		selectedAnswers := args[2]

		res := quizcli.PostAnswer(quizID, userID, selectedAnswers, beautify)
		fmt.Println(res)
	},
}

func init() {
	postanswerCmd.Flags().BoolVarP(&beautify, "beautify", "b", false, "Print result as beautiful json")
	rootCmd.AddCommand(postanswerCmd)
}
