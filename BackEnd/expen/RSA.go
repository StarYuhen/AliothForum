package expen

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"os"
)

// RSAEncrypt RSA_Encrypt RSA加密
// plainText 要加密的数据
// path 公钥匙文件地址
func RSAEncrypt(plainText []byte, path string) []byte {
	// 打开文件
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	// 读取文件的内容
	info, _ := file.Stat()
	buf := make([]byte, info.Size())
	file.Read(buf)
	// pem解码
	block, _ := pem.Decode(buf)
	// x509解码

	publicKeyInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		panic(err)
	}
	// 类型断言
	publicKey := publicKeyInterface.(*rsa.PublicKey)
	// 对明文进行加密
	cipherText, err := rsa.EncryptPKCS1v15(rand.Reader, publicKey, plainText)
	if err != nil {
		panic(err)
	}
	// 返回密文
	return cipherText
}

// RSADecrypt RSA_Decrypt RSA解密
// cipherText 需要解密的byte数据
// path 私钥文件路径
func RSADecrypt(cipherText []byte, private string) (string, error) {
	block, _ := pem.Decode([]byte(private))
	// X509解码
	privateKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return "", err
	}
	// 对密文进行解密
	plainText, err := rsa.DecryptPKCS1v15(rand.Reader, privateKey, cipherText)
	// 返回明文
	return string(plainText), err
}

// func main() {
// 	// 加密
// 	data := []byte("hello world")
// 	encrypt := RSAEncrypt(data, "public.pem")
// 	fmt.Println(string(encrypt))
//
// 	// 解密
// 	decrypt := RSADecrypt(encrypt, "private.pem")
// 	fmt.Println("解密后的内容", string(decrypt))
// }
