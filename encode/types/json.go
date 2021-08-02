package types

import (
	"bytes"
	"fmt"
	jsoniter "github.com/json-iterator/go"
	"go-help/encode"
	"io"
	"os"
)

type jsonType struct {
}

var (
	Json    encode.Decoder = jsonType{}
	jsonLib                = jsoniter.ConfigCompatibleWithStandardLibrary
)

func (jsonType) UnmarshalFromReader(b io.Reader, v interface{}) error {
	return jsonLib.NewDecoder(b).Decode(v)
}

func (jsonType) UnmarshalFromFile(fileName string, v interface{}) error {
	file, err := os.Open(fileName)
	if err != nil {
		return fmt.Errorf("json decode file err:%s", err)
	}

	defer file.Close()

	err = jsonLib.NewDecoder(file).Decode(v)
	if err != nil {
		return fmt.Errorf("json decode file err:%s", err)
	}
	return nil
}

func (jsonType) Unmarshal(b []byte, v interface{}) error {
	return jsonLib.Unmarshal(b, v)
}

func (jsonType) UnmarshalStr(str string, v interface{}) error {
	return jsonLib.UnmarshalFromString(str, v)
}

func (jsonType) Marshal(b interface{}) (result []byte, err error) {
	return jsonLib.Marshal(b)
}

func (jsonType) MarshalStr(b interface{}) (result string, err error) {
	r, err := jsonLib.MarshalToString(b)
	if err != nil {
		return "", fmt.Errorf("json encode str err:%s", err)
	}
	return r, nil
}

func (jsonType) MarshalToReader(b interface{}) (result io.Reader, err error) {
	d, err := jsonLib.Marshal(b)
	if err != nil {
		return nil, fmt.Errorf("json marshal reader err:%s", err)
	}
	return bytes.NewReader(d), nil
}

func (jsonType) MarshalToFile(fileName string, v interface{}) error {
	file, err := os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY, os.ModePerm)
	if err != nil {
		return fmt.Errorf("json marshal file err:%s", err)
	}
	defer file.Close()
	return jsonLib.NewEncoder(file).Encode(v)
}

func (jsonType) MarshalToWriter(w io.Writer, v interface{}) error {
	return jsonLib.NewEncoder(w).Encode(v)
}
