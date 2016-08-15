package main

import (
	"runtime"
	"testing"
)

//TestGetMachineInfo :
func TestGetMachineInfo(t *testing.T) {
	//skip if not linux
	if runtime.GOOS != "linux" {
		t.Skipf("Skip if not in linux")
		return
	}

	d := NewDataCollector()
	if d ==
		nil {
		t.Error("Constructor error")
	}

	ret := d.GetMachineInfo()
	if ret == "" {
		t.Error("Get machine info error.")
	}
}

//GetProcessInfo :
func TestGetProcessInfo(t *testing.T) {
	d := NewDataCollector()
	if d == nil {
		t.Error("Constructor error")
	}

	ret := d.GetProcessInfo()
	if ret == "" {
		t.Error("Get process info error.")
	}
}
