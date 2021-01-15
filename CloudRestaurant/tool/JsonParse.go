/*
	提供一些解析结构体的方法
*/
package tool

import (
	"encoding/json"
	"io"
)

type JsonParse struct {
}

func Decode(io io.ReadCloser, v interface{}) error {

	return json.NewDecoder(io).Decode(v)

}
