/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"challenge/pkg/proto"
	"context"
	"fmt"
	"log"

	"github.com/spf13/cobra"
)

// shortenerCmd represents the shortener command
var shortenerCmd = &cobra.Command{
	Use:   "shortener",
	Short: "Make a long link short",
	Long:  "Use a Bitly API to make a long link short",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			cmd.PrintErr("not enought arguments")
			return
		}
		startShortener(args[0])
	},
}

func init() {
	rootCmd.AddCommand(shortenerCmd)

	// shortenerCmd.Flags().StringP()
}

func startShortener(link string) {
	linkResponse, err := getClient().MakeShortLink(context.Background(), &proto.Link{Data: link})
	if err!= nil {
		log.Fatal(err)
	}
	
	fmt.Printf("Short link: %s\n", linkResponse.GetData())
}
