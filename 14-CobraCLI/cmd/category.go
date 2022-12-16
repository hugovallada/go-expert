/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// categoryCmd represents the category command
var categoryCmd = &cobra.Command{
	Use:   "category",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		exists, _ := cmd.Flags().GetBool("exists")
		names, _ := cmd.Flags().GetStringSlice("names")
		fmt.Println(names)
		fmt.Println(exists)
		fmt.Println("category called: ", category)
	},
	PreRun: func(cmd *cobra.Command, args []string) {
		fmt.Println("Antes do run")
	},
	PostRun: func(cmd *cobra.Command, args []string) {
		fmt.Println("Depois do run")
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		return fmt.Errorf("ocorreu um erro")
	},
}

var category string

func init() {
	rootCmd.AddCommand(categoryCmd)
	categoryCmd.PersistentFlags().StringVarP(&category, "name", "n", "", "Name of the category") // serve pra todos os filhos do category
	categoryCmd.Flags().BoolP("exists", "e", false, "Check if category exists")
	categoryCmd.Flags().StringSlice("names", []string{}, "")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// categoryCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// categoryCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
