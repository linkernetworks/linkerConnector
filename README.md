linkerConnector: Linux system data collector and forward to target destination.
======================
[![Build Status](https://travis-ci.org/LinkerNetworks/linkerConnector.svg)](https://travis-ci.org/LinkerNetworks/linkerConnector)

[![](https://goreportcard.com/badge/github.com/LinkerNetworks/linkerConnector)](https://goreportcard.com/report/github.com/LinkerNetworks/linkerConnector)

A tool to get Linux system data via `/proc` refer more detail on [spec](http://www.mjmwired.net/kernel/Documentation/filesystems/proc.txt).


Install
--------------

#### Install if your system installed golang 

```
go get -u -x github.com/LinkerNetworks/linkerConnector
```

#### Install using binary.

TBD

Usage
---------------------

```
linkerConnector  (-server=XXX) (-dest=XXX) ...

## Check all command destription
linkerConnector --help

## optional: using make
make compile-static-amd64: static compile
make docker-build: build a docker image whose name is linkerrepository/linker_connector
make container-run: run the image built above
make container-stop: stop container
make container-clean: remove container
make clean: clean binary file
```    

Options
---------------

- `readProcPath`/`r` : File path of proc for linkerConnector to read from.
- `cAdvisorAddr`/`c` : Http Url for cAdvisor REST API.
- `server`/`s` : The comma separated list of server could be brokers in the Kafka cluster or spark address.
- `topic`/`t` : The topic to kafka produce.
- `interval`/`i` : Time interval(second) to retrieval data , default 0 is not repeat.
- `dest`/`d` : Destination to `kafka`, `spark`, `cassandra` and `stdout`.
- `pipe`/`p`: Using pipe mode to forward data. usage `cat data.txt | linkerConnector -p true`
- `dsiableFileSave` / `d`: Disable local file save.
- `--help` : Display detail command help.

Important Data
---------------

- Disk I/O `/proc/{PID}/io`
- Process Status: (Zombie, VM Data) `/proc/{PID}/status`.
  - `state` (R is running, S is sleeping, D is sleeping in an uninterruptible wait, Z is zombie, T is traced or stopped).
  - Zombie: No parent process.
- How to find parent process:
  -  `ps axjf` (Linux only)
- Hardware Infomation Dump:
  - `dmidecode`, only dump memory `dmidecode -- type memory` 

TODOs
---------------

- [ ] DataCollect
	- [X] ProcessInfo
		- [X] Proc Information
		- [X] File Info
		- [X] Host's all Container Info
	- [ ] MachineIndo
		- [X] CPUInfo
		- [X] MemInfo
		- [X] DMIINFO- Bios Information
		- [X] DMIINFO- System Information
		- [ ] NetInfo
		- [ ] DiskInfo
	- [ ] Last Login history
	- [ ] Syslog
- [ ] Data Forwarding
	- [X] Kafka
	- [X] Spark Streaming
	- [X] Stdout
	- [X] Pipe mode
	- [X] Cassandra
	- [ ] TBC
- [ ] Data Generator
- [ ] Randomize data generator
