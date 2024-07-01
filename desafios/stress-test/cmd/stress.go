/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/andremelinski/pos-goexpert/desafios/Stress-test/internal/usecases"
	"github.com/spf13/cobra"
)

// stressCmd represents the stress command
func newStressCmd() *cobra.Command{
	return &cobra.Command{
		Use:   "stress",
		Short: "A brief description of your command",
		Long: `A longer description `,
		RunE: runStress,
	}
}

func runStress(cmd *cobra.Command, args []string) error {
	url, _ := cmd.Flags().GetString("url")
	requests, _ := cmd.Flags().GetUint64("requests")
	concurrency, _ := cmd.Flags().GetUint64("concurrency")

	input := usecases.StressTestInput{
		URL:         url,
		Requests:    requests,
		Concurrency: concurrency,
	}

	fmt.Println(input)
	usecases.Aqui()


	return nil
}

func init() {
	createCmd := newStressCmd()
	rootCmd.AddCommand(createCmd)

	createCmd.Flags().StringP("url", "u", "", "service URL to test")
	createCmd.Flags().Uint64P("requests","r", 0, "number of requests to perform")
	createCmd.Flags().Uint64P("concurrency", "c",0, "number of simultaneous requests to make at a time")
	createCmd.MarkFlagsRequiredTogether("url", "requests", "concurrency")
	
}
