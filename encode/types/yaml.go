package types

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"os"
)

type YAML struct {
}

func NewYaml() YAML {
	return YAML{}
}

func (YAML) DecodeFile(filePath string, v interface{}) error {
	file, err := os.Open(filePath)
	if err != nil {
		return fmt.Errorf("yaml unmarshal file:%s", err)
	}

	err = yaml.NewDecoder(file).Decode(v)
	if err != nil {
		return fmt.Errorf("yaml unmarshal file:%s", err)
	}
	return nil
}

func (YAML) DecodeStr(str string, v interface{}) error {
	return yaml.Unmarshal([]byte(str), v)
}

func (YAML) DecodeByte(b []byte, v interface{}) error {
	return yaml.Unmarshal(b, v)
}

func (YAML) EncodeStr(b interface{}) (result string, err error) {
	r, err := yaml.Marshal(b)
	if err != nil {
		return "", fmt.Errorf("json encode str err:%s", err)
	}
	return string(r), nil
}

func (YAML) EncodeBytes(b interface{}) (result []byte, err error) {
	r, err := yaml.Marshal(b)
	if err != nil {
		return []byte{}, fmt.Errorf("json encode bytes err:%s", err)
	}
	return r, nil
}
