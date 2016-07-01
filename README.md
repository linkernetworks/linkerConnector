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
```    

Options
---------------

- `server`/`s` : The comma separated list of server could be brokers in the Kafka cluster or spark address.
- `topic`/`t` : The topic to kafka produce.
- `interval`/`i` : Time interval(second) to retrieval data , default 0 is not repeat.
- `dest`/`d` : Destination to `kafka`, `spark` and `stdout`.
- `--help` : Display detail command help.


TODOs
---------------

- [ ] DataCollect
	- [X] ProcessInfo
		- [X] Proc Information
		- [X] File Info
	- [ ] MachineIndo
		- [X] CPUInfo
		- [X] MemInfo
		- [ ] NetInfo
		- [ ] DiskInfo
- [ ] Data Forwarding
	- [X] Kafka
	- [X] Spark Streaming
	- [X] JSON file storage
	- [ ] Google Cloud Storage
- [ ] Data Generator
- [ ] Randomize data generator