/*
	对密码进行加密、解密
	明文加密成密文
	密文解密成明文
*/
package tool

import (
	"crypto/md5"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"io"
)

func EncoderSha256(data string) string {

	h := sha256.New()
	h.Write([]byte(data))
	sum := h.Sum(nil)

	// 由于是十六进制，因此需要转换
	s := hex.EncodeToString(sum)
	return string(s)

}

func Md5(data string) string {
	w := md5.New()
	io.WriteString(w, data)
	bydate := w.Sum(nil)
	result := fmt.Sprintf("%x", bydate)
	return result
}

func Base64Encode(data string) string {
	return base64.StdEncoding.EncodeToString([]byte(data))
}

func Base64Decode(data string) ([]byte, error) {
	return base64.StdEncoding.DecodeString(data)
}
