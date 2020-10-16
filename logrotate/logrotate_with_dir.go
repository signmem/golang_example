package main

import (
	"io"
	"os"
	"path"
	"path/filepath"
	"strconv"
	"strings"
)

func main() {
	// 主要目标
	// 从配置中的文件中，扫描当前日志目录
	// 把所有大于 10M 的文件执行 5 次日志备份
	logFilePath := "/apps/logs/falcon-agent.log"
	FalconLogRotate
}

func FalconLogRotate(logFilePath string) {
	logFile := logFilePath
	logFileLinuxFormat := pathFormat(logFile)
	logDir, _ := path.Split(logFileLinuxFormat)
	walk(logDir)
}


func pathFormat(basePath string) string {

	// use to compile windows path format 

	return strings.ReplaceAll(basePath, "\\", "/")
}

func walk(fp string) {
	var linuxPath = pathFormat(fp)

	filepath.Walk(linuxPath, func(path string, info os.FileInfo, err error) error {
		fileFullPath := pathFormat(path)

		fileSuffix := strings.HasSuffix(fileFullPath, ".log")
		if fileSuffix && info.Size() > 10000000 {
			logRotate(fileFullPath)
		}
		return nil
	})
}


func logRotate(fileName string ) {

	// use to logrotate
	for num := 4 ; num >= 1; num-- {
		logName := strconv.Itoa(num)
		bakName := strconv.Itoa(num + 1)
		originFileName := fileName + "." + logName
		backupFileName := fileName + "." + bakName
		_ = os.Rename(originFileName, backupFileName)
	}
	// copy file and truncate origin file
	_ = copy(fileName, fileName + ".1")
	_ = os.Truncate(fileName, 0)
}

func copy(src, dst string)  error {
	// copy file procedure
	sourceFileStat, err := os.Stat(src)
	if err != nil {
		return  err
	}

	if !sourceFileStat.Mode().IsRegular() {
		return  err
	}

	source, err := os.Open(src)
	if err != nil {
		return  err
	}
	defer source.Close()

	destination, err := os.Create(dst)
	if err != nil {
		return  err
	}
	defer destination.Close()
	_, err = io.Copy(destination, source)
	if err != nil {
		return  err
	}
	return nil
}
