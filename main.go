package main

import (
	"github.com/Shopify/sarama"
	"github.com/vmutter/kafta/cmd"
)

func init() {

}

func main() {
	addrs := []string{""}
	sarama.NewClient(addrs, nil)

	cmd.Execute()
}
