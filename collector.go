package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os/exec"
	"regexp"
	"strconv"
	"time"

	linuxproc "github.com/c9s/goprocinfo/linux"
)

var (
	processID = regexp.MustCompile(`^[0-9]*$`)
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
	files, err := ioutil.ReadDir("/proc")
	if err != nil {
		log.Fatal(err)
	}

	var retProcessInfo ProcessInfo

	for _, file := range files {
		if processID.MatchString(file.Name()) {

			var procDetail ProcessDetail
			procDetail.ProcID, _ = strconv.ParseUint(file.Name(), 10, 64)

			status, err := linuxproc.ReadProcessStatus(fmt.Sprintf("/proc/%s/status", file.Name()))
			if err != nil {
				log.Println("status read fail.")
			} else {
				procDetail.StatusInfo = *status
			}

			stat, err := linuxproc.ReadProcessStat(fmt.Sprintf("/proc/%s/stat", file.Name()))
			if err != nil {
				log.Println("status read fail.")
			} else {
				procDetail.StateInfo = *stat
			}

			retProcessInfo.Procs = append(retProcessInfo.Procs, procDetail)
		}
		log.Println(file.Name())
	}

	stdout, err := exec.Command("hostname").CombinedOutput()
	if err != nil {
		log.Println("hostname cannot retrieval.")
	} else {
		retProcessInfo.MachineID = string(stdout)
		log.Println("Machine ID:", retProcessInfo.MachineID)
	}

	retProcessInfo.Timestamp = time.Now().Unix()

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
