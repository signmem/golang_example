package main

import (
    "fmt"
    "flag"
    "os"
    "io/ioutil" 
    "encoding/json"
)


type cephGlobalConfig []struct {
	UUID       string `json:"uuid"`
	MonNumber  int    `json:"monNumber"`
	M3WriteURL string `json:"m3writeUrl"`
}


func main() {

    var cfg = flag.String(  "c",  "cfg.json",  "config file location" )
    var version = flag.Bool( "v",  false,  "show version" )
    var ip = flag.Int("f", 1, "help message for flagname")
   
    flag.Parse()

    if *version {
        fmt.Println ("20200930")
        return
    } 

   if *cfg != "" {
        globalConfig := *cfg 
        parseConfig(globalConfig)
        return
   }
   

   if *ip > 1 {
      fmt.Println(*ip)
      return
   }

}

func parseConfig(cfg string) {
   if cfg == "" {
       fmt.Println ( "use -c define config file")
       return
   }

   if !exists(cfg) {

       fmt.Println("config file is not exists")

   }else{
       configInfo , _:= readAll(cfg) 

       var myConfig cephGlobalConfig

       // 读入 cfg.json 后需要调用 json.Unmarshal 放入 struct 结构
       err := json.Unmarshal(configInfo, &myConfig)

       if err != nil {
          fmt.Println("xxxxx")
       }
       for _, i := range myConfig {
           fmt.Println(i.UUID)
       }
   }
}


func exists(path string) bool {  
    _, err := os.Stat(path)    //os.Stat获取文件信息  
    if err != nil {  
        if os.IsExist(err) {  
            return true  
        }  
        return false  
    }  
    return true  
}  

func readAll(filePth string) ([]byte, error) {

    // readAll 返回的是 []byte 格式，可以直接被 json 函数调用

    f, err := os.Open(filePth)
    if err != nil {
        return nil, err
    }
    return ioutil.ReadAll(f)
}

