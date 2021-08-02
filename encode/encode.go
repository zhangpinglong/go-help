package encode

import (
	"go-help/encode/types"
	"sync"
)

func init() {
	decoderReg(JsonDecode, types.NewJson())
	decoderReg(YamlDecode, types.NewYaml())
}

type DecodeType string

const (
	JsonDecode DecodeType = "json"
	YamlDecode DecodeType = "yaml"
)

type Decoder interface {
	//根据文件解析
	DecodeFile(filePath string, v interface{}) error
	//根据字符串解析
	DecodeStr(str string, v interface{}) error
	//根据byte解析
	DecodeByte(b []byte, v interface{}) error

	//解析成字符串
	EncodeStr(b interface{}) (result string, err error)
	//解析成byte
	EncodeBytes(b interface{}) (result []byte, err error)
}

//增加解析器入口
var (
	unmarshalMethods = map[DecodeType]Decoder{}
	mMutex           sync.RWMutex
)

//DecoderReg 添加解析器
func decoderReg(deName DecodeType, de Decoder) {
	mMutex.Lock()
	defer mMutex.Unlock()
	unmarshalMethods[deName] = de
}

//DecoderGet 获取解析器
func DecoderGet(deType DecodeType) Decoder {
	mMutex.RLock()
	defer mMutex.RUnlock()
	return unmarshalMethods[deType]
}
