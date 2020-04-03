package crypto

import (
	"fmt"

	"github.com/speps/go-hashids"
)

// Encrypt 加密
func Encrypt(key string, minLength int, params []int) (string, error) {
	data := hashids.NewData()
	data.Salt = key
	data.MinLength = minLength
	hash, err := hashids.NewWithData(data)
	if err != nil {
		fmt.Printf("hash failed, err:%v\n", err)
		return "", err
	}
	encode, err := hash.Encode(params)
	if err != nil {
		fmt.Printf("encode failed, err:%v\n", err)
		return "", err
	}
	return encode, nil
}

// Decrypt 解密
func Decrypt(key string, minLength int, param string) ([]int, error) {
	data := hashids.NewData()
	data.Salt = key
	data.MinLength = minLength
	hash, err := hashids.NewWithData(data)
	if err != nil {
		fmt.Printf("hash failed, err:%v\n", err)
		return nil, err
	}
	ints, err := hash.DecodeWithError(param)
	if err != nil {
		fmt.Printf("decode failed, err:%v\n", err)
		return nil, err
	}
	return ints, nil
}
