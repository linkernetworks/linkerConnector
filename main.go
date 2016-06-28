package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/Shopify/sarama"
	linuxproc "github.com/c9s/goprocinfo/linux"
)

var (
	serverAddr = flag.String("server", "localhost:9092", "The comma separated list of server could be brokers in the Kafka cluster or spark address")
	topic      = flag.String("topic", "", "The topic to produce")
	interval   = flag.Int("interval", 0, "interval to retrieval data(millisecond), default 0 is not repeat.")
	dest       = flag.String("dest", "kafka", "Destination to kafka, spark and stdout")
	logger     = log.New(os.Stderr, "", log.LstdFlags)
)

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

func sendData(topic string, key string, value string) {
	switch {
	case *dest == "kafka":
		kafkaProducer(*serverAddr, topic, key, value)
	case *dest == "spark":
		netStringPut(*serverAddr, fmt.Sprintf("%s:%s", key, value))
	case *dest == "stdout":
		log.Println("key=", key, " val=", value)
	default:
		log.Println("key=", key, " val=", value)
	}

}

func main() {
	flag.Parse()

	//default with verbose
	sarama.Logger = logger

	for {

		stat, err := linuxproc.ReadStat("/proc/stat")
		if err != nil {
			log.Fatal("stat read fail")
		}

		for _, s := range stat.CPUStats {
			log.Println("Get data:", s)
			sendData(*topic, "User", strconv.FormatUint(s.User, 64))
			sendData(*topic, "Nice", strconv.FormatUint(s.Nice, 64))
		}

		sendData(*topic, "Processes", strconv.FormatUint(stat.Processes, 64))

		time.Sleep(time.Millisecond * time.Duration(*interval))

	}

}
