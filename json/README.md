# 目的

>  利用 golang 进行定义，输出 json 格式  
>  golang 其实对 json 格式定义建议使用 struct 格式， 比较清晰，方便


# json 格式
>  常见 json 格式为  list ,  dict 两种

参考下面输出格式

## dict 
	{
		"fstype": "ext4",
		"host": "host1"
	}

## list
	["host1", "host2", "host3", "host4"]

## dict + list
	{
		"fstype": "ext4",
		"host": ["host1", "host2", "host3", "host4"]
	}

## list + dict 
	[{
			"fstype": "ext4",
			"host": "host1"
		},
		{
			"fstype": "ext4",
			"host": "host2"
	}]

## golang 输出 json 方法
### dict

参考 json_001.go


	func makeExtr() map[string]string {           //  定义
	        testExtr := make(map[string]string,0)
	        testExtr["fstype"] = "ext4"
	        testExtr["host"] = "host1"
	        return testExtr
	}
	
	func buildExtr() string  {                     // json 输出
	        hardExtr := makeExtr()
	        jsonExtr, _ := json.Marshal(hardExtr)
	        return fmt.Sprintf(string(jsonExtr))
	}
	

输出结果

	# go run json_001.go
	{"fstype":"ext4","host":"host1"}

###  list 

参考  json_002.go
	
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
	

执行结果如下

	# go run json_002.go
	["host1","host2","host3","host4"]


简单说明

>  json dict 使用 key: value 方法可以通过 map[string]string 进行定义  
>  json  list 直接通过  []string 进行定义
>  最简单直接方法进行 json 格式定义为  map[string]string  
>  但假如 json 格式复杂时候，则很难维护  


换个方式，我们使用 struct 进行重新定义上面例子

参考 json_004
	
	type HostInfo struct {
	        FsType    string `json:"fstype"`
	        HostName  string `json:"host"`
	}
	
	func makeExtr() HostInfo {
	        var testExtr HostInfo
	        testExtr.FsType = "ext4"
	        testExtr.HostName= "host1"
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

运行结果：

	# go run json_003.go
	{"fstype":"ext4","host":"host1"}

看上去比直接定义麻烦，但请参考下面例子

	{
		"name": "monitorinfo",
		"value": [{
			"body": [{
				"metric": "cpu_idel",
				"datasourceId": 11,
				"from": 1587005847,
				"step": 60,
				"source": "123456789abcdefg",
				"to": 1587092247,
				"type": "Augix"
			}],
			"chart": {
				"legend": "myLegend",
				"title": "mytitile"
			},
			"url": "http://www.m3db.com"
		}]
	}

### 思路
> 类似上述 json 格式，必须要通过 struct 进行解决才可以方便后续维护
>  struct 定义时候需要进行分组处理  

参考 example

参考 extra.go

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
	        Value   []M3Body   `json:"body"`
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
	
思路

> 先对每个 dict 分别独立进行分层定义  
>  在获取数据时候分别进行组合  
>  注意一下 slice 的用法 （利用 struct + append 的方法实现）  
