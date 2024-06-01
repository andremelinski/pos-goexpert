/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"database/sql"
	"log"
	"os"

	"github.com/andremelinski/15-cobra/internal/database"
	_ "github.com/mattn/go-sqlite3"
	"github.com/spf13/cobra"
)

func GetDb() *sql.DB{
	db, err := sql.Open("sqlite3", "./data.db")
	
	if err != nil {
		log.Fatal(err)
	}
	return db
}

func CategoryDB(dbConfig *sql.DB) database.Category {
	return *database.NewCategory(dbConfig)
}

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "15-cobra",
	Short: "A brief description of your application",
	Long: `A longer description `,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.15-cobra.yaml)")

	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}


