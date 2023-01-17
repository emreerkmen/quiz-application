package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

type QuizCli struct{}

var rootCmd = &cobra.Command{
	Use:   "quizcli",
	Short: "quizcli - a simple CLI to make a litle quizzes with RestFull Api",
	Long: `quizcli a simple CLI to make a litle quizzes with RestFull Api
   
For now you can just make a quiz that already found in app.`,
	Run: func(cmd *cobra.Command, args []string) {

	},
}

func (qc QuizCli) Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Whoops. There was an error while executing your CLI '%s'", err)
		os.Exit(1)
	}
}
