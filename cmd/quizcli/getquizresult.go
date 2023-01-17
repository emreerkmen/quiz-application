package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"quiz-app/quiz-api/pkg/quizcli"
)

var getQuizResultCmd = &cobra.Command{
	Use:     "get-quiz-result",
	Aliases: []string{"gqr"},
	Short:   "Get quiz result with given id",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		res := quizcli.GetQuizResult(args[0], beautify)
		fmt.Println(res)
	},
}

func init() {
	getQuizResultCmd.Flags().BoolVarP(&beautify, "beautify", "b", false, "Print result as beautiful json")
	rootCmd.AddCommand(getQuizResultCmd)
}
