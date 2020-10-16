package main 

import (
	"fmt"
	_ "path"
	"path/filepath"
	"strings"
	"os"
)


func pathFormat(basePath string) string {
	return strings.ReplaceAll(basePath, "\\", "/")
}

func walk(fp string) {
	var linuxPath = pathFormat(fp)
 
	filepath.Walk(linuxPath, func(path string, info os.FileInfo, err error) error {
		fileFullPath := pathFormat(path)
		fileSuffix := strings.HasSuffix(fileFullPath, ".log")
		if fileSuffix {
			fmt.Println(info.Size())
			fmt.Println(fileFullPath)
		}
		return nil
	})
}

func main() {

   // 用于扫描 path 目录下的所有文件
   var path = "/var/log/"
   walk(path)

}
