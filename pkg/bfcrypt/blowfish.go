package bfcrypt

import (
	"bytes"
	"crypto/cipher"
	"errors"

	"github.com/shengdoushi/base58"
	"golang.org/x/crypto/blowfish"
)

// Decrypt58 decrypt
func Decrypt58(src string, key []byte) (out []byte, err error) {
	var enc []byte
	enc, err = base58.Decode(src, base58.RippleAlphabet)
	if err != nil {
		return
	}

	return BlowFishECBDecrypt(enc, key)
}

// Encrypt58 encrypt
func Encrypt58(src, key []byte) (out string, err error) {
	if len(src) == 0 || len(key) == 0 {
		err = errors.New("input is empty")
		return
	}
	var enc []byte
	enc, err = BlowFishECBEncrypt(src, key)
	if err != nil {
		return
	}
	out = base58.Encode(enc, base58.RippleAlphabet)
	return
}

// Decrypt58 decrypt
func Decrypt(src, key []byte) (out []byte, err error) {
	return BlowFishECBDecrypt(src, key)
}

// Encrypt encrypt
func Encrypt(src, key []byte) (out []byte, err error) {
	return BlowFishECBEncrypt(src, key)
}

// BlowFishECBEncrypt  Encrypt58
func BlowFishECBEncrypt(src, key []byte) ([]byte, error) {
	block, err := blowfish.NewCipher(key)
	if err != nil {
		return nil, err
	}
	if len(src) == 0 {
		return nil, errors.New("plain content empty")
	}
	ecb := NewECBEncrypter(block)

	content := PKCS5Padding(src, block.BlockSize())
	crypted := make([]byte, len(content))
	ecb.CryptBlocks(crypted, content)
	return crypted, nil
}

// BlowFishECBDecrypt Decrypt58
func BlowFishECBDecrypt(crypted, key []byte) ([]byte, error) {
	block, err := blowfish.NewCipher([]byte(key))
	if err != nil {
		return nil, err
	}
	blockMode := NewECBDecrypter(block)
	origData := make([]byte, len(crypted))
	blockMode.CryptBlocks(origData, crypted)
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

type ECBEncrypter ecb

// NewECBEncrypter returns a BlockMode which encrypts in electronic code book
// mode, using the given Block.
func NewECBEncrypter(b cipher.Block) cipher.BlockMode {
	return (*ECBEncrypter)(newECB(b))
}
func (x *ECBEncrypter) BlockSize() int { return x.blockSize }
func (x *ECBEncrypter) CryptBlocks(dst, src []byte) {
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

type ECBDecrypter ecb

// NewECBDecrypter returns a BlockMode which decrypts in electronic code book
// mode, using the given Block.
func NewECBDecrypter(b cipher.Block) cipher.BlockMode {
	return (*ECBDecrypter)(newECB(b))
}

func (x *ECBDecrypter) BlockSize() int {
	return x.blockSize
}

func (x *ECBDecrypter) CryptBlocks(dst, src []byte) {
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

// PKCS5Padding _
func PKCS5Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

// PKCS5UnPadding _
func PKCS5UnPadding(origData []byte) []byte {
	length := len(origData)
	// 去掉最后一个字节 unpadding 次
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}
