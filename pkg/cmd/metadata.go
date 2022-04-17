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
	"google.golang.org/grpc/metadata"
)

// metadataCmd represents the metadata command
var metadataCmd = &cobra.Command{
	Use:   "metadata",
	Short: "Read metadata",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
		startMetadata()
	},
}

func init() {
	rootCmd.AddCommand(metadataCmd)
}

func startMetadata() {
	ctx := metadata.NewOutgoingContext(context.Background(), metadata.MD{})
	ctx = metadata.AppendToOutgoingContext(ctx, "somekey", "some important data")

	phResp, err := getClient().ReadMetadata(ctx, &proto.Placeholder{})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Short link: %s\n", phResp.GetData())
}
