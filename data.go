package main

import (
	linuxproc "github.com/c9s/goprocinfo/linux"
	"github.com/google/cadvisor/info/v1"
)

//ProcessDetail :
type ProcessDetail struct {
	ProcID     uint64                  `json:"proc_id"`
	StatusInfo linuxproc.ProcessStatus `json:"status_info"`
	StateInfo  linuxproc.ProcessStat   `json:"stat_info"`
}

//ProcessIO :
type ProcessIO struct {
	ProcID     uint64 `json:"proc_id"`
	DiskInput  int    `json:"disk_input"`
	DiskOutput int    `json:"disk_output"`
	NetInput   int    `json:"net_input"`
	NetOutput  int    `json:"net_output"`
}

//ProcessInfo :
type ProcessInfo struct {
	MachineID string `json:"machine_id"`
	Timestamp int64  `json:"timestamp"`

	Procs  []ProcessDetail `json:"procs"`
	ProcIO []ProcessIO     `json:"proc_io"`
	Stat   linuxproc.Stat  `json:"stat"`
	// DockerStat []ContainerInfo `json:"container_info"`
	DockerStat []v1.ContainerInfo `json:"docker_stat"`
}

//MachineInfo :Machine information
type MachineInfo struct {
	MachineID string            `json:"machine_id"`
	Timestamp int64             `json:"timestamp"`
	CPUInfo   linuxproc.CPUInfo `json:"cpu_info"`
	MemInfo   linuxproc.MemInfo `json:"mem_info"`
	BiosInfo  BIOSInfo          `json:"bios_info"`
	SysInfo   SystemInformation `json:"system_info"`
	Stat   	  linuxproc.Stat    `json:"stat"`

}

// BIOS Information
// 	Vendor: innotek GmbH
// 	Version: VirtualBox
// 	Release Date: 12/01/2006
// 	Address: 0xE0000
// 	Runtime Size: 128 kB
// 	ROM Size: 128 kB
// 	Characteristics:
// 		ISA is supported
// 		PCI is supported
// 		Boot from CD is supported
// 		Selectable boot is supported
// 		8042 keyboard services are supported (int 9h)
// 		CGA/mono video services are supported (int 10h)
// 		ACPI is supported

//BIOSInfo :
type BIOSInfo struct {
	Vendor          string `json:"vendor"`
	Version         string `json:"version"`
	ReleaseDate     string `json:"release_date"`
	Address         string `json:"address"`
	RuntimeSize     string `json:"runtime_size"`
	ROMSize         string `json:"rom_size"`
	Characteristics string `json:"characteristics"`
}

// System Information
// 	Manufacturer: innotek GmbH
// 	Product Name: VirtualBox
// 	Version: 1.2
// 	Serial Number: 0
// 	UUID: F548DD5F-057D-4F7F-9465-FC529E045C08
// 	Wake-up Type: Power Switch
// 	SKU Number: Not Specified
// 	Family: Virtual Machine

//SystemInformation :
type SystemInformation struct {
	Manufacturer string `json:"manufacturer"`
	ProductName  string `json:"product_name"`
	Version      string `json:"version"`
	SerialNumber string `json:"serial_number"`
	UUID         string `json:"uuid"`
	WakeupType   string `json:"wakeup_type"`
	SKUNumber    string `json:"sku_number"`
	Family       string `json:"Family"`
}

