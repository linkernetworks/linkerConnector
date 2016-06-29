package main

import (
	"fmt"
	"log"
	"os"
	"runtime"
	"time"

	"github.com/Shopify/sarama"
	"github.com/spf13/cobra"
)

var (
	logger = log.New(os.Stderr, "", log.LstdFlags)
)

func main() {
	sarama.Logger = logger

	var serverAddr, topic, dest string
	var interval int
	rootCmd := &cobra.Command{
		Use:   "linkerConnector",
		Short: "Linux system data collector and send to target destination",
		Run: func(cmd *cobra.Command, args []string) {
			if runtime.GOOS != "linux" {
				fmt.Println("Collect data from linux for now, application exist.")
				return
			}

			for {
				data := NewDataCollector()
				procInfo := data.GetProcessInfo()
				machineInfo := data.GetMachineInfo()

				sendData(dest, serverAddr, topic, "ProcessInfo", procInfo)
				sendData(dest, serverAddr, topic, "MachineInfo", machineInfo)

				if interval == 0 {
					return
				}

				time.Sleep(time.Millisecond * time.Duration(interval))
			}
		},
	}

	rootCmd.Flags().IntVarP(&interval, "interval", "i", 0, "Interval to retrieval data(millisecond), default 0 is not repeat.")

	rootCmd.Flags().StringVarP(&serverAddr, "server", "s", "", "The comma separated list of server could be brokers in the Kafka cluster or spark address")
	rootCmd.Flags().StringVarP(&topic, "topic", "t", "", "The topic to kafka produce")
	rootCmd.Flags().StringVarP(&dest, "dest", "d", "stdout", "Destination to kafka, spark and stdout")

	rootCmd.Execute()
}
