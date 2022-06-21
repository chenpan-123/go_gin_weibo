package utils

import (
	"crypto/md5"
	"fmt"
	"time"
)

func MD5(str string) string {
	md5str := fmt.Sprintf("%x", md5.Sum([]byte(str)))
	return md5str
}

func SwitchTimeStampToData(timestamp int64) string {
	t := time.Unix(timestamp, 0)
	return t.Format("2006-01-02 15:04:0")

}
