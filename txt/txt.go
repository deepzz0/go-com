package txt

import (
	"fmt"
	"os"
	"strings"
)

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
