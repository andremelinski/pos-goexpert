/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/spf13/cobra"
)

var catgory string

// categoryCmd represents the category command
var categoryCmd = &cobra.Command{
	Use:   "category",
	Short: "A brief description of your command",
	Long: `A longer description `,
	Run: func(cmd *cobra.Command, args []string) {
		// cmd.Help()
		// name, _ := cmd.Flags().GetString("name")
		// shortName, _ := cmd.Flags().GetString("short-name")
		catName, _ := cmd.Flags().GetString("cat-name")
		// println(name)
		// println(shortName)
		println(catName)
	},
}

func init() {
	rootCmd.AddCommand(categoryCmd)
	// flag "global" para category e sub comandos. Da o valor a category 
	// pode ser utilizado em qualquer outro subcomando de category
	// categoryCmd.PersistentFlags().String("name", "", "Name of the category global scope")
	// categoryCmd.PersistentFlags().StringP("short-name", "n", "", "Name of the category global scope")
	// categoryCmd.Flags().String("local-name", "", "Name of the category local scope")

	// valor por referencia: valor eh mantido pelo ponteiro
	categoryCmd.PersistentFlags().StringVarP(&catgory, "cat-name", "c", "","Name of the category global scope" )
	// categoryCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
