/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"challenge/pkg/proto"
	"context"
	"fmt"
	"io"
	"log"

	"github.com/spf13/cobra"
)

// timerCmd represents the timer command
var (
	Seconds   int64
	Frequency int64

	timerCmd = &cobra.Command{
		Use:   "timer",
		Short: "Starts timer",
		Long:  "Starts timer",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) < 1 {
				log.Fatal("not enought arguments")
			}

			startTimer(args[0], Seconds, Frequency)
		},
	}
)

func init() {
	rootCmd.AddCommand(timerCmd)

	timerCmd.Flags().Int64VarP(&Seconds, "seconds", "s", int64(5), "Timer duration in seconds (5 by default)")
	timerCmd.Flags().Int64VarP(&Frequency, "frequency", "f", int64(1), "Timer update frequency in seconds (1 by default)")
}

func startTimer(name string, seconds int64, frequency int64) {
	ptimer := &proto.Timer{
		Name:      name,
		Seconds:   seconds,
		Frequency: frequency,
	}
	fmt.Printf("%v\n", ptimer)

	timerClient, err := getClient().StartTimer(context.Background(), ptimer)
	if err != nil {
		log.Fatal(err)
	}

	for {
		timer, err := timerClient.Recv()
		if err == io.EOF {
			fmt.Println("Finished")
			return
		}
		if err != nil {
			log.Fatal(err)
		}

		if timer.Error != "" {
			fmt.Printf("\nTimer error: %s\n", timer.Error)
		} else {
			fmt.Printf("\nTimer name: %s\nSeconds remainnig: %v\n", timer.Name, timer.Seconds)
		}
	}
}
