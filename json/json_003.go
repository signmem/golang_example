package main

import (
	"fmt"
	"encoding/json"	
)

type HostInfo struct {
	FsType    string `json:"fstype"`
	HostName  string `json:"host"`
}

func makeExtr() HostInfo {
	var testExtr HostInfo
	testExtr.FsType = "ext4"
	testExtr.HostName= "host1"
	return testExtr
}

func buildExtr() string  {
	hardExtr := makeExtr()
	jsonExtr, _ := json.Marshal(hardExtr)
	return fmt.Sprintf(string(jsonExtr))
}

func main() {
	info := buildExtr()
	fmt.Println(info)
}
