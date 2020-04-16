package main

import (
	"fmt"
	"encoding/json"	
)


func makeExtr() []string {
	testExtr := []string{
		"host1",
		"host2",
		"host3",
		"host4",
	}
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
