/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/andremelinski/15-cobra/internal/database"
	"github.com/spf13/cobra"
)
type RunEFunc  func(cmd *cobra.Command, args []string) error

// decopling and inject DB connection
func newCreateCmd(categoryDB database.Category) *cobra.Command{
	return &cobra.Command{
		Use:   "create",
		Short: "A brief description of your command",
		Long: `A longer description `,
		RunE: runCreate(categoryDB),
	}
}

// createCmd represents the create command
// var createCmd = &cobra.Command{
// 	Use:   "create",
// 	Short: "A brief description of your command",
// 	Long: `A longer description `,
	// Run: func(cmd *cobra.Command, args []string) {
	// 	db := GetDb()
	// 	category := CategoryDB(db)

	// 	catName, _ := cmd.Flags().GetString("cat-name")
	// 	description, _ := cmd.Flags().GetString("description")
	// 	println(catName)
	// 	cat, err := category.Create(catName, description)
	// 	if err != nil {
	// 		cmd.PrintErrln(err)
	// 		return
	// 	}
	// 	cmd.Println(cat)
	// },

	// --> decopling using RunEFunc type and inject DB dependencies in categoryDB
// 	RunE: runCreate(CategoryDB(GetDb())),
// }

func runCreate(categoryDB database.Category) RunEFunc {
	return func(cmd *cobra.Command, args []string) error {
		catName, _ := cmd.Flags().GetString("cat-name")
		description, _ := cmd.Flags().GetString("description")
		cat, err := categoryDB.Create(catName, description)
		cmd.Println(cat)
		if err !=nil{
			return err
		}
		return nil
	}
}

func init() {
	createCmd := newCreateCmd(CategoryDB(GetDb()))
	categoryCmd.AddCommand(createCmd)
	createCmd.Flags().StringP("description", "d", "", "category description")
	// createCmd.MarkFlagsRequiredTogether("cat-name", "description")
	createCmd.MarkFlagRequired("description")
	categoryCmd.MarkFlagRequired("cat-name")

}
