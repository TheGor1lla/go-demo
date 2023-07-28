/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/jboursiquot/go-proverbs"
	"github.com/spf13/cobra"
)

// printCmd represents the print command
var printCmd = &cobra.Command{
	Use:   "print",
	Short: "Prints proverbs",
	Long: `Nifty little tool that prints Go Proverbs in your terminal`,
	Run: func(cmd *cobra.Command, args []string) {
		count, _ :=cmd.Flags().GetInt("count")
		for i := 0; i < count; i++ {
			fmt.Println(proverbs.Random().Saying)			
		}
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		count, err :=cmd.Flags().GetInt("count")
		if err != nil {
			return err
		}
		for i := 0; i < count; i++ {
			fmt.Println(proverbs.Random().Saying)			
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(printCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// printCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// printCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	printCmd.Flags().IntP("count", "c", 1, "Count of proverbs to print")
}
