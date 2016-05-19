package useragent

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

var ConfigCache = NewCache()

var gopath = os.Getenv("GOPATH")
var wdpath = "github.com/deepzz0/go-common/useragent/"

type Cache struct {
	Robots       []Robot
	OSs          []OS
	Vendors      []Vendor
	Apps         []Client
	Browsers     []Browser
	Engines      []Engine
	Libraries    []Client
	Pims         []Client
	MediaPlayers []Client
	Readers      []Reader
	Cameras      []Device
	CarBrowsers  []Device
	Consoles     []Device
	Medias       []Device
	Mobiles      []Device
	TVs          []Device
}

func NewCache() *Cache {
	cache := &Cache{}
	cache.loadRobots()
	cache.loadOSs()
	cache.loadVendor()
	cache.loadApps()
	cache.loadBrowser()
	cache.loadEngines()
	cache.loadLibraries()
	cache.loadPims()
	cache.loadMeidaPlayers()
	cache.loadReaders()
	cache.loadCameras()
	cache.loadCarBrowser()
	cache.loadConsole()
	cache.loadMedia()
	cache.loadMobile()
	cache.loadTV()
	return cache
}

func (this *Cache) loadRobots() {
	f, err := os.Open(fmt.Sprintf("%s/src/%s/%s", gopath, wdpath, "robot.json"))
	if err != nil {
		checkErr(err)
	}
	data, err := ioutil.ReadAll(f)
	if err != nil {
		checkErr(err)
	}
	err = json.Unmarshal(data, &this.Robots)
	if err != nil {
		checkErr(err)
	}
}

func (this *Cache) loadOSs() {
	f, err := os.Open(fmt.Sprintf("%s/src/%s/%s", gopath, wdpath, "os.json"))
	if err != nil {
		checkErr(err)
	}
	data, err := ioutil.ReadAll(f)
	if err != nil {
		checkErr(err)
	}
	err = json.Unmarshal(data, &this.OSs)
	if err != nil {
		checkErr(err)
	}
}

func (this *Cache) loadVendor() {
	f, err := os.Open(fmt.Sprintf("%s/src/%s/%s", gopath, wdpath, "vendorfragment.json"))
	if err != nil {
		checkErr(err)
	}
	data, err := ioutil.ReadAll(f)
	if err != nil {
		checkErr(err)
	}
	temp := make(map[string][]string)
	err = json.Unmarshal(data, &temp)
	if err != nil {
		checkErr(err)
	}
	for k, v := range temp {
		ven := Vendor{}
		ven.Producer = k
		for _, value := range v {
			ven.Regex = append(ven.Regex, value)
		}
		this.Vendors = append(this.Vendors, ven)
	}
}

func (this *Cache) loadApps() {
	f, err := os.Open(fmt.Sprintf("%s/src/%s/%s", gopath, wdpath, "client/app.json"))
	if err != nil {
		checkErr(err)
	}
	data, err := ioutil.ReadAll(f)
	if err != nil {
		checkErr(err)
	}
	err = json.Unmarshal(data, &this.Apps)
	if err != nil {
		checkErr(err)
	}
}

func (this *Cache) loadBrowser() {
	f, err := os.Open(fmt.Sprintf("%s/src/%s/%s", gopath, wdpath, "client/browser.json"))
	if err != nil {
		checkErr(err)
	}
	data, err := ioutil.ReadAll(f)
	if err != nil {
		checkErr(err)
	}
	err = json.Unmarshal(data, &this.Browsers)
	if err != nil {
		checkErr(err)
	}
}

func (this *Cache) loadEngines() {
	f, err := os.Open(fmt.Sprintf("%s/src/%s/%s", gopath, wdpath, "client/engine.json"))
	if err != nil {
		checkErr(err)
	}
	data, err := ioutil.ReadAll(f)
	if err != nil {
		checkErr(err)
	}
	err = json.Unmarshal(data, &this.Engines)
	if err != nil {
		checkErr(err)
	}
}

