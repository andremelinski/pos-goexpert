/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/spf13/cobra"
)

// testCmd represents the test command
var testCmd = &cobra.Command{
	Use:   "test",
	Short: "A brief description of test",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		command, _ := cmd.Flags().GetString("comando")
		if command == "ping"{
			cmd.Print("return ping \n")
		}else{
			cmd.Print(command + "\n")
		}
	},
}

func init() {
	rootCmd.AddCommand(testCmd)
	//  cria a flag
	testCmd.Flags().StringP("comando", "c", "", "escolha ping ou outra coisa")
	// faz a flag ser obrigatoria 
	testCmd.MarkFlagRequired("comando")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// testCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// testCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
