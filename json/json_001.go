package main

import (
	"fmt"
	"encoding/json"	
)


func makeExtr() map[string]string {
	testExtr := make(map[string]string,0)
	testExtr["fstype"] = "ext4"
	testExtr["host"] = "host1"
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
