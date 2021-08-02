package types

import (
	"bytes"
	"fmt"
	"go-help/encode"
	"gopkg.in/yaml.v3"
	"io"
	"os"
)

type YAML struct {
}

var (
	Yaml encode.Decoder = YAML{}
)

func (YAML) UnmarshalFromReader(b io.Reader, v interface{}) error {
	return yaml.NewDecoder(b).Decode(v)
}
func (YAML) UnmarshalFromFile(fileName string, v interface{}) error {
	file, err := os.Open(fileName)
	if err != nil {
		return fmt.Errorf("yaml unmarshal file:%s", err)
	}
	defer file.Close()
	err = yaml.NewDecoder(file).Decode(v)
	if err != nil {
		return fmt.Errorf("yaml unmarshal file:%s", err)
	}
	return nil
}

func (YAML) UnmarshalStr(str string, v interface{}) error {
	return yaml.Unmarshal([]byte(str), v)
}

func (YAML) Unmarshal(b []byte, v interface{}) error {
	return yaml.Unmarshal(b, v)
}

func (YAML) MarshalStr(b interface{}) (result string, err error) {
	r, err := yaml.Marshal(b)
	if err != nil {
		return "", fmt.Errorf("yaml marshal str err:%s", err)
	}
	return string(r), nil
}

func (YAML) Marshal(b interface{}) (result []byte, err error) {
	r, err := yaml.Marshal(b)
	if err != nil {
		return []byte{}, fmt.Errorf("yaml marshal err:%s", err)
	}
	return r, nil
}

func (YAML) MarshalToReader(b interface{}) (result io.Reader, err error) {
	r, err := yaml.Marshal(b)
	if err != nil {
		return nil, fmt.Errorf("yaml marshal reader err:%s", err)
	}
	return bytes.NewReader(r), nil
}

func (YAML) MarshalToFile(fileName string, v interface{}) error {
	file, err := os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY, os.ModePerm)
	if err != nil {
		return fmt.Errorf("yaml marshal file err:%s", err)
	}
	defer file.Close()

	return yaml.NewEncoder(file).Encode(v)
}

func (YAML) MarshalToWriter(w io.Writer, v interface{}) error {
	return yaml.NewEncoder(w).Encode(v)
}
