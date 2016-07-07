package Sendr

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

//SendDataParam :
type SendDataParam struct {
	Dest    string
	SerAddr string
	Topic   string //use for keyspace/database if target is mysql/cassandra
	Key     string
	Value   string
	Table   string //only use if target is database
}

//Sender :
type Sender struct {
	App string
}

//NewSender :
func NewSender(app string) *Sender {
	s := new(Sender)
	s.App = app
	return s
}

//SendData :
func (s *Sender) SendData(senderParam SendDataParam) {
	switch senderParam.Dest {
	case "kafka":
		kafkaProducer(senderParam.SerAddr, senderParam.Topic, senderParam.Key, senderParam.Value)
	case "spark":
		netStringPut(senderParam.SerAddr, fmt.Sprintf("%s:%s", senderParam.Key, senderParam.Value))
	case "stdout":
		log.Println("key=", senderParam.Key, " val=", senderParam.Value)
	case "cassandra":
		var dbConfig DBConfig
		dbConfig.KeySpace = senderParam.Topic
		dbConfig.ServerList = append(dbConfig.ServerList, senderParam.SerAddr)
		c := NewCassandra(dbConfig)
		c.InsertKV(senderParam.Table, senderParam.Key, senderParam.Value)
	default:
		log.Println("key=", senderParam.Key, " val=", senderParam.Value)
	}

	//Write to file locally for all collect data.
	tmpDir, err := ioutil.TempDir("", fmt.Sprintf("%s-%s", s.App, senderParam.Key))
	if err != nil {
		panic(err)
	}

	temFile := fmt.Sprintf(path.Join(tmpDir, "%s-%d.json"), senderParam.Key, time.Now().UnixNano())
	log.Println("---> Write file to :", temFile)
	err = writeFile(temFile, senderParam.Value)

	if err != nil {
		log.Fatal("Write file error:", err)
	}
}

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
		log.Fatalln("FAILED to open the producer:", err)
	}
	defer producer.Close()

	partition, offset, err := producer.SendMessage(&sarama.ProducerMessage{
		Topic: topic,
		Key:   keyEncoder,
		Value: valueEncoder,
	})

	if err != nil {
		log.Println("FAILED to produce message:", err)
	} else {
		fmt.Printf("topic=%s\tpartition=%d\toffset=%d\n", topic, partition, offset)
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
