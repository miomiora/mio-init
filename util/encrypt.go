package util

import (
	"crypto/md5"
	"encoding/hex"
)

func Md5(original string) string {
	hash := md5.New()
	hash.Write([]byte(Secret))
	hash.Write([]byte(original))
	encryptString := hex.EncodeToString(hash.Sum(nil))
	return encryptString
}
