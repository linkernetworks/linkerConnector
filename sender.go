package main

import (
	"io"
	"log"
	"net"
	"strings"
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