func (this *Cache) loadLibraries() {
	f, err := os.Open(fmt.Sprintf("%s/src/%s/%s", gopath, wdpath, "client/library.json"))
	if err != nil {
		checkErr(err)
	}
	data, err := ioutil.ReadAll(f)
	if err != nil {
		checkErr(err)
	}
	err = json.Unmarshal(data, &this.Libraries)
	if err != nil {
		checkErr(err)
	}
}

func (this *Cache) loadPims() {
	f, err := os.Open(fmt.Sprintf("%s/src/%s/%s", gopath, wdpath, "client/pim.json"))
	if err != nil {
		checkErr(err)
	}
	data, err := ioutil.ReadAll(f)
	if err != nil {
		checkErr(err)
	}
	err = json.Unmarshal(data, &this.Pims)
	if err != nil {
		checkErr(err)
	}
}

func (this *Cache) loadMeidaPlayers() {
	f, err := os.Open(fmt.Sprintf("%s/src/%s/%s", gopath, wdpath, "client/player.json"))
	if err != nil {
		checkErr(err)
	}
	data, err := ioutil.ReadAll(f)
	if err != nil {
		checkErr(err)
	}
	err = json.Unmarshal(data, &this.MediaPlayers)
	if err != nil {
		checkErr(err)
	}
}

func (this *Cache) loadReaders() {
	f, err := os.Open(fmt.Sprintf("%s/src/%s/%s", gopath, wdpath, "client/reader.json"))
	if err != nil {
		checkErr(err)
	}
	data, err := ioutil.ReadAll(f)
	if err != nil {
		checkErr(err)
	}
	err = json.Unmarshal(data, &this.Readers)
	if err != nil {
		checkErr(err)
	}
}

func (this *Cache) loadCameras() {
	f, err := os.Open(fmt.Sprintf("%s/src/%s/%s", gopath, wdpath, "device/camera.json"))
	if err != nil {
		checkErr(err)
	}
	data, err := ioutil.ReadAll(f)
	if err != nil {
		checkErr(err)
	}
	err = json.Unmarshal(data, &this.Cameras)
	if err != nil {
		checkErr(err)
	}
}

func (this *Cache) loadCarBrowser() {
	f, err := os.Open(fmt.Sprintf("%s/src/%s/%s", gopath, wdpath, "device/car_browser.json"))
	if err != nil {
		checkErr(err)
	}
	data, err := ioutil.ReadAll(f)
	if err != nil {
		checkErr(err)
	}
	err = json.Unmarshal(data, &this.CarBrowsers)
	if err != nil {
		checkErr(err)
	}
}

func (this *Cache) loadConsole() {
	f, err := os.Open(fmt.Sprintf("%s/src/%s/%s", gopath, wdpath, "device/console.json"))
	if err != nil {
		checkErr(err)
	}
	data, err := ioutil.ReadAll(f)
	if err != nil {
		checkErr(err)
	}
	err = json.Unmarshal(data, &this.Consoles)
	if err != nil {
		checkErr(err)
	}
}

func (this *Cache) loadMedia() {
	f, err := os.Open(fmt.Sprintf("%s/src/%s/%s", gopath, wdpath, "device/media.json"))
	if err != nil {
		checkErr(err)
	}
	data, err := ioutil.ReadAll(f)
	if err != nil {
		checkErr(err)
	}
	err = json.Unmarshal(data, &this.Medias)
	if err != nil {
		checkErr(err)
	}
}

func (this *Cache) loadMobile() {
	f, err := os.Open(fmt.Sprintf("%s/src/%s/%s", gopath, wdpath, "device/mobile.json"))
	if err != nil {
		checkErr(err)
	}
	data, err := ioutil.ReadAll(f)
	if err != nil {
		checkErr(err)
	}
	err = json.Unmarshal(data, &this.Mobiles)
	if err != nil {
		checkErr(err)
	}
}

func (this *Cache) loadTV() {
	f, err := os.Open(fmt.Sprintf("%s/src/%s/%s", gopath, wdpath, "device/tv.json"))
	if err != nil {
		checkErr(err)
	}
	data, err := ioutil.ReadAll(f)
	if err != nil {
		checkErr(err)
	}
	err = json.Unmarshal(data, &this.TVs)
	if err != nil {
		checkErr(err)
	}
}
