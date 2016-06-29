package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"strings"

	"github.com/Shopify/sarama"
)

func netStringPut(host string, out string) error {
	conn, err := net.Dial("tcp", host)
	if err != nil {
		log.Println("err:", err)
		return err
	}
	defer conn.Close()
	_, err = io.Copy(conn, strings.NewReader(out))
	return err
}

func kafkaProducer(srv string, topic string, key string, value string) {
	partitionerConstructor := sarama.NewHashPartitioner

	var keyEncoder, valueEncoder sarama.Encoder
	if key != "" {
		keyEncoder = sarama.StringEncoder(key)
	}
	if value != "" {
		valueEncoder = sarama.StringEncoder(value)
	}

	config := sarama.NewConfig()
	config.Producer.Partitioner = partitionerConstructor

	producer, err := sarama.NewSyncProducer(strings.Split(srv, ","), config)
	if err != nil {
		logger.Fatalln("FAILED to open the producer:", err)
	}
	defer producer.Close()

	partition, offset, err := producer.SendMessage(&sarama.ProducerMessage{
		Topic: topic,
		Key:   keyEncoder,
		Value: valueEncoder,
	})

	if err != nil {
		logger.Println("FAILED to produce message:", err)
	} else {
		fmt.Printf("topic=%s\tpartition=%d\toffset=%d\n", topic, partition, offset)
	}
}

func sendData(dest string, servAddr string, topic string, key string, value string) {
	switch {
	case dest == "kafka":
		kafkaProducer(servAddr, topic, key, value)
	case dest == "spark":
		netStringPut(servAddr, fmt.Sprintf("%s:%s", key, value))
	case dest == "stdout":
		log.Println("key=", key, " val=", value)
	default:
		log.Println("key=", key, " val=", value)
	}

}
