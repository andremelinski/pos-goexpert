/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/andremelinski/pos-goexpert/desafios/stress-test/internal/usecases"
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
	requests, _ := cmd.Flags().GetInt64("requests")
	concurrency, _ := cmd.Flags().GetInt64("concurrency")
	concent, err := usecases.NewStressURL(url, requests, concurrency).Stress()
	if err != nil{
		return err
	}
	cmd.Print("report completed\n")
	cmd.Println(*concent)
	return nil
}

func init() {
	createCmd := newStressCmd()
	rootCmd.AddCommand(createCmd)

	createCmd.Flags().StringP("url", "u", "", "service URL to test")
	createCmd.Flags().Int64P("requests","r", 0, "number of requests to perform")
	createCmd.Flags().Int64P("concurrency", "c",0, "number of simultaneous requests to make at a time")
	createCmd.MarkFlagsRequiredTogether("url", "requests", "concurrency")
}
