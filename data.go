package main

import (
	linuxproc "github.com/c9s/goprocinfo/linux"
)

//ProcessDetail :
type ProcessDetail struct {
	ProcID     uint64                  `json:"proc_id"`
	StatusInfo linuxproc.ProcessStatus `json:"status_info"`
	StateInfo  linuxproc.ProcessStat   `json:"stat_info"`
	inputRate  int                     `json:"input_rate"`
	outputRate int                     `json:"output_rate"`
}

//ProcessInfo :
type ProcessInfo struct {
	MachineID string `json:"machine_id"`
	//Timestamp : Unix time
	Timestamp int64 `json:"timestamp"`

	Procs []ProcessDetail `json:"procs"`
}

//MachineInfo :Machine information
type MachineInfo struct {
	MachineID string `json:"machine_id"`
	//Timestamp : Unix time
	Timestamp int64             `json:"timestamp"`
	CPUInfo   linuxproc.CPUInfo `json:"cpu_info"`
	MemInfo   linuxproc.MemInfo `json:"mem_info"`

	NetInfo []struct {
		Protocal   string `json:"protocal"`
		Mac        string `json:"mac"`
		IP         string `json:"ip"`
		Rate       string `json:"rate"`
		Errs       string `json:"errs"`
		Drop       string `json:"drop"`
		Compressed string `json:"compressed"`
	} `json:"net_info"`
	DiskInfo []struct {
		inputRate    int     `json:"input_rate"`
		outputRate   int     `json:"output_rate"`
		errRate      int     `json:"err_rate"`
		seriesNumber string  `json:"serues_number"`
		typeOfDisk   string  `json:"disk_type"`
		diskSize     int     `json:"disk_size"`
		usuage       float32 `json:"usuage"`
	} `json:"disk_info"`
}
