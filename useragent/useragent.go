// user-agent parser, Used for web page statistics
package useragent

import (
	"github.com/deepzz0/go-com/log"
	"net/http"
	"regexp"
	"strings"
	"sync"
	// "time"
)

const (
	CAMERA      = "camera"
	CAR_BROWSER = "car_browser"
	CONSOLE     = "console"
	MEDIA       = "media"
	MOBILE      = "mobile"
	TV          = "tv"
)

const (
	APP     = "app"
	BROWSER = "browser"
	LIBRARY = "library"
	PIM     = "pim"
	PLAYER  = "player"
	READER  = "reader"
	ROBOT   = "robot"
	VENDOR  = "vendor"
)

const (
	TYPE     = "type"
	PRODUCER = "producer"
	MODEL    = "model"
	NAME     = "name"
	VERSION  = "version"
	URL      = "url"
	SUB_TYPE = "sub_type"
	ENGINE   = "engine"
)

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

func ParseByString(useragent string) UserAgent {
	ua := UserAgent{Client: make(map[string]string)}
	ua.queryRobot(useragent)
	if ua.Robot.Name != "" {
		return ua
	}
	waitgroup := sync.WaitGroup{}
	waitgroup.Add(4)
	go func() {
		ua.queryDevice(useragent)
		waitgroup.Done()
	}()
	go func() {
		ua.queryClient(useragent)
		waitgroup.Done()
	}()
	go func() {
		ua.queryOS(useragent)
		waitgroup.Done()
	}()
	go func() {
		ua.queryVendor(useragent)
		waitgroup.Done()
	}()
	waitgroup.Wait()
	return ua
}

func ParseByRequest(requst *http.Request) UserAgent {
	return ParseByString(requst.UserAgent())
}

func (ua *UserAgent) queryDevice(useragent string) {
	for _, mobile := range ConfigCache.Mobiles {
		reg, err := regexp.Compile(mobile.Regex)
		checkErr(err)
		if str := reg.FindString(useragent); str != "" {
			ua.Device.Producer = mobile.Producer
			ua.Device.Type = MOBILE
			if mobile.Models != nil {
				ua.Device.Model = regexModel(mobile.Models, str)
			} else {
				ua.Device.Model = mobile.Model
			}
			return
		}
	}
	for _, console := range ConfigCache.Consoles {
		reg, err := regexp.Compile(console.Regex)
		checkErr(err)
		if str := reg.FindString(useragent); str != "" {
			ua.Device.Producer = console.Producer
			ua.Device.Type = CONSOLE
			if console.Models != nil {
				ua.Device.Model = regexModel(console.Models, str)
			} else {
				ua.Device.Model = console.Model
			}
			return
		}
	}
	for _, camera := range ConfigCache.Cameras {
		reg, err := regexp.Compile(camera.Regex)
		checkErr(err)
		if str := reg.FindString(useragent); str != "" {
			ua.Device.Producer = camera.Producer
			ua.Device.Type = CAMERA
			if camera.Models != nil {
				ua.Device.Model = regexModel(camera.Models, str)
			} else {
				ua.Device.Model = camera.Model
			}
			return
		}
	}
	for _, car := range ConfigCache.CarBrowsers {
		reg, err := regexp.Compile(car.Regex)
		checkErr(err)
		if str := reg.FindString(useragent); str != "" {
			ua.Device.Producer = car.Producer
			ua.Device.Type = CAR_BROWSER
			if car.Models != nil {
				ua.Device.Model = regexModel(car.Models, str)
			} else {
				ua.Device.Model = car.Model
			}
			return
		}
	}
	for _, media := range ConfigCache.Medias {
		reg, err := regexp.Compile(media.Regex)
		checkErr(err)
		if str := reg.FindString(useragent); str != "" {
			ua.Device.Producer = media.Producer
			ua.Device.Type = MEDIA
			if media.Models != nil {
				ua.Device.Model = regexModel(media.Models, str)
			} else {
				ua.Device.Model = media.Model
			}
			return
		}
	}
	for _, tv := range ConfigCache.TVs {
		reg, err := regexp.Compile(tv.Regex)
		checkErr(err)
		if str := reg.FindString(useragent); str != "" {
			ua.Device.Producer = tv.Producer
			ua.Device.Type = TV
			if tv.Models != nil {
				ua.Device.Model = regexModel(tv.Models, str)
			} else {
				ua.Device.Model = tv.Model
			}
			return
		}
	}
}

