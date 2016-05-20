## useragent

The project used to parse the useragent, You can easily statistical data you want.

## Usage

###### Download and install

```
go get github.com/deepzz0/go-common
```

###### Call

```
func ParseByString(useragent string) UserAgent

func ParseByRequest(request *http.Request) UserAgent
```
###### Return params

``` go
type UserAgent struct {
	Type   string `json:"类型"`
	Device struct {
		Type     string `json:"类型"`
		Producer string `json:"厂家"`
		Model    string `json:"型号"`
	} `json:"设备"`
	Client map[string]string `json:"客户端"`
	OS     struct {
		Name    string `json:"名称"`
		Version string `json:"版本号"`
	} `json:"操作系统"`
	Robot struct {
		Name     string `json:"名称"`
		URL      string `json:"网址"`
		Producer struct {
			Name string `json:"名称"`
			URL  string `json:"网址"`
		} `json:"厂家"`
	} `json:"机器人"`
	Vendor string `json:"供应商"`
}

// client 常量如下：
const (
	// TYPE     = "type"
	// PRODUCER = "producer"
	// MODEL    = "model"
	// NAME     = "name"
	// VERSION  = "version"
	// URL      = "url"
	// SUB_TYPE = "sub_type"
	// ENGINE   = "engine"
	TYPE     = "类型"
	PRODUCER = "厂家"
	MODEL    = "型号"
	NAME     = "名称"
	VERSION  = "版本号"
	URL      = "网址"
	SUB_TYPE = "详细类型"
	ENGINE   = "引擎"
)
```

###### Example

``` go
package main

import(
	"fmt"

	"github.com/deepzz0/go-common/useragent"
)

func main(){
	var str = "Mozilla/5.0 (iPhone; CPU iPhone OS 5_0 like Mac OS X) AppleWebKit/534.46 (KHTML, like Gecko) Version/5.1 Mobile/9A334 Safari/7534.48.3"
	agent := useragent.ParseByString(str)
	fmt.Printf("%#v\n", *agent)
}
```
或者

``` go
package main

import(
	"fmt"
	"net/http"

	"github.com/deepzz0/go-common/useragent"
)

func main(){
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		agent := useragent.ParseByRequest(r)
		fmt.Printf("%#v\n", *agent)
	})
	http.ListenAndServe(":8080", nil)
}
```
