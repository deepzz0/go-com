## useragent

The project used to parse the useragent, You can easily statistical data you want. website: [useragent parser](http://blog.deepzz.com/plugin/useragent.html)

## Usage

###### Download and install

```
go get github.com/deepzz0/go-com
```

###### Call

```
func ParseByString(useragent string) UserAgent

func ParseByRequest(request *http.Request) UserAgent
```
###### Return params

``` go
type UserAgent struct {
	Type   string
	Device struct {
		Type     string
		Producer string
		Model    string
	}
	Client map[string]string
	OS     struct {
		Name    string
		Version string
	}
	Robot struct {
		Name     string
		URL      string
		Producer struct {
			Name string
			URL  string
		}
	}
	Vendor string
}
```


###### Example

``` go
package main

import(
	"fmt"

	"github.com/deepzz0/go-com/useragent"
)

func main(){
	var str = "Mozilla/5.0 (iPhone; CPU iPhone OS 5_0 like Mac OS X) AppleWebKit/534.46 (KHTML, like Gecko) Version/5.1 Mobile/9A334 Safari/7534.48.3"
	agent := useragent.ParseByString(str)
	fmt.Printf("%#v\n", agent)
}
```
或者

``` go
package main

import(
	"fmt"
	"net/http"

	"github.com/deepzz0/go-com/useragent"
)

func main(){
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		agent := useragent.ParseByRequest(r)
		fmt.Printf("%#v\n", agent)
	})
	http.ListenAndServe(":8080", nil)
}
```