func (ua *UserAgent) queryClient(useragent string) {
	for _, app := range ConfigCache.Apps {
		reg, err := regexp.Compile(app.Regex)
		checkErr(err)
		if str := reg.FindString(useragent); str != "" {
			ua.Type = APP

			sli := reg.FindStringSubmatch(str)
			if app.Name == "$1" {
				app.Name = sli[1]
			}
			ua.Client[NAME] = app.Name
			if app.Version == "$1" {
				app.Version = sli[1]
			}
			if app.Version == "$2" {
				app.Version = sli[2]
			}
			ua.Client[VERSION] = app.Version
			return
		}
	}
	for _, browser := range ConfigCache.Browsers {
		reg, err := regexp.Compile(browser.Regex)
		checkErr(err)
		if str := reg.FindString(useragent); str != "" {
			ua.Type = BROWSER
			ua.Client[NAME] = browser.Name

			sli := reg.FindStringSubmatch(str)
			if browser.Version == "$1" {
				browser.Version = sli[1]
			}
			if browser.Engine.Versions != nil {
				for k, v := range browser.Engine.Versions {
					if browser.Version == k {
						ua.Client[ENGINE] = v
					}
				}
			} else {
				ua.Client[ENGINE] = browser.Engine.Default
			}
			if ua.Client[ENGINE] == "" {
				for _, engine := range ConfigCache.Engines {
					reg, err = regexp.Compile(engine.Regex)
					checkErr(err)
					if reg.MatchString(useragent) {
						ua.Client[ENGINE] = engine.Name
						break
					}
				}
			}
			ua.Client[VERSION] = browser.Version
			return
		}
	}

	for _, library := range ConfigCache.Libraries {
		reg, err := regexp.Compile(library.Regex)
		checkErr(err)
		if str := reg.FindString(useragent); str != "" {
			ua.Type = LIBRARY
			ua.Client[NAME] = library.Name
			if library.Version == "$1" {
				sli := reg.FindStringSubmatch(str)
				library.Version = sli[1]
			}
			ua.Client[VERSION] = library.Version
			return
		}
	}
	for _, pim := range ConfigCache.Pims {
		reg, err := regexp.Compile(pim.Regex)
		checkErr(err)
		if str := reg.FindString(useragent); str != "" {
			ua.Type = PIM
			ua.Client[NAME] = pim.Name
			if pim.Version == "$1" {
				sli := reg.FindStringSubmatch(str)
				pim.Version = sli[1]
			}
			ua.Client[VERSION] = pim.Version
			return
		}
	}
	for _, player := range ConfigCache.MediaPlayers {
		reg, err := regexp.Compile(player.Regex)
		checkErr(err)
		if str := reg.FindString(useragent); str != "" {
			ua.Type = PLAYER
			ua.Client[NAME] = player.Name
			if player.Version == "$1" {
				sli := reg.FindStringSubmatch(str)
				player.Version = sli[1]
			}
			ua.Client[VERSION] = player.Version
			return
		}
	}
	for _, reader := range ConfigCache.Readers {
		reg, err := regexp.Compile(reader.Regex)
		checkErr(err)
		if str := reg.FindString(useragent); str != "" {
			ua.Type = READER
			ua.Client[NAME] = reader.Name
			if reader.Version == "$1" {
				sli := reg.FindStringSubmatch(str)
				reader.Version = sli[1]
			}
			ua.Client[VERSION] = reader.Version
			ua.Client[URL] = reader.URL
			ua.Client[SUB_TYPE] = reader.Type
		}
	}
}

func (ua *UserAgent) queryOS(useragent string) {
	for _, os := range ConfigCache.OSs {
		reg, err := regexp.Compile(os.Regex)
		checkErr(err)
		if str := reg.FindString(useragent); str != "" {
			sli := reg.FindStringSubmatch(str)
			if os.Name == "$1" {
				os.Name = sli[1]
			}
			ua.OS.Name = os.Name
			if os.Version == "$1" {
				os.Version = sli[1]
			}
			if os.Version == "$2" {
				os.Version = sli[2]
			}
			ua.OS.Version = os.Version
			return
		}
	}
}

func (ua *UserAgent) queryRobot(useragent string) {
	for _, robot := range ConfigCache.Robots {
		reg, err := regexp.Compile(robot.Regex)
		checkErr(err)
		if reg.MatchString(useragent) {
			ua.Type = ROBOT
			ua.Robot.Name = robot.Name
			ua.Robot.URL = robot.URL
			ua.Robot.Producer.Name = robot.Producer.Name
			ua.Robot.Producer.URL = robot.Producer.URL
			return
		}
	}
}

func (ua *UserAgent) queryVendor(useragent string) {
	for _, ven := range ConfigCache.Vendors {
		for _, regex := range ven.Regex {
			reg, err := regexp.Compile(regex)
			checkErr(err)
			if reg.MatchString(useragent) {
				ua.Type = VERSION
				ua.Vendor = ven.Producer
				return
			}
		}
	}
}

func regexModel(models []Model, str string) (result string) {
	for _, model := range models {
		reg, err := regexp.Compile(model.Regex)
		checkErr(err)
		if reg.MatchString(str) {
			if strings.Contains(model.Model, "$1") {
				sli := reg.FindStringSubmatch(str)
				model.Model = strings.Replace(model.Model, "$1", sli[1], 1)
			}
			result = model.Model
			return
		}
	}
	return
}

func checkErr(err error) {
	if err != nil {
		log.Error(err)
	}
}
