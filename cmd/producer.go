package cmd

import (
	"log"

	"github.com/Shopify/sarama"
	"github.com/spf13/cobra"
)

var Message string

func init() {
	rootCmd.AddCommand(producerCmd)

	producerCmd.PersistentFlags().StringVarP(&Message, "message", "m", "sample message", "message to send to broker")
}

var producerCmd = &cobra.Command{
	Use:     "producer",
	Aliases: []string{"p"},
	Short:   "Produce a message to a topic",
	Long:    "Produce a message to a topic",
	Run: func(cmd *cobra.Command, args []string) {
		producer, err := newProducer()
		if err != nil {
			log.Fatalln("Could not create producer: ", err)
		}

		msg := prepareMessage(KafkaTopic, Message)
		partition, offset, err := producer.SendMessage(msg)
		if err != nil {
			log.Fatalln("error to send message.", err.Error())
		} else {
			log.Printf("Message was saved to partion: %d.\nMessage offset is: %d.\n", partition, offset)
		}
	},
}

func newProducer() (sarama.SyncProducer, error) {
	config := sarama.NewConfig()
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Return.Successes = true
	producer, err := sarama.NewSyncProducer([]string{KafkaBroker}, config)

	return producer, err
}

func prepareMessage(topic, message string) *sarama.ProducerMessage {
	msg := &sarama.ProducerMessage{
		Topic:     topic,
		Partition: -1,
		Value:     sarama.StringEncoder(message),
	}

	return msg
}
