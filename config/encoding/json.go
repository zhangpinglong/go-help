package encoding

import (
	"fmt"
	jsoniter "github.com/json-iterator/go"
	"go-help/config"
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
	jsoniter.Unmarshal()
	err = j.jsonIns.NewDecoder(file).Decode(v)
	if err != nil {
		return fmt.Errorf("json decode file err:%s", err)
	}
	return nil
}

func (j Json) DecodeStr(str string, v interface{}) error {
	return j.jsonIns.UnmarshalFromString(str, v)
}

func (j Json) DecodeFiles(files []*config.Files) {

}
