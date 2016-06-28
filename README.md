linkerConnector: Linux system dat collector and forward to Spark/Kafka.
======================
[![Build Status](https://travis-ci.org/LinkerNetworks/linkerConnector.svg)](https://travis-ci.org/LinkerNetworks/linkerConnector)

[![](https://goreportcard.com/badge/github.com/LinkerNetworks/linkerConnector)](https://goreportcard.com/report/github.com/LinkerNetworks/linkerConnector)

A tool to get Linux system data via `/proc` refer more detail on [spec](http://www.mjmwired.net/kernel/Documentation/filesystems/proc.txt).


Install
--------------

    go get -u -x github.com/LinkerNetworks/linkerConnector

Usage
---------------------

    linkerConnector  

Options
---------------

- `-w` number of workers. (concurrency), default workers is "25"



