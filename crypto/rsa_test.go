package crypto

import (
	"encoding/base64"
	"fmt"
	"testing"
)

func TestRsaGenKey(t *testing.T) {
	err := RsaGenKey(1024)
	if err != nil {
		t.Error(err)
	}
}

func TestRsaEncryptDecrypt(t *testing.T) {
	var privateKey = []byte(`
-----BEGIN RSA PRIVATE KEY-----
MIICXQIBAAKBgQCwXFGvLHiirKNxGPaOpA9O4Idf0d6UpQwkIncAKU9z7tZ2dkC1
mR7LNodjD8sNpU6jrocWhkZZWtTY1zAaBs3Udwv6wUf0QHnjPZcaQZxuY8pzn1AF
V5VDFl3F+xv8JnBp1ZHraadjeLDUwIkz1KBb69Gn/4h2N1cD7UDoHflbOwIDAQAB
AoGBAJE3s1sWt07b8MkT0RGrLZ+5aj8QRMMJFHI8nthXK8E+jQGGZcoihyS1hc2g
F4bo81P8RefaMNsq29ChgE4uBBTrKr/EAkVUyr8XkCcS1ScfukGpQ/l+kxGXjVkq
dZ7ZIaufv86AimlqyGIF4hsgNShmqZXEuNCu3NsrphBvqMYhAkEAwAyfZ4BPyUQ0
vCX3BZ5Lr8ih+oDQFqIInF00kbwm/voiqYFwqNJjZGgzik49FU57xYBDYU0xxtDX
nggYvMbELwJBAOsWTazYbfQDDiaXHUMIFqZDIFtds6g7UOSwKjkgB8l5LX3Lh8mw
doTzzg83QkJiDxPE89XHpGtzSF87Qe1iOrUCQCnNVZEYy8UaVIQzm04cw4qymBdH
nIOgp1EptHyYQMC1P4A3zYbhrIK5b6aGGyOdHrHBlmkCfXgyEwyx5HiKpz8CQCwv
jv3z2AbLJDfAo3Fb7dXmPAiwPfpa28OAEQ+Xo58Mta41ORqBnmUy5gIaIswTXj4b
ALGnypGfo3Sy0JtroRkCQQCkWrn8K+uwh4IsUXeyHziK2NhsBF1rpNgjJd69zW1U
RgB/VKxCK4oLC7ueOfexvxR4rqUMOK/tkNFsKVW0Kdj+
-----END RSA PRIVATE KEY-----
`)

	var publicKey = []byte(`
-----BEGIN PUBLIC KEY-----
MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQCwXFGvLHiirKNxGPaOpA9O4Idf
0d6UpQwkIncAKU9z7tZ2dkC1mR7LNodjD8sNpU6jrocWhkZZWtTY1zAaBs3Udwv6
wUf0QHnjPZcaQZxuY8pzn1AFV5VDFl3F+xv8JnBp1ZHraadjeLDUwIkz1KBb69Gn
/4h2N1cD7UDoHflbOwIDAQAB
-----END PUBLIC KEY-----
`)

	data, err := RsaEncrypt(publicKey, []byte("github.com/deepzz0"))
	if err != nil {
		t.Error(err)
	}
	fmt.Println("rsa encrypt base64:" + base64.StdEncoding.EncodeToString(data))

	origData, err := RsaDecrpt(privateKey, data)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(string(origData))
}
