package encoding

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
