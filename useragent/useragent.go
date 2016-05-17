// user-agent parser, Used for web page statistics
package useragent

import (
	"net/http"
	"regexp"
	"strings"

	"github.com/deepzz0/go-common/log"
)

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
	Device_Model           string //
}

func ParseByString(useragent string) *UserAgent {
	agent := &UserAgent{}
	if robot := IsRobot(useragent); robot != "" {
		agent.Agent_Type = ROBOT
		agent.Agent_Name = robot
		return agent
	}
	reg, _ := regexp.Compile(`\(.*\)`)
	index := reg.FindStringIndex(useragent)
	if len(index) < 2 {
		log.Warn(useragent)
		return agent
	}
	osInfo := useragent[index[0]:index[1]]
	slice := strings.Split(osInfo, ";")
	agent.OS_Type = slice[0][1:len(slice[0])]
	agent.OS_Encryption = GetOSSecurity(osInfo)
	regLanguage, _ := regexp.Compile(`en-[\w]+`)
	agent.OS_Language = regLanguage.FindString(osInfo)
	switch agent.OS_Type {
	case "iPad", "iPhone", "iPod", "Macintosh":
		regVersion, _ := regexp.Compile(`(?:(Mac OS X [\d\_\.]+)|(iPhone OS [\d\_]+))`)
		agent.OS_Name = strings.Replace(regVersion.FindString(osInfo), "_", ".", -1)
	case "Android":
		regVersion, _ := regexp.Compile(`rv:[\d\.]+`)
		agent.OS_Name = regVersion.FindString(osInfo)
	case "Linux":
		regVersion, _ := regexp.Compile(`Android [\d\.]+`)
		agent.OS_Name = regVersion.FindString(osInfo)
	case "X11":
		regVersion, _ := regexp.Compile(`Linux x[\d]+_[\d]+`)
		agent.OS_Name = regVersion.FindString(osInfo)
	case "compatible", "Windows":
		regVersion, _ := regexp.Compile(`(?:(Windows NT [\d\.]+)|(Windows Phone( OS)? [\d\.]+))`)
		agent.OS_Name = regVersion.FindString(osInfo)
	case "PlayBook":
		regVersion, _ := regexp.Compile(`RIM Tablet OS [\d\.]+`)
		agent.OS_Name = regVersion.FindString(osInfo)
	default:
		if strings.Contains(agent.OS_Type, "Windows NT") {
			agent.OS_Name = agent.OS_Type
		}
	}
	if strings.Contains(useragent, "Mobile") {
		agent.Device_Type = "mobile"
	}
	agent.Device_Model = FindDevice(slice)
	// ------------------------------------------------------------
	browserInfo := useragent[index[0]:]
	var Name, Version string
	// 是否是web
	if name, version := IsWebBrowser(browserInfo); name != "" {
		Name, Version = name, version
		agent.Agent_Type = WEB_BROWSER
	}
	// 是否是
	if name, version := IsMobileBrowser(browserInfo); name != "" {
		Name, Version = name, version
		agent.Agent_Type = MOBILE_BROWSER
	}
	// 是否是
	if name, version := IsTextBrowser(browserInfo); name != "" {
		Name, Version = name, version
		agent.Agent_Type = TEXT_BROWSER
	}
	//
	if name, version := IsEmailClient(browserInfo); name != "" {
		Name, Version = name, version
		agent.Agent_Type = EMAIL_CLIENT
	}
	//
	if name, version := IsTool(browserInfo); name != "" {
		Name, Version = name, version
		agent.Agent_Type = TOOL
	}
	//
	if name, version := IsApp(browserInfo); name != "" {
		Name, Version = name, version
		agent.Agent_Type = APP
	}
	agent.Agent_Name = Name
	agent.Agent_Version = Version
	agent.Agent_Rendering_Engine = GetEngine(agent.Agent_Type, Name)
	agent.Agent_Producer = GetProducer(Name)
	return agent
}

func ParseByRequest(request *http.Request) *UserAgent {
	return ParseByString(request.UserAgent())
}

func IsWebBrowser(useragent string) (string, string) {
	for k, v := range Agent_types[WEB_BROWSER] {
		reg, err := regexp.Compile(k)
		if err != nil {
			log.Error(err)
			return "", ""
		}
		if str := reg.FindString(useragent); str == "" {
			continue
		} else {
			return v, TypeToName[v](str)
		}
	}
	return "", ""
}

func IsMobileBrowser(useragent string) (string, string) {
	for k, v := range Agent_types[MOBILE_BROWSER] {
		reg, err := regexp.Compile(k)
		if err != nil {
			log.Error(err)
			return "", ""
		}
		if str := reg.FindString(useragent); str == "" {
			continue
		} else {
			return v, TypeToName[v](str)
		}
	}
	return "", ""
}

func IsTextBrowser(useragent string) (string, string) {
	for k, v := range Agent_types[TEXT_BROWSER] {
		reg, err := regexp.Compile(k)
		if err != nil {
			log.Error(err)
			return "", ""
		}
		if str := reg.FindString(useragent); str == "" {
			continue
		} else {
			return v, TypeToName[v](str)
		}
		return k, v
	}
	return "", ""
}

func IsEmailClient(useragent string) (string, string) {
	for k, v := range Agent_types[EMAIL_CLIENT] {
		reg, err := regexp.Compile(k)
		if err != nil {
			log.Error(err)
			return "", ""
		}
		if str := reg.FindString(useragent); str == "" {
			continue
		} else {
			return v, TypeToName[v](str)
		}
		return k, v
	}
	return "", ""
}

func IsTool(useragent string) (string, string) {
	for k, v := range Agent_types[TOOL] {
		reg, err := regexp.Compile(k)
		if err != nil {
			log.Error(err)
			return "", ""
		}
		if str := reg.FindString(useragent); str == "" {
			continue
		} else {
			return v, TypeToName[v](str)
		}
		return k, v
	}
	return "", ""
}

func IsApp(useragent string) (string, string) {
	for k, v := range Agent_types[APP] {
		reg, err := regexp.Compile(k)
		if err != nil {
			log.Error(err)
			return "", ""
		}
		if str := reg.FindString(useragent); str == "" {
			continue
		} else {
			return v, TypeToName[v](str)
		}
	}
	return "", ""
}

func IsRobot(useragent string) string {
	for k, v := range Agent_types[ROBOT] {
		reg, err := regexp.Compile(k)
		if err != nil {
			log.Error(err)
			return ""
		}
		if str := reg.FindString(useragent); str == "" {
			continue
		} else {
			return v
		}
	}
	return ""
}

func FindDevice(slice []string) (model string) {
	for _, v := range slice {
		if strings.Contains(v, "Build") {
			model = strings.TrimSpace(v)
		}
		if strings.Contains(model, ")") {
			model = model[:strings.Index(model, ")")]
		}
	}
	return
}

func GetOSSecurity(osinfo string) string {
	reg, _ := regexp.Compile(`(?:U|I|N)\;`)
	str := reg.FindString(osinfo)
	if str != "" {
		str = str[:len(str)-1]
	}
	return str
}

func GetProducer(name string) string {
	if producer := TypeToProducer[name]; producer == "" {
		return "UNKNOWN"
	} else {
		return producer
	}
}

func GetEngine(typ, name string) string {
	if engine := TypeToRenderingEngine[typ][name]; engine == "" {
		return "UNKNOWN"
	} else {
		return engine
	}
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
