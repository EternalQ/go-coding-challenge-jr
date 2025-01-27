package cmd

import (
	"challenge/pkg/proto"
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
	"google.golang.org/grpc"
)

// rootCmd represents the root command
var (
	rootCmd = &cobra.Command{
		Use:   "root",
		Short: "Challenge gRPC client",
		Long:  "Challenge gRPC client",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("root called")
		},
	}
)

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// rootCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func getClient() proto.ChallengeServiceClient {
	conn, err := grpc.Dial(":8080", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}

	return proto.NewChallengeServiceClient(conn)
}
