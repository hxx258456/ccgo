// Copyright 2022 s1ren@github.com/hxx258456.

package sm4

import (
	"bytes"
	"crypto/cipher"
	"crypto/rand"
	"errors"
	"fmt"
	"io"
)

// PKCS7Padding 根据pkcs7标准填充明文
func PKCS7Padding(src []byte) []byte {
	padding := BlockSize - len(src)%BlockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(src, padtext...)
}

// PKCS7UnPadding 根据pkcs7标准去除填充
func PKCS7UnPadding(src []byte) ([]byte, error) {
	length := len(src)
	if length == 0 {
		return nil, errors.New("invalid pkcs7 padding (len(padtext) == 0)")
	}
	unpadding := int(src[length-1])
	if unpadding > BlockSize || unpadding == 0 {
		return nil, fmt.Errorf("invalid pkcs7 padding (unpadding > BlockSize || unpadding == 0). unpadding: %d, BlockSize: %d", unpadding, BlockSize)
	}

	pad := src[len(src)-unpadding:]
	for i := 0; i < unpadding; i++ {
		if pad[i] != byte(unpadding) {
			return nil, errors.New("invalid pkcs7 padding (pad[i] != unpadding)")
		}
	}

	return src[:(length - unpadding)], nil
}

// Sm4EncryptCbc sm4加密，CBC模式
//
//goland:noinspection GoNameStartsWithPackageName
func Sm4EncryptCbc(plainData, key []byte) (iv, encryptData []byte, err error) {
	block, err := NewCipher(key)
	if err != nil {
		return nil, nil, err
	}
	paddedData := PKCS7Padding(plainData)
	encryptData = make([]byte, len(paddedData))
	iv = make([]byte, BlockSize)
	if _, err = io.ReadFull(rand.Reader, iv); err != nil {
		return nil, nil, err
	}
	mode := cipher.NewCBCEncrypter(block, iv)
	mode.CryptBlocks(encryptData, paddedData)
	return
}

// Sm4DecryptCbc sm4解密，CBC模式
//
//goland:noinspection GoNameStartsWithPackageName
func Sm4DecryptCbc(encryptData, key, iv []byte) (plainData []byte, err error) {
	block, err := NewCipher(key)
	if err != nil {
		return nil, err
	}
	// 长度检查
	length := len(encryptData)
	if length < BlockSize || length%BlockSize != 0 {
		return nil, fmt.Errorf("sm4.Sm4DecryptCbc: 密文长度不正确,不是Block字节数的整数倍. Block字节数: [%d]", BlockSize)
	}
	paddedData := make([]byte, len(encryptData))
	mode := cipher.NewCBCDecrypter(block, iv)
	mode.CryptBlocks(paddedData, encryptData)
	plainData, err = PKCS7UnPadding(paddedData)
	if err != nil {
		return nil, err
	}
	return
}

// Sm4EncryptCfb sm4加密，CFB模式
//
//goland:noinspection GoNameStartsWithPackageName
func Sm4EncryptCfb(plainData, key []byte) (iv, encryptData []byte, err error) {
	block, err := NewCipher(key)
	if err != nil {
		return nil, nil, err
	}
	encryptData = make([]byte, len(plainData))
	iv = make([]byte, BlockSize)
	if _, err = io.ReadFull(rand.Reader, iv); err != nil {
		return nil, nil, err
	}
	mode := cipher.NewCFBEncrypter(block, iv)
	mode.XORKeyStream(encryptData, plainData)
	return
}

// Sm4DecryptCfb sm4解密，CFB模式
//
//goland:noinspection GoNameStartsWithPackageName
func Sm4DecryptCfb(encryptData, key, iv []byte) (plainData []byte, err error) {
	block, err := NewCipher(key)
	if err != nil {
		return nil, err
	}
	plainData = make([]byte, len(encryptData))
	mode := cipher.NewCFBDecrypter(block, iv)
	mode.XORKeyStream(plainData, encryptData)
	return
}

// Sm4EncryptOfb sm4加密，OFB模式
//
//goland:noinspection GoNameStartsWithPackageName
func Sm4EncryptOfb(plainData, key []byte) (iv, encryptData []byte, err error) {
	block, err := NewCipher(key)
	if err != nil {
		return nil, nil, err
	}
	encryptData = make([]byte, len(plainData))
	iv = make([]byte, BlockSize)
	if _, err = io.ReadFull(rand.Reader, iv); err != nil {
		return nil, nil, err
	}
	mode := cipher.NewOFB(block, iv)
	mode.XORKeyStream(encryptData, plainData)
	return
}

// Sm4DecryptOfb sm4解密，OFB模式
//
//goland:noinspection GoNameStartsWithPackageName
func Sm4DecryptOfb(encryptData, key, iv []byte) (plainData []byte, err error) {
	block, err := NewCipher(key)
	if err != nil {
		return nil, err
	}
	plainData = make([]byte, len(encryptData))
	mode := cipher.NewOFB(block, iv)
	mode.XORKeyStream(plainData, encryptData)
	return
}

// Sm4EncryptGcm sm4加密，GCM模式
//
//goland:noinspection GoNameStartsWithPackageName
func Sm4EncryptGcm(plainData, key []byte) (nonce, encryptData []byte, err error) {
	block, err := NewCipher(key)
	if err != nil {
		return nil, nil, err
	}
	sm4gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, nil, err
	}
	nonce = make([]byte, sm4gcm.NonceSize())
	_, err = io.ReadFull(rand.Reader, nonce)
	if err != nil {
		return nil, nil, err
	}
	out := sm4gcm.Seal(nonce, nonce, plainData, nil)
	encryptData = out[sm4gcm.NonceSize():]
	return
}

// Sm4DecryptGcm sm4解密，GCM模式
//
//goland:noinspection GoNameStartsWithPackageName
func Sm4DecryptGcm(encryptData, key, nonce []byte) ([]byte, error) {
	block, err := NewCipher(key)
	if err != nil {
		return nil, err
	}
	sm4gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}
	// nonce, ciphertext := data[:sm4gcm.NonceSize()], data[sm4gcm.NonceSize():]
	out, err := sm4gcm.Open(nil, nonce, encryptData, nil)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Sm4EncryptGcmWithNonce sm4加密，GCM模式
//
//goland:noinspection GoNameStartsWithPackageName
func Sm4EncryptGcmWithNonce(plainData, key, nonce, dst []byte) (encryptData []byte, err error) {
	block, err := NewCipher(key)
	if err != nil {
		return nil, err
	}
	sm4gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}
	out := sm4gcm.Seal(dst, nonce, plainData, dst)
	encryptData = out[len(dst):]
	return
}

// Sm4DecryptGcmWithNonce sm4解密，GCM模式
//
//goland:noinspection GoNameStartsWithPackageName
func Sm4DecryptGcmWithNonce(encryptData, key, nonce, dst []byte) ([]byte, error) {
	block, err := NewCipher(key)
	if err != nil {
		return nil, err
	}
	sm4gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}
	out, err := sm4gcm.Open(encryptData[:0], nonce, encryptData, dst)
	if err != nil {
		return nil, err
	}
	return out, nil
}
