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

	"github.com/docker/engine-api/client"
	"github.com/docker/engine-api/types"
	"golang.org/x/net/context"

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

			pStat, err := linuxproc.ReadProcessStat(fmt.Sprintf("/proc/%s/stat", file.Name()))
			if err != nil {
				log.Println("status read fail.")
			} else {
				procDetail.StateInfo = *pStat
			}

			retProcessInfo.Procs = append(retProcessInfo.Procs, procDetail)
		}
	}

	retProcessInfo.DockerStat = getDockerContainerStat()
	retProcessInfo.MachineID = getMachineID()
	retProcessInfo.Timestamp = getUnixTimestamp()

	//json marshal
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
	retMachineInfo.MachineID = getMachineID()
	retMachineInfo.Timestamp = getUnixTimestamp()

	mInfo, err := linuxproc.ReadMemInfo("/proc/meminfo")
	if err != nil {
		log.Println("memory info read fail.")
	} else {
		retMachineInfo.MemInfo = *mInfo
	}

	cInfo, err := linuxproc.ReadCPUInfo("/proc/cpuinfo")
	if err != nil {
		log.Println("CPU info read fail.")
	} else {
		retMachineInfo.CPUInfo = *cInfo
	}

	stat, err := linuxproc.ReadStat("/proc/stat")
	if err != nil {
		log.Println("stat read fail.")
	} else {
		retMachineInfo.Stat = *stat
	}
	
	err = d.GetDMIInfo(&retMachineInfo)
	if err != nil {
		log.Println("Get DMI error:", err)
	}

	retJSON, err := json.Marshal(retMachineInfo)
	if err != nil {
		log.Println("marshall json failed:", err)
		return ""
	}
	return string(retJSON)
}

func getUnixTimestamp() int64 {
	return time.Now().Unix()
}
func getMachineID() string {
	stdout, err := exec.Command("hostname").CombinedOutput()
	if err != nil {
		log.Println("hostname cannot retrieval.")
		return ""
	}

	return string(stdout)
}

//GetDMIInfo :
func (d *DataCollector) GetDMIInfo(mInfo *MachineInfo) error {
	dmi := NewDMI()
	err := dmi.Run()
	if err != nil {
		log.Println("Get DMI error:", err, ", Please install dmidecode before use this.")
		return err
	}

	si, err := dmi.SearchByName("System Information")
	if err != nil {
		log.Println("Parse SI failed.")
	}

	rSI := SystemInformation{
		Manufacturer: si["Manufacturer"],
		ProductName:  si["Product Name"],
		Version:      si["Version"],
		SerialNumber: si["Serial Number"],
		UUID:         si["UUID"],
		WakeupType:   si["Wakeup Type"],
		SKUNumber:    si["SKU Number"],
		Family:       si["Family"]}
	mInfo.SysInfo = rSI

	bi, err := dmi.SearchByName("BIOS Information")
	if err != nil {
		log.Println("Parse BI failed.")
	}

	rBI := BIOSInfo{
		Vendor:          bi["Vendor"],
		Version:         bi["Version"],
		ReleaseDate:     bi["Release Date"],
		Address:         bi["Address"],
		RuntimeSize:     bi["Runtime Size"],
		ROMSize:         bi["ROM Size"],
		Characteristics: bi["Characteristics"]}
	mInfo.BiosInfo = rBI
	return nil
}

//GetDockerContainerStat :
func getDockerContainerStat() []ContainerInfo {
	var retContainerInfo []ContainerInfo

	defaultHeaders := map[string]string{"User-Agent": "engine-api-cli-1.0"}
	cli, err := client.NewClient("unix:///var/run/docker.sock", "v1.22", nil, defaultHeaders)
	if err != nil {
		log.Println("No Docker CLient: err=", err)
		return nil
	}

	options := types.ContainerListOptions{All: true}
	containers, err := cli.ContainerList(context.Background(), options)
	if err != nil {
		log.Println("No ContainerList : err=", err)
		return nil
	}

	for _, c := range containers {
		fmt.Println(c.ID)
		body, err := cli.ContainerStats(context.Background(), c.ID, false)
		defer body.Close()
		if err != nil {
			log.Println("No ContainerStats : err=", err)
			return nil
		}
		var jsonC ContainerInfo
		bytContent, err := ioutil.ReadAll(body)

		if err := json.Unmarshal(bytContent, &jsonC); err != nil {
			log.Println("Unmarshal error : err=", err)
			return nil
		}
		retContainerInfo = append(retContainerInfo, jsonC)
		// log.Println("Container:", c.ID, string(bytContent))
	}

	return nil
}
