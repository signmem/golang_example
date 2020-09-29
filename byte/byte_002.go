package main

import (
	"fmt"
	"strconv"
	"io/ioutil"
	"time"
)

type TimeStruct struct {
	Hour   string
	Minite string
}

func GetNow() TimeStruct {

	var nowTime TimeStruct
	hour, _   := strconv.Atoi(time.Now().Format("15"))
	minite, _ := strconv.Atoi(time.Now().Format("04"))

	if hour < 10 {
		nowTime.Hour = "0" + strconv.Itoa(hour)
	} else {
		nowTime.Hour = strconv.Itoa(hour)
	}

	if minite < 10 {
		nowTime.Minite = "0" + strconv.Itoa(minite)
	} else {
		nowTime.Minite = strconv.Itoa(minite)
	}

	return nowTime
}

func TimeToFile (timeNow TimeStruct) bool {

	/*
	在使用 ioutil.WriteFile 时候不可以直接输入 string, 需要转换为 []byte

	参考
	func WriteFile(filename string, data []byte, perm os.FileMode) error

	*/

	inputStr := timeNow.Hour + "-" + timeNow.Minite
	
	err := ioutil.WriteFile("/tmp/testfile.txt", []byte(inputStr), 0644)
	if err != nil {
		return false
	}
	return true
}


func main () {
	now := GetNow()
	status :=  TimeToFile(now)
	fmt.Println(status)

}
