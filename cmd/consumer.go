package cmd

import (
	"fmt"

	"github.com/Shopify/sarama"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(consumerCmd)
}

var consumerCmd = &cobra.Command{
	Use:     "consumer",
	Aliases: []string{"c"},
	Short:   "Listen to a topic from a Kafka broker",
	Long:    "Listen to a topic from a Kafka broker",
	Run: func(cmd *cobra.Command, args []string) {
		consumer, err := sarama.NewConsumer([]string{KafkaBroker}, nil)
		if err != nil {
			fmt.Println("Could not create consumer: ", err)
		}

		subscribe(KafkaTopic, consumer)

		waitExit()
	},
}

func subscribe(topic string, consumer sarama.Consumer) {
	partitionList, err := consumer.Partitions(topic)
	if err != nil {
		fmt.Println("Error retrieving partitionList ", err)
	}

	initialOffset := sarama.OffsetNewest

	for _, partition := range partitionList {
		pc, _ := consumer.ConsumePartition(topic, partition, initialOffset)

		go func(pc sarama.PartitionConsumer) {
			for message := range pc.Messages() {
				fmt.Println(string(message.Value))
			}
		}(pc)
	}
}
