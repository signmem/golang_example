package main

import (
   "encoding/json"
   "os"
   "sync"
   "fmt"
   "io/ioutil"
)

type TransferConfig struct {
        Enabled  bool     `json:"enabled"`
        Addrs    []string `json:"addrs"`
        Interval int      `json:"interval"`
        Timeout  int      `json:"timeout"`
}

type HeartbeatConfig struct {
        Enabled  bool   `json:"enabled"`
        Addr     string `json:"addr"`
        Interval int    `json:"interval"`
        Timeout  int    `json:"timeout"`
}

type GlobalConfig struct {
        Heartbeat     *HeartbeatConfig  `json:"heartbeat"`
        Transfer      *TransferConfig   `json:"transfer"`
}

var (
        ConfigFile string
        config     *GlobalConfig
        lock       = new(sync.RWMutex)
)

func Config() *GlobalConfig {
        lock.RLock()
        defer lock.RUnlock()
        return config
}

func FileReadAll(filePth string) ([]byte, error) {
    f, err := os.Open(filePth)
    if err != nil {
        return nil, err
    }
    return ioutil.ReadAll(f)
}

func  main () {

    ConfigFile = "cfg.json"
    var c GlobalConfig
    configContent, _ := FileReadAll(ConfigFile)  

    _ = json.Unmarshal([]byte(configContent), &c)

    fmt.Println(c.Heartbeat.Interval)
    fmt.Println(c.Heartbeat.Addr)
    fmt.Println(c.Heartbeat)

    fmt.Println("============")
    fmt.Println(c.Transfer.Addrs)
    fmt.Println(c.Transfer.Addrs[0])
    fmt.Println("============")
    for _, i := range  c.Transfer.Addrs  {
       fmt.Println(i)
    }
}

