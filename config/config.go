package config

import (
	"errors"
	"fmt"
	"go-help/config/encoding"
	"strings"
	"sync"
)

type File struct {
	//文件路径
	FilePath string

	V interface{}
	//类型名称
	DeName string
}

type DecodeType int

const (
//YAML DecodeType
)

type Decoder interface {
	//根据文件解析
	DecodeFile(filePath string, v interface{}) error
	//根据字符串解析
	DecodeStr(str string, v interface{}) error
}

//增加解析器入口
var (
	unmarshalMethods = map[string]Decoder{
		"json": encoding.NewJson(),
		"yaml": encoding.NewYaml(),
	}

	mMutex sync.RWMutex
)

//UnmarshalMethodsSet 添加解析器
func DecoderReg(deName string, de Decoder) {
	mMutex.Lock()
	defer mMutex.Unlock()
	unmarshalMethods[deName] = de
}

//UnmarshalMethodsGet 添加解析器
func decoderGet(deName string) (Decoder, error) {
	mMutex.RLock()
	defer mMutex.RUnlock()
	if v, ok := unmarshalMethods[deName]; ok {
		return v, nil
	}
	return nil, errors.New(deName + " does not exist")
}

//UnmarshalFile 配置文件格式解析
//列：config.UnmarshalFile("./fileName.json", &data)
//会根据文件名称中的.json判定需要什么类型解析器
func DecodeFile(fileName string, v interface{}, deName string) error {
	if fileName == "" {
		return errors.New("config DecodeFile: the file name cannot be empty")
	}

	fileNameStrs := strings.Split(strings.Trim(fileName, "."), ".")

	if len(fileNameStrs) <= 1 {
		return errors.New("config DecodeFile: incorrect file name")
	}

	decoder, err := decoderGet(deName)
	if err != nil {
		return err
	}

	return decoder.DecodeFile(fileName, v)
}

//DecodeStr 根据字符串解析
func DecodeStr(str string, v interface{}, deName string) error {
	if str == "" {
		return fmt.Errorf("config DecodeStr: str empty")
	}

	decoder, err := decoderGet(deName)
	if err != nil {
		return err
	}

	return decoder.DecodeStr(str, v)
}
