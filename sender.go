package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net"
	"os"
	"path"
	"strings"
	"time"

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

	//Write to file locally for all collect data.
	tmpDir, err := ioutil.TempDir("", fmt.Sprintf("linkerConnector-%s", key))
	if err != nil {
		panic(err)
	}

	temFile := fmt.Sprintf(path.Join(tmpDir, "%s-%d.json"), key, time.Now().UnixNano())
	log.Println("---> Write file to :", temFile)
	err = writeFile(temFile, value)

	if err != nil {
		log.Fatal("Write file error:", err)
	}
}

func writeFile(file string, data string) error {
	f, err := os.OpenFile(file, os.O_RDWR|os.O_CREATE, 0600)
	if err != nil {
		return err
	}

	_, err = fmt.Fprintf(f, "%s", data)

	if err != nil {
		f.Close()
		return err
	}

	f.Sync()
	f.Close()
	return nil
}
