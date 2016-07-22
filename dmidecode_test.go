// ADDED BY DROP - https://github.com/matryer/drop (v0.6)
//  source: github.com/dselans/dmidecode (08eabb429b4ad1353e56dd634048b05d811ba062)
//  update: drop -f github.com/dselans/dmidecode
// license:  (see repo for details)

package main

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"testing"
)

const (
	fakeBinary string = "time4soup"
	testDir    string = "test_data"
)

func TestFindBin(t *testing.T) {
	if runtime.GOOS != "linux" {
		t.Skipf("Skip if not in linux")
		return
	}

	dmi := NewDMI()

	if _, err := dmi.FindBin("time4soup"); err == nil {
		t.Skip("Should not be able to find obscure binary")
		return
	}

	bin, findErr := dmi.FindBin("dmidecode")
	if findErr != nil {
		t.Skip("Should be able to find dmidecode. Error: %v", findErr)
		return
	}

	_, statErr := os.Stat(bin)

	if statErr != nil {
		t.Skip("Should be able to lookup found file. Error: %v", statErr)
		return
	}
}

func TestExecDmidecode(t *testing.T) {
	if runtime.GOOS != "linux" {
		t.Skipf("Skip if not in linux")
		return
	}

	dmi := NewDMI()

	if _, err := dmi.ExecDmidecode("/bin/" + fakeBinary); err == nil {
		t.Skip("Should get an error trying to execute a fake binary. Error: %v", err)
		return
	}

	bin, findErr := dmi.FindBin("dmidecode")
	if findErr != nil {
		t.Skip("Should be able to find binary. Error: %v", findErr)
		return
	}

	output, execErr := dmi.ExecDmidecode(bin)

	if execErr != nil {
		t.Errorf("Should not get errors executing '%v'. Error: %v", bin, execErr)
	}

	if len(output) == 0 {
		t.Errorf("Output should not be empty")
	}
}

func TestParseDmidecodeByFile(t *testing.T) {
	dmi := NewDMI()
	byt, err := ioutil.ReadFile("test_data/centos_6.5_64bit_oem.txt")
	if err != nil {
		t.Error("No file")
	}
	//log.Println("byt:", string(byt))
	dmi.ParseDmidecode(string(byt))
	ret, err := dmi.SearchByName("Memory Device")
	if err != nil {
		t.Error("parse error", err)
	}
	log.Println("ret:", ret)

	si, err := dmi.SearchByName("System Information")
	if err != nil {
		t.Error("Parse SI failed.")
	}
	log.Println(si["Serial Number"], si["Family"], si["DMIType"])
	log.Println("SI:", si)

	rSI := SystemInformation{
		Manufacturer: si["Manufacturer"],
		ProductName:  si["Product Name"],
		Version:      si["Version"],
		SerialNumber: si["Serial Number"],
		UUID:         si["UUID"],
		WakeupType:   si["Wakeup Type"],
		SKUNumber:    si["SKU Number"],
		Family:       si["Family"]}

	log.Println(rSI)

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

	log.Println(rBI, bi)
}

func TestParseDmidecode(t *testing.T) {
	if runtime.GOOS != "linux" {
		t.Skipf("Skip if not in linux")
		return
	}

	dmi := NewDMI()

	bin, findErr := dmi.FindBin("dmidecode")
	if findErr != nil {
		t.Skip("Should be able to find binary. Error: %v", findErr)
		return
	}

	output, execErr := dmi.ExecDmidecode(bin)

	if execErr != nil {
		t.Errorf("Should not get errors executing '%v'. Error: %v", bin, execErr)
	}

	if err := dmi.ParseDmidecode(output); err != nil {
		t.Error("Should not receive an error after parsing dmidecode output")
	}

	if len(dmi.Data) == 0 {
		t.Error("Parsed data structure should have more than 0 entries")
	}

	files, globErr := filepath.Glob(testDir + "/*")
	if globErr != nil {
		t.Errorf("Should not receive errors during '%v' glob. Error: %v", testDir, globErr)
	}

	for _, file := range files {
		// Let's clear it out, each iteration (just in case)
		dmi.Data = make(map[string]map[string]string)

		data, readErr := ioutil.ReadFile(file)
		if readErr != nil {
			t.Errorf("Should not receive errors while reading contents of '%v'. Error: %v", file, readErr)
		}

		if err := dmi.ParseDmidecode(string(data)); err != nil {
			t.Errorf("Should not get errors while parsing '%v'. Error:%v", file, err)
		}

		if len(dmi.Data) == 0 {
			t.Errorf("Data length should be larger than 0 after reading '%v'", file)
		}
	}
}

func TestRun(t *testing.T) {
	if runtime.GOOS != "linux" {
		t.Skipf("Skip if not in linux")
		return
	}

	dmi := NewDMI()

	if err := dmi.Run(); err != nil {
		t.Skip("Run() should not return any errors. Error: %v", err)
	}
}

func TestSearchBy(t *testing.T) {
	if runtime.GOOS != "linux" {
		t.Skipf("Skip if not in linux")
		return
	}

	dmi := NewDMI()

	if _, err := dmi.SearchByName("System Information"); err == nil {
		t.Error("Should have received an error when Search ran prior to .Run()")
	}

	if _, err := dmi.SearchByType(1); err == nil {
		t.Error("Should have received an error when Search ran prior to .Run()")
	}

	if _, err := dmi.GenericSearchBy("DMIName", "System Information"); err == nil {
		t.Error("Should have received an error when Search ran prior to .Run()")
	}

	if err := dmi.Run(); err != nil {
		t.Skip("Run() should not return any errors. Error: %v", err)
		return
	}

	byNameData, byNameErr := dmi.SearchByName("System Information")
	if byNameErr != nil {
		t.Errorf("Shouldn't receive errors when searching by name. Error: %v", byNameErr)
	}

	if len(byNameData) == 0 {
		t.Error("Returned data should have more than 0 records")
	}

	byTypeData, byTypeErr := dmi.SearchByType(1)
	if byTypeErr != nil {
		t.Errorf("Shouldn't receive errors when searching by name. Error: %v", byTypeErr)
	}

	if len(byTypeData) == 0 {
		t.Error("Returned data should have more than 0 records")
	}

	genericData, genericErr := dmi.GenericSearchBy("DMIName", "System Information")
	if genericErr != nil {
		t.Errorf("Shouldn't receive errors when searching by name. Error: %v", genericErr)
	}

	if len(genericData) == 0 {
		t.Error("Returned data should have more than 0 records")
	}
}
