package aes

import (
	"fmt"
	"testing"
	"time"
)

var addr = "0x8114B66d176B2BebB84729385709445166D7A463"
var pwd = "testPwd"

var rawData = []byte("RawData")
var rawKey = []byte(addr)[:32]

func TestAes(t *testing.T) {
	str, _ := EncryptByAes(rawData, rawKey)
	str1,_:= DecryptByAes(str, rawKey)
	fmt.Printf("Data: %v, Encrypto：%v, Decrypto：%s\n ",string(rawData), str, str1)
}

func TestAesTime(t *testing.T) {
	startTime := time.Now()
	count := 10000
	for i := 0; i < count; i++ {
		str, _ := EncryptByAes(rawData, rawKey)
		_, _ = DecryptByAes(str, rawKey)
	}
	fmt.Printf("%v次 - %v", count, time.Since(startTime))
}

func TestCachePwd(t *testing.T) {
	if CheckPwdCacheExist(DefaultCacheDirPath, addr) {
		data, err := GetPwdByReadCache(DefaultCacheDirPath, addr)
		if err != nil {
			fmt.Printf("err is %v\n", err)
		}
		fmt.Printf("Read %v, pwd is %v\n", data, []byte(data))
	} else {
		data, err := EncryptByAesAndWriteToFile(DefaultCacheDirPath, addr, pwd)
		if err != nil {
			fmt.Printf("err is %v\n", err)
		}
		fmt.Printf("Write %v, data is %v\n", pwd, data)
	}
}