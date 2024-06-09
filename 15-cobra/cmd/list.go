/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("list called")
		catName, _ := cmd.Flags().GetString("cat-name")
		println(catName + " - catName called by ref")
	},
	// roda antes do comando ser executado
	PreRun: func(cmd *cobra.Command, args []string) {
		fmt.Println("PreRun called")
	},
	// roda depois do comando ser executado
	PostRun: func(cmd *cobra.Command, args []string) {
		fmt.Println("PostRun called")
	},
	//  outra forma de executar o run mas esse retorna um erro 
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("RunE called")
		return fmt.Errorf("erro")
	},
}

func init() {
	categoryCmd.AddCommand(listCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}