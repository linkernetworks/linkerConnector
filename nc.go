package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"strings"

	linuxproc "github.com/c9s/goprocinfo/linux"
)

func netStringPut(host string, port int, out string) error {
	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", host, port))
	if err != nil {
		log.Println("err:", err)
		return err
	}
	defer conn.Close()
	_, err = io.Copy(conn, strings.NewReader(out))
	return err
}

func main() {
	// netStringPut("localhost", 9999, "asdffff")
	stat, err := linuxproc.ReadStat("/proc/stat")
	if err != nil {
		log.Fatal("stat read fail")
	}
	for _, s := range stat.CPUStats {
		// s.User
		// s.Nice
		// s.System
		// s.Idle
		// s.IOWait
		log.Println(s.User)
	}

	// stat.CPUStatAll
	// stat.CPUStats
	// stat.Processes
	// stat.BootTime
}
