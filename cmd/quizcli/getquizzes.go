package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"quiz-app/quiz-api/pkg/quizcli"
)

var beautify bool
var getQuizzesCmd = &cobra.Command{
	Use:     "get-quizzes",
	Aliases: []string{"gq"},
	Short:   "Get all quizzes",
	Args:    cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		res := quizcli.GetQuizzes(beautify)
		fmt.Println(res)
	},
}

func init() {
	getQuizzesCmd.Flags().BoolVarP(&beautify, "beautify", "b", false, "Print result as beautiful json")
	rootCmd.AddCommand(getQuizzesCmd)
}
