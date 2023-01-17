package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"quiz-app/quiz-api/pkg/quizcli"
)

var getQuizResultsCmd = &cobra.Command{
	Use:     "get-quiz-results",
	Aliases: []string{"gqrs"},
	Short:   "Get all quiz results",
	Args:    cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		res := quizcli.GetQuizResults(beautify)
		fmt.Println(res)
	},
}

func init() {
	getQuizResultsCmd.Flags().BoolVarP(&beautify, "beautify", "b", false, "Print result as beautiful json")
	rootCmd.AddCommand(getQuizResultsCmd)
}
