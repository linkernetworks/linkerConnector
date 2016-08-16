package main

import (
	"time"

	linuxproc "github.com/c9s/goprocinfo/linux"
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

	Procs      []ProcessDetail `json:"procs"`
	ProcIO     []ProcessIO     `json:"proc_io"`
	Stat       linuxproc.Stat  `json:"stat"`
	DockerStat []ContainerInfo `json:"container_info"`
}

//MachineInfo :Machine information
type MachineInfo struct {
	MachineID string            `json:"machine_id"`
	Timestamp int64             `json:"timestamp"`
	CPUInfo   linuxproc.CPUInfo `json:"cpu_info"`
	MemInfo   linuxproc.MemInfo `json:"mem_info"`
	BiosInfo  BIOSInfo          `json:"bios_info"`
	SysInfo   SystemInformation `json:"system_info"`

	// NetInfo []struct {
	// 	Protocal   string `json:"protocal"`
	// 	Mac        string `json:"mac"`
	// 	IP         string `json:"ip"`
	// 	Rate       string `json:"rate"`
	// 	Errs       string `json:"errs"`
	// 	Drop       string `json:"drop"`
	// 	Compressed string `json:"compressed"`
	// 	InputRate  int    `json:"input_rate"`
	// 	OutputRate int    `json:"output_rate"`
	// } `json:"net_info"`
	// DiskInfo []struct {
	// 	InputRate    int     `json:"input_rate"`
	// 	OutputRate   int     `json:"output_rate"`
	// 	ErrRate      int     `json:"err_rate"`
	// 	SeriesNumber string  `json:"serues_number"`
	// 	TypeOfDisk   string  `json:"disk_type"`
	// 	DiskSize     int     `json:"disk_size"`
	// 	Usuage       float32 `json:"usuage"`
	// } `json:"disk_info"`
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

//ContainerInfo :
type ContainerInfo struct {
	Read        time.Time `json:"read"`
	PrecpuStats struct {
		CPUUsage struct {
			TotalUsage        int   `json:"total_usage"`
			PercpuUsage       []int `json:"percpu_usage"`
			UsageInKernelmode int   `json:"usage_in_kernelmode"`
			UsageInUsermode   int   `json:"usage_in_usermode"`
		} `json:"cpu_usage"`
		SystemCPUUsage int64 `json:"system_cpu_usage"`
		ThrottlingData struct {
			Periods          int `json:"periods"`
			ThrottledPeriods int `json:"throttled_periods"`
			ThrottledTime    int `json:"throttled_time"`
		} `json:"throttling_data"`
	} `json:"precpu_stats"`
	CPUStats struct {
		CPUUsage struct {
			TotalUsage        int   `json:"total_usage"`
			PercpuUsage       []int `json:"percpu_usage"`
			UsageInKernelmode int   `json:"usage_in_kernelmode"`
			UsageInUsermode   int   `json:"usage_in_usermode"`
		} `json:"cpu_usage"`
		SystemCPUUsage int64 `json:"system_cpu_usage"`
		ThrottlingData struct {
			Periods          int `json:"periods"`
			ThrottledPeriods int `json:"throttled_periods"`
			ThrottledTime    int `json:"throttled_time"`
		} `json:"throttling_data"`
	} `json:"cpu_stats"`
	MemoryStats struct {
		Usage    int `json:"usage"`
		MaxUsage int `json:"max_usage"`
		Stats    struct {
			ActiveAnon              int   `json:"active_anon"`
			ActiveFile              int   `json:"active_file"`
			Cache                   int   `json:"cache"`
			Dirty                   int   `json:"dirty"`
			HierarchicalMemoryLimit int64 `json:"hierarchical_memory_limit"`
			HierarchicalMemswLimit  int64 `json:"hierarchical_memsw_limit"`
			InactiveAnon            int   `json:"inactive_anon"`
			InactiveFile            int   `json:"inactive_file"`
			MappedFile              int   `json:"mapped_file"`
			Pgfault                 int   `json:"pgfault"`
			Pgmajfault              int   `json:"pgmajfault"`
			Pgpgin                  int   `json:"pgpgin"`
			Pgpgout                 int   `json:"pgpgout"`
			Rss                     int   `json:"rss"`
			RssHuge                 int   `json:"rss_huge"`
			Swap                    int   `json:"swap"`
			TotalActiveAnon         int   `json:"total_active_anon"`
			TotalActiveFile         int   `json:"total_active_file"`
			TotalCache              int   `json:"total_cache"`
			TotalDirty              int   `json:"total_dirty"`
			TotalInactiveAnon       int   `json:"total_inactive_anon"`
			TotalInactiveFile       int   `json:"total_inactive_file"`
			TotalMappedFile         int   `json:"total_mapped_file"`
			TotalPgfault            int   `json:"total_pgfault"`
			TotalPgmajfault         int   `json:"total_pgmajfault"`
			TotalPgpgin             int   `json:"total_pgpgin"`
			TotalPgpgout            int   `json:"total_pgpgout"`
			TotalRss                int   `json:"total_rss"`
			TotalRssHuge            int   `json:"total_rss_huge"`
			TotalSwap               int   `json:"total_swap"`
			TotalUnevictable        int   `json:"total_unevictable"`
			TotalWriteback          int   `json:"total_writeback"`
			Unevictable             int   `json:"unevictable"`
			Writeback               int   `json:"writeback"`
		} `json:"stats"`
		Failcnt int `json:"failcnt"`
		Limit   int `json:"limit"`
	} `json:"memory_stats"`
	BlkioStats struct {
		IoServiceBytesRecursive []struct {
			Major int    `json:"major"`
			Minor int    `json:"minor"`
			Op    string `json:"op"`
			Value int    `json:"value"`
		} `json:"io_service_bytes_recursive"`
		IoServicedRecursive []struct {
			Major int    `json:"major"`
			Minor int    `json:"minor"`
			Op    string `json:"op"`
			Value int    `json:"value"`
		} `json:"io_serviced_recursive"`
		IoQueueRecursive       []interface{} `json:"io_queue_recursive"`
		IoServiceTimeRecursive []interface{} `json:"io_service_time_recursive"`
		IoWaitTimeRecursive    []interface{} `json:"io_wait_time_recursive"`
		IoMergedRecursive      []interface{} `json:"io_merged_recursive"`
		IoTimeRecursive        []interface{} `json:"io_time_recursive"`
		SectorsRecursive       []interface{} `json:"sectors_recursive"`
	} `json:"blkio_stats"`
	PidsStats struct {
		Current int `json:"current"`
	} `json:"pids_stats"`
	Networks struct {
		Eth0 struct {
			RxBytes   int `json:"rx_bytes"`
			RxPackets int `json:"rx_packets"`
			RxErrors  int `json:"rx_errors"`
			RxDropped int `json:"rx_dropped"`
			TxBytes   int `json:"tx_bytes"`
			TxPackets int `json:"tx_packets"`
			TxErrors  int `json:"tx_errors"`
			TxDropped int `json:"tx_dropped"`
		} `json:"eth0"`
	} `json:"networks"`
}
