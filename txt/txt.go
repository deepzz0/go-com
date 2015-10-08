package txt

import (
	"fmt"
	"os"
	"strings"
)

func GetValueString(key string) string {
	return ConfigData[key]
}

//Bit sizes 0, 8, 16, 32, and 64 correspond to int, int8, int16, int32, and int64
func GetValueInt64(key string) (int64, error) {
	data := ConfigData[key]
	toint64, err := strconv.ParseInt(data, 10, 64)
	if err != nil {
		return -1, err
	}
	return toint64, nil
}

func GetValueInt(key string) (int, error) {
	data := ConfigData[key]
	toint, err := strconv.ParseInt(data, 10, 0)
	if err != nil {
		return -1, err
	}
	return int(toint), nil
}

func GetValueBool(key string) (bool, error) {
	data := ConfigData[key]
	tobool, err := strconv.ParseBool(data)
	if err != nil {
		return false, err
	}
	return tobool, nil
}

func GetValueFloat32(key string) (float32, error) {
	data := ConfigData[key]
	tofloat64, err := strconv.ParseFloat(data, 32)
	if err != nil {
		return -1.1, err
	}
	return float32(tofloat64), nil
}

func GetValueFloat64(key string) (float64, error) {
	data := ConfigData[key]
	tofloat64, err := strconv.ParseFloat(data, 64)
	if err != nil {
		return -1.1, err
	}
	return tofloat64, nil
}

func ParseConf(filename string) map[string]string {
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
