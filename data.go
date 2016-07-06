package main

import (
	linuxproc "github.com/c9s/goprocinfo/linux"
)

//ProcessDetail :
type ProcessDetail struct {
	ProcID     uint64                  `json:"proc_id"`
	StatusInfo linuxproc.ProcessStatus `json:"status_info"`
	StateInfo  linuxproc.ProcessStat   `json:"stat_info"`
}

type ProcessIO struct {
	ProcID     uint64  `json:"proc_id"`
	DiskInput  int     `json:"disk_input"`
	DiskOutput int     `json:"disk_output"`
	NetInput   int     `json:"net_input"`
	NetOutput  int     `json:"net_output"`
}

//ProcessInfo :
type ProcessInfo struct {
	MachineID string `json:"machine_id"`
	//Timestamp : Unix time
	Timestamp int64 `json:"timestamp"`

	Procs []ProcessDetail `json:"procs"`

	ProcIO []ProcessIO `json:"proc_io"`
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
		InputRate  int    `json:"input_rate"`
		OutputRate int    `json:"output_rate"`
	} `json:"net_info"`
	DiskInfo []struct {
		InputRate    int     `json:"input_rate"`
		OutputRate   int     `json:"output_rate"`
		ErrRate      int     `json:"err_rate"`
		SeriesNumber string  `json:"serues_number"`
		TypeOfDisk   string  `json:"disk_type"`
		DiskSize     int     `json:"disk_size"`
		Usuage       float32 `json:"usuage"`
	} `json:"disk_info"`
}
