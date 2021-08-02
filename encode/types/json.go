package types

import (
	"fmt"
	jsoniter "github.com/json-iterator/go"
	"os"
)

type Json struct {
	jsonIns jsoniter.API
}

func NewJson() Json {
	return Json{
		jsonIns: jsoniter.ConfigCompatibleWithStandardLibrary,
	}
}

func (j Json) DecodeFile(filePath string, v interface{}) error {
	file, err := os.Open(filePath)
	if err != nil {
		return fmt.Errorf("json decode file err:%s", err)
	}

	err = j.jsonIns.NewDecoder(file).Decode(v)
	if err != nil {
		return fmt.Errorf("json decode file err:%s", err)
	}
	return nil
}

func (j Json) DecodeStr(str string, v interface{}) error {
	return j.jsonIns.UnmarshalFromString(str, v)
}

func (j Json) DecodeByte(b []byte, v interface{}) error {
	return j.jsonIns.Unmarshal(b, v)
}

func (j Json) EncodeStr(b interface{}) (result string, err error) {
	r, err := j.jsonIns.Marshal(b)
	if err != nil {
		return "", fmt.Errorf("json encode str err:%s", err)
	}
	return string(r), nil
}

func (j Json) EncodeBytes(b interface{}) (result []byte, err error) {
	r, err := j.jsonIns.Marshal(b)
	if err != nil {
		return []byte{}, fmt.Errorf("json encode bytes err:%s", err)
	}
	return r, nil
}
