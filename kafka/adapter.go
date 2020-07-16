package main

import (
	"fmt"
	"os"
	"os/signal"
	"pdk/src/adapter/elasticContainer"
	"pdk/src/adapter/kafkaContainer"
	"syscall"

	"gopkg.in/confluentinc/confluent-kafka-go.v1/kafka"
)

func main() {

	if len(os.Args) < 4 {
		fmt.Fprintf(os.Stderr, "Usage: %s <broker> <group> <topics..>\n",
			os.Args[0])
		os.Exit(1)
	}

	broker := os.Args[1]
	group := os.Args[2]
	topics := os.Args[3:]

	sigchan := make(chan os.Signal, 1)
	signal.Notify(sigchan, syscall.SIGINT, syscall.SIGTERM)
	output := make(chan string)
	exitKafka := make(chan int)
	exitElastic := make(chan int)

	kConf := kafkaContainer.MakeConfig(broker, group)
	c, err := kafka.NewConsumer(kConf)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to create consumer: %s\n", err)
		os.Exit(1)
	}
	fmt.Printf("Created Consumer %v\n", c)
	err = c.SubscribeTopics(topics, nil)
	kafkaContainer.RunKafkaConsumer(c, output, exitKafka)

	es, err := elasticContainer.NewElasticCli()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to create es: %s\n", err)
		os.Exit(1)
	}
	elasticContainer.RunElasticCli(es, output, exitElastic)

	<-sigchan
	exitElastic <- 0
	exitKafka <- 0
}
