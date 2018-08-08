package crypto

import (
	"crypto/md5"
	"io"
	"fmt"
)

const SALT = "salt4shorturl"

func GetMD5(url string) string {
	h := md5.New()
	io.WriteString(h, url+SALT)
	urlmd5 := fmt.Sprintf("%x", h.Sum(nil))
	return urlmd5
}
