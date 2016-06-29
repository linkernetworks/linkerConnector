package main

import (
	"encoding/json"
	"log"
	"time"

	linuxproc "github.com/c9s/goprocinfo/linux"
)

//DataCollector : Data Collector currently working for Linux first.
type DataCollector struct {
}

//NewDataCollector : Object constructor for data collector
func NewDataCollector() *DataCollector {
	d := new(DataCollector)
	return d
}

//GetProcessInfo :Get ProcessInfo JSON format string.
func (d *DataCollector) GetProcessInfo() string {
	var retProcessInfo ProcessInfo
	status, err := linuxproc.ReadProcessStatus("/proc/self/status")
	if err != nil {
		log.Fatal("status read fail")
	}
	retProcessInfo.ProcInfo = *status
	retProcessInfo.Timestamp = time.Now().Unix()
	//TODO. Still need implement since FileInfo is not exist in proc

	//json marshaling
	retJSON, err := json.Marshal(retProcessInfo)
	if err != nil {
		log.Println("marshall json failed:", err)
		return ""
	}

	return string(retJSON)
}

//GetMachineInfo :Get MachineInfo JSON format string.
func (d *DataCollector) GetMachineInfo() string {
	var retMachineInfo MachineInfo
	retMachineInfo.Timestamp = time.Now().Unix()
	//TODO. Still need implement since all information cannot get from /proc

	retJSON, err := json.Marshal(retMachineInfo)
	if err != nil {
		log.Println("marshall json failed:", err)
		return ""
	}
	return string(retJSON)
}
