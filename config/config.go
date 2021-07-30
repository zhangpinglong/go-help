package config

import (
	"encoding/json"
	"errors"
	"github.com/BurntSushi/toml"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"strings"
	"sync"
)

type UnmarshalFiler interface {
	UnmarshalFile(fileName string, v interface{}) error
}

type JSONUnmarshalFile struct{}

func (JSONUnmarshalFile) UnmarshalFile(fileName string, v interface{}) error {
	d, err := ioutil.ReadFile(fileName)
	if err != nil {
		return err
	}
	return json.Unmarshal(d, v)
}

type YAMLUnmarshalFile struct{}

func (YAMLUnmarshalFile) UnmarshalFile(fileName string, v interface{}) error {
	d, err := ioutil.ReadFile(fileName)
	if err != nil {
		return err
	}
	return yaml.Unmarshal(d, v)
}

type TOMLUnmarshalFile struct{}

func (TOMLUnmarshalFile) UnmarshalFile(fileName string, v interface{}) error {
	d, err := ioutil.ReadFile(fileName)
	if err != nil {
		return err
	}
	return toml.Unmarshal(d, v)
}

//增加解析器入口
var (
	unmarshalMethods = map[string]interface{}{
		"json": JSONUnmarshalFile{},
		"yaml": YAMLUnmarshalFile{},
		"toml": TOMLUnmarshalFile{},
	}
	mMutex sync.RWMutex
)

//UnmarshalMethodsAdd 添加解析器
func UnmarshalMethodsAdd(deName string, de UnmarshalFiler) error {
	mMutex.Lock()
	defer mMutex.Unlock()
	if _, ok := unmarshalMethods[deName]; ok {
		return errors.New("unmarshalFiler already exist")
	}
	unmarshalMethods[deName] = de
	return nil
}

//UnmarshalMethodsGet 添加解析器
func UnmarshalMethodsGet(deName string) (UnmarshalFiler, error) {
	mMutex.RLock()
	defer mMutex.RUnlock()
	if v, ok := unmarshalMethods[deName]; ok {
		return v.(UnmarshalFiler), nil
	}
	return nil, errors.New(deName + " does not exist")
}

//UnmarshalFile 配置文件格式解析
//列：config.UnmarshalFile("./fileName.json", &data)
//会根据文件名称中的.json判定需要什么类型解析器
func UnmarshalFile(fileName string, v interface{}) error {
	if fileName == "" {
		return errors.New("the file name cannot be empty")
	}

	fileNameStrs := strings.Split(strings.Trim(fileName, "."), ".")

	if len(fileNameStrs) <= 1 {
		return errors.New("incorrect file name")
	}

	unmarshalFiler, err := UnmarshalMethodsGet(fileNameStrs[len(fileNameStrs)-1])
	if err != nil {
		return err
	}

	return unmarshalFiler.UnmarshalFile(fileName, v)
}
