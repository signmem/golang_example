package main

import (
   "fmt"
   "strings"
)


type MetricValue struct {
        Endpoint  string      `json:"endpoint"`
        Metric    string      `json:"metric"`
        Value     interface{} `json:"value"`
        Step      int64       `json:"step"`
        Type      string      `json:"counterType"`
        Tags      string      `json:"tags"`
        Timestamp int64       `json:"timestamp"`
}


func GaugeValue(metric string, val interface{}, tags ...string) *MetricValue {
        return NewMetricValue(metric, val, "GAUGE", tags...)
}


func NewMetricValue(metric string, val interface{}, dataType string, tags ...string) *MetricValue {
        mv := MetricValue{
                Metric: metric,
                Value:  val,
                Type:   dataType,
        }

        size := len(tags)

        if size > 0 {
                mv.Tags = strings.Join(tags, ",")
        }

        return &mv
}



func main() {

    var metric string = "terrytsang"
    var val interface{} = "gogogo"
    tags := []string{"t1","t2","t3","t4"}
    a := GaugeValue(metric, val, tags...)
    fmt.Println(a)

}
