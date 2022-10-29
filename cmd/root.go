package cmd

import (
	"fmt"
	"os"
	"os/signal"

	"github.com/spf13/cobra"
)

var KafkaBroker string
var KafkaTopic string

var rootCmd = &cobra.Command{
	Use:   "kafta",
	Short: "Kafta is a simple Apache Kafta client",
	Long:  "Kafta is a simple Apache Kafta client",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(cmd.Short)
	},
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&KafkaBroker, "broker", "b", "localhost:9092", "kafka broker address")
	rootCmd.PersistentFlags().StringVarP(&KafkaTopic, "topic", "t", "sample", "kafka topic")
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func waitExit() {
	done := make(chan os.Signal, 1)

	signal.Notify(done, os.Interrupt)

	for range done {
		os.Exit(0)
	}
}
