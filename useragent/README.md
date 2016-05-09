## useragent

The project used to parse the useragent, You can easily statistical data you want.

## usage

###### Download and install

```
go get github.com/deepzz0/go-common
```

###### Call

```
func ParseByString(useragent string) *UserAgent

func ParseByRequest(request *http.Request) *UserAgent
```

####### Example

```
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

```
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

###### About UserAgent

```
type UserAgent struct {
	Agent_Type             string // brower (mobile)
	Agent_Name             string // chrome
	Agent_Version          string // 50.0.2661.86
	Agent_Rendering_Engine string // Webkit
	Agent_Producer         string // 生产者
	Agent_Producer_Url     string // 网址
	OS_Type                string // OS X
	OS_Name                string // iPhone OS 10.6.6
	OS_Language            string // en-US
	OS_Encryption          string // 加密等级  N无   I弱   U强
	Device_Type            string // mobile
	Device_Model           string // Lumia 930
}
```