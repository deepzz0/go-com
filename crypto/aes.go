package crypto

import (
	"crypto/aes"
	"crypto/cipher"
)

// AES-128。key长度：16, 24, 32 bytes 对应 AES-128, AES-192, AES-256
func AesEncrypt(origData, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	blockSize := block.BlockSize()
	origData = PKCS5Padding(origData, blockSize)
	blockMode := cipher.NewCBCEncrypter(block, key[:blockSize])
	// origData = ZeroPadding(origData, block.BlockSize())
	crypted := make([]byte, len(origData))
	// crypted := origData // 根据CryptBlocks方法的说明，如下方式初始化crypted也可以
	blockMode.CryptBlocks(crypted, origData)
	return crypted, nil
}

func AesDecrypt(crypted, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	blockSize := block.BlockSize()
	blockMode := cipher.NewCBCDecrypter(block, key[:blockSize])
	origData := make([]byte, len(crypted))
	// origData := crypted
	blockMode.CryptBlocks(origData, crypted)
	origData = PKCS5UnPadding(origData)
	// origData = ZeroUnPadding(origData)

	return origData, nil
}

////////////////////////////////////////////////////////////////
// example:
// data, err := aes.AesEncrypt([]byte(str), []byte(AES_KEY))
// if err != nil {
// 	fmt.Println(err)
// }
// hexdata := hex.EncodeToString(data)
// fmt.Println(hex.EncodeToString(data))
// data, err = hex.DecodeString(hexdata)
// if err != nil {
// 	fmt.Println(err)
// }
// fmt.Println(string(data))
func AesEcbEncrypt(origData, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	ecb := NewECBEncrypter(block)
	origData = PKCS5Padding(origData, block.BlockSize())
	crypted := make([]byte, len(origData))
	ecb.CryptBlocks(crypted, origData)
	return crypted, nil
}

func AesEcbDecrypt(crypted, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	// blockSize := block.BlockSize()
	ecb := NewECBDecrypter(block)
	origData := make([]byte, len(crypted))
	ecb.CryptBlocks(origData, crypted)
	origData = PKCS5UnPadding(origData)
	return origData, nil
}

type ecb struct {
	b         cipher.Block
	blockSize int
}

func newECB(b cipher.Block) *ecb {
	return &ecb{
		b:         b,
		blockSize: b.BlockSize(),
	}
}

type ecbEncrypter ecb

// NewECBEncrypter returns a BlockMode which encrypts in electronic code book
// mode, using the given Block.
func NewECBEncrypter(b cipher.Block) cipher.BlockMode {
	return (*ecbEncrypter)(newECB(b))
}
func (x *ecbEncrypter) BlockSize() int { return x.blockSize }
func (x *ecbEncrypter) CryptBlocks(dst, src []byte) {
	if len(src)%x.blockSize != 0 {
		panic("crypto/cipher: input not full blocks")
	}
	if len(dst) < len(src) {
		panic("crypto/cipher: output smaller than input")
	}
	for len(src) > 0 {
		x.b.Encrypt(dst, src[:x.blockSize])
		src = src[x.blockSize:]
		dst = dst[x.blockSize:]
	}
}

type ecbDecrypter ecb

// NewECBDecrypter returns a BlockMode which decrypts in electronic code book
// mode, using the given Block.
func NewECBDecrypter(b cipher.Block) cipher.BlockMode {
	return (*ecbDecrypter)(newECB(b))
}
func (x *ecbDecrypter) BlockSize() int { return x.blockSize }
func (x *ecbDecrypter) CryptBlocks(dst, src []byte) {
	if len(src)%x.blockSize != 0 {
		panic("crypto/cipher: input not full blocks")
	}
	if len(dst) < len(src) {
		panic("crypto/cipher: output smaller than input")
	}
	for len(src) > 0 {
		x.b.Decrypt(dst, src[:x.blockSize])
		src = src[x.blockSize:]
		dst = dst[x.blockSize:]
	}
}
