package text

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type config struct {
	ConfigData map[string]string
}

func NewConfig(filename string) *config {
	c := &config{parseConf(filename)}
	return c
}

func ParseConf(filename string) map[string]string {
	return parseConf(filename)
}

func (c *config) GetValueString(key string) string {
	return c.ConfigData[key]
}

//Bit sizes 0, 8, 16, 32, and 64 correspond to int, int8, int16, int32, and int64
func (c *config) GetValueInt64(key string) (int64, error) {
	data := c.ConfigData[key]
	toint64, err := strconv.ParseInt(data, 10, 64)
	if err != nil {
		return -1, err
	}
	return toint64, nil
}

func (c *config) GetValueInt(key string) (int, error) {
	data := c.ConfigData[key]
	toint, err := strconv.ParseInt(data, 10, 0)
	if err != nil {
		return -1, err
	}
	return int(toint), nil
}

func (c *config) GetValueBool(key string) (bool, error) {
	data := c.ConfigData[key]
	tobool, err := strconv.ParseBool(data)
	if err != nil {
		return false, err
	}
	return tobool, nil
}

func (c *config) GetValueFloat32(key string) (float32, error) {
	data := c.ConfigData[key]
	tofloat64, err := strconv.ParseFloat(data, 32)
	if err != nil {
		return -1.1, err
	}
	return float32(tofloat64), nil
}

func (c *config) GetValueFloat64(key string) (float64, error) {
	data := c.ConfigData[key]
	tofloat64, err := strconv.ParseFloat(data, 64)
	if err != nil {
		return -1.1, err
	}
	return tofloat64, nil
}

func parseConf(filename string) map[string]string {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	fi, err := file.Stat()
	if err != nil {
		panic(err)
	}
	bt := make([]byte, fi.Size())
	_, err = file.Read(bt)
	if err != nil {
		panic(err)
	}
	lines := byteReadLine(bt)
	result := make(map[string]string, len(lines))
	for _, line := range lines {
		if strings.HasPrefix(line, "#") {
			continue
		}
		k, v := parseToMap(line)
		result[k] = v
	}
	return result
}

func byteReadLine(bt []byte) []string {
	var result = make([]string, 0)
	count := -1
	for i, b := range bt {
		if string(b) == "\n" {
			temp := strings.Replace(strings.TrimSpace(string(bt[count+1:i])), "\"", "", -1)
			if temp != "" {
				result = append(result, temp)
				count = i
			}
		}
		if i == len(bt)-1 {
			temp := strings.Replace(strings.TrimSpace(string(bt[count+1:i+1])), "\"", "", -1)
			if temp != "" {
				result = append(result, temp)
				count = i
			}
		}
	}
	return result
}

func parseToMap(str string) (string, string) {
	if !strings.Contains(str, "=") {
		panic(fmt.Errorf("config file incomplete:%s", str))
	}
	sli := strings.Split(str, "=")
	key := strings.Replace(sli[0], " ", "", -1)
	value := strings.Replace(sli[1], " ", "", -1)
	return key, value
}
