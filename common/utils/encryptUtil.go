package utils

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"strings"
)

//对明文进行填充,ZeroPadding/NOPadding
func Padding(plainText []byte, blockSize int) []byte {
	//计算要填充的长度
	n := blockSize - len(plainText)%blockSize
	//对原来的明文填充n个0
	temp := bytes.Repeat([]byte{byte(0)}, n)
	plainText = append(plainText, temp...)
	return plainText
}

//对密文删除填充
func UnPadding(cipherText []byte) []byte {
	//取出密文最后一个字节end
	end := cipherText[len(cipherText)-1]
	//删除填充
	cipherText = cipherText[:len(cipherText)-int(end)]
	return cipherText
}

//AEC加密（CBC模式）
func AES_CBC_Encrypt(plainText []byte, key []byte) []byte {
	//指定加密算法，返回一个AES算法的Block接口对象
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}
	//进行填充
	plainText = Padding(plainText, block.BlockSize())
	//指定初始向量vi,长度和block的块尺寸一致
	iv := key
	//指定分组模式，返回一个BlockMode接口对象
	blockMode := cipher.NewCBCEncrypter(block, iv)
	//加密连续数据库
	cipherText := make([]byte, len(plainText))
	blockMode.CryptBlocks(cipherText, plainText)
	//返回密文
	return cipherText
}

//AEC解密（CBC模式）
func AES_CBC_Decrypt(cipherText []byte, key []byte) []byte {
	//指定解密算法，返回一个AES算法的Block接口对象
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}
	//指定初始化向量IV,和加密的一致
	iv := []byte(key)
	//指定分组模式，返回一个BlockMode接口对象
	blockMode := cipher.NewCBCDecrypter(block, iv)
	//解密
	plainText := make([]byte, len(cipherText))
	blockMode.CryptBlocks(plainText, cipherText)
	//删除填充
	plainText = UnPadding(plainText)
	return plainText
}

func MD5_ENCRYPT(s []byte) string {
	m := md5.New()
	// 去除字节数组后面的0
	if bytes.IndexByte(s, 0) > -1 {
		m.Write(s[:bytes.IndexByte(s, 0)])
	} else {
		m.Write(s)
	}
	return hex.EncodeToString(m.Sum(nil))
}

func main() {
	var aeskey = []byte("1234567887654321")
	pass := []byte("mccpwd")
	xpass := AES_CBC_Encrypt(pass, aeskey)

	pass64 := base64.StdEncoding.EncodeToString(xpass)
	fmt.Printf("加密后:%v\n", pass64)

	//bytesPass, err := base64.StdEncoding.DecodeString(pass64)
	bytesPass, err := base64.StdEncoding.DecodeString("PNlwSCDOLpqb6tW40QSbNA==")
	if err != nil {
		fmt.Println(err)
		return
	}

	tpass := AES_CBC_Decrypt(bytesPass, aeskey)
	fmt.Printf("解密后:%s, %d\n", tpass, len(tpass))
	fmt.Printf("MD5加密后:%s\n", strings.ToUpper(MD5_ENCRYPT(tpass)))
}
