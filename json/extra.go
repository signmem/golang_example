package main

import (
    "fmt"
    "encoding/json"
    "time"
)

type MonitorData struct {
	Name  string    `json:"name"`
	Value []M3Value `json:"value"`
}

type M3Body struct {
	Metric       string `json:"metric"`
	DatasourceID int    `json:"datasourceId"`
	From         int64  `json:"from"`
	Step         int    `json:"step"`
	Source       string `json:"source"`
	To           int64  `json:"to"`
	Type         string `json:"type"`
}

type M3Chart struct {
	Legend string `json:"legend"`
	Title  string `json:"title"`
}


type M3Value struct {
	Value	[]M3Body   `json:"body"`
	Chart   M3Chart    `json:"chart"`
	URL     string     `json:"url"` 
}

func buildM3Chart() M3Chart{
	var newM3Chart M3Chart
	newM3Chart.Legend = "myLegend"
	newM3Chart.Title = "mytitile"
	return newM3Chart
}

func buildM3Body() M3Body {
	var newBody M3Body
	newBody.Metric = "cpu_idel"
	newBody.DatasourceID = 11
	newBody.From = time.Now().Unix()
	newBody.Step = 60
	newBody.Source = "123456789abcdefg"
	newBody.To = time.Now().Unix() + 86400
	newBody.Type = "Augix"
	return newBody
}

func buildM3Value() M3Value {
	var newM3Value M3Value
	var m3BodySlice []M3Body
	m3BodyData := buildM3Body()
	m3BodySlice = append(m3BodySlice, m3BodyData)
	newM3Value.Value = m3BodySlice
	newM3Value.Chart = buildM3Chart()
	newM3Value.URL = "http://www.m3db.com"
	return newM3Value
}

func buildMonitorData() MonitorData {
	var newMonitorData MonitorData
	var monitorDataValue M3Value
	var monitorDataValueSlice []M3Value
	monitorDataValue = buildM3Value()
	newMonitorData.Name = "monitorinfo"
	monitorDataValueSlice = append(monitorDataValueSlice, monitorDataValue)
	newMonitorData.Value = monitorDataValueSlice
	return newMonitorData
}


func main() {

	jsonInfo := buildMonitorData()
	goJson, _ := json.Marshal(jsonInfo)
	data := fmt.Sprintf( string(goJson))
	fmt.Println(data)

}
