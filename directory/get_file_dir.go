package main

import (
	"path"
	"fmt"
	"strings"
)

func winFile() {        

	// windows 下的目录需要完成 "\"  -> "/" 转换， 否则报错

	fileName := "c:\\falcon-agent\\logs\\falcon-agent.log"
        fileLinuxPath := strings.ReplaceAll(fileName, "\\","/")
	dir, _ :=path.Split(fileLinuxPath)
	fmt.Println(dir)
}

func linuxFile() {
        // linux 下可以使用 path.Dir 获取路径 或 path.Split 也可以
        // 注意 path.Dir 不带结尾的 "/" 

	fileName := "/apps/logs/falcon-agent.log"
	fileDir := path.Dir(fileName)
	fmt.Println(fileDir)
	fileDir1, _ := path.Split(fileName)
	fmt.Println(fileDir1)

}

func main() {
	winFile()
	linuxFile()
}
