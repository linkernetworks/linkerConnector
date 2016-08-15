package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"runtime"
	"time"

	sendr "github.com/LinkerNetworks/linkerConnector/sender"
	"github.com/spf13/cobra"
)

var (
	send *sendr.Sender
)

func main() {
	send = sendr.NewSender("linkerConnector")
	var serverAddr, topic, dest string
	var interval int
	var usingPipe, disableFile bool
	rootCmd := &cobra.Command{
		Use:   "linkerConnector",
		Short: "Linux system data collector and send to target destination",
		Run: func(cmd *cobra.Command, args []string) {
			if usingPipe {
				info, _ := os.Stdin.Stat()

				if info.Size() > 0 {
					reader := bufio.NewReader(os.Stdin)
					processPipe(reader, dest, serverAddr, topic, disableFile)
				}
				return
			}
			if runtime.GOOS != "linux" {
				fmt.Println("Collect data from linux for now, application exit.")
				return
			}

			for {
				data := NewDataCollector()
				procInfo := data.GetProcessInfo()
				machineInfo := data.GetMachineInfo()

				send.SendData(sendr.SendDataParam{Dest: dest, SerAddr: serverAddr, Topic: topic, Key: "ProcessInfo", Value: procInfo, DisableFileSave: disableFile})
				send.SendData(sendr.SendDataParam{Dest: dest, SerAddr: serverAddr, Topic: topic, Key: "MachineInfo", Value: machineInfo, DisableFileSave: disableFile})

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
	rootCmd.Flags().BoolVarP(&usingPipe, "pipe", "p", false, "Using pipe mode to forward data")
	rootCmd.Flags().BoolVarP(&disableFile, "dsiableFileSave", "f", false, "Disable local file save.")

	rootCmd.Execute()

}

func processPipe(reader *bufio.Reader, dest, serverAddr, topic string, disableFileSave bool) {
	line := 1
	for {
		input, err := reader.ReadString('\n')
		if err != nil && err == io.EOF {
			break
		}
		send.SendData(sendr.SendDataParam{Dest: dest, SerAddr: serverAddr, Topic: topic, Key: "Pipe", Value: input, DisableFileSave: disableFileSave})
		line++
	}
}
