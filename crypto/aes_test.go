package crypto

import (
	"encoding/base64"
	"fmt"
	"testing"
)

func TestAes(t *testing.T) {
	key := []byte("sfe023f_9fd&fwfl")
	result, err := AesEncrypt([]byte("github.com/deepzz0"), key)
	if err != nil {
		panic(err)
	}
	fmt.Println(base64.StdEncoding.EncodeToString(result))
	origData, err := AesDecrypt(result, key)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(origData))
}
