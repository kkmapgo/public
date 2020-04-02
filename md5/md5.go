package md5

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io"
	"os"
)

// FileMd5 对文件进行md5加密
func FileMd5(fileName string) (string, error) {
	fileObj, err := os.Open(fileName)
	if err != nil {
		fmt.Printf("open file failed, err:%v\n", err)
		return "", err
	}
	hash := md5.New()
	_, err = io.Copy(hash, fileObj)
	if err != nil {
		fmt.Printf("copy to md5 obj failed, err:%v\n", err)
		return "", err
	}
	return hex.EncodeToString(hash.Sum(nil)), nil
}

// StringMd5 对字符串进行md5加密
func StringMd5(str string) (string, error) {
	hash := md5.New()
	_, err := hash.Write([]byte(str))
	if err != nil {
		fmt.Printf("write string to md5 obj failed, err:%v\n", err)
		return "", err
	}
	return hex.EncodeToString(hash.Sum(nil)), nil
}
