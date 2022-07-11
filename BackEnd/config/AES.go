package config

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"github.com/sirupsen/logrus"
	"os"
)

// Key aes-128加密解密  密钥 16位
var Key = []byte(StoreConfig.WebEncrypt.AESKey)

var KeyContext = keyCon()

// Creat 生成加密数据
func Creat(data string) string {
	plaintext := []byte(data)
	// 如果传入加密串的话，plaint 就是传入的字符串
	if len(os.Args) > 1 {
		plaintext = []byte(os.Args[1])
	}

	// 加密字符串
	cfb := cipher.NewCFBEncrypter(KeyContext, Key)
	ciphertext := make([]byte, len(plaintext))
	cfb.XORKeyStream(ciphertext, plaintext)
	date := base64.StdEncoding.EncodeToString(ciphertext)
	logrus.Info("加密的对应数据==>", date, "---对应的字符串数据==>", data)
	return date
}

// Decrypt 解密卡密的数据
func Decrypt(date string) string {
	data, err := base64.StdEncoding.DecodeString(date)
	if err != nil {
		logrus.Error("base64转化后的结果:", err)
		return ""
	}
	cfb := cipher.NewCFBDecrypter(KeyContext, Key)
	plaintextCopy := make([]byte, len(data))
	cfb.XORKeyStream(plaintextCopy, data)
	str := string(plaintextCopy)
	logrus.Info("解密前的数据==>", date, "---转化为字符串的==>", str)
	return str
}

// aes加密构造器
func keyCon() cipher.Block {
	keyText := StoreConfig.WebEncrypt.AESEncrypt
	// 创建加密算法 aes
	c, err := aes.NewCipher([]byte(keyText))
	if err != nil {
		logrus.Error("创建加密算法失败---Error: NewCipher", len(keyText), err)
		return nil
	}
	logrus.Info("已初始化AES加密算法")
	return c
}
