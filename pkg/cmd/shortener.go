/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// shortenerCmd represents the shortener command
var shortenerCmd = &cobra.Command{
	Use:   "shortener",
	Short: "Make a long link short",
	Long:  "Use a Bitly API to make a long link short",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("shortener called")
	},
}

func init() {
	rootCmd.AddCommand(shortenerCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// shortenerCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// shortenerCmd.Flags().StringP()
}
