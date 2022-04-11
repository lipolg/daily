package service

import (
	"crypto/sha1"
	"fmt"
	"sort"
)

type VerifyService struct {
	Signature string
	Timestamp string
	Nonce     string
	Echostr   string
}

const TOKEN = "lipolg"

func (s VerifyService) Verify() string {
	strs := sort.StringSlice{TOKEN, s.Timestamp, s.Nonce}
	sort.Strings(strs)
	str := ""
	for _, v := range strs {
		str += v
	}
	h := sha1.New()
	h.Write([]byte(str))
	result := fmt.Sprintf("%x", h.Sum(nil))
	if result != s.Signature {
		return "签名校验失败！"
	}
	return s.Echostr
}
