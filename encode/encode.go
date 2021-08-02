package encode

import (
	"io"
)

type Decoder interface {
	//
	UnmarshalFromReader(b io.Reader, v interface{}) error
	//根据文件解析
	UnmarshalFromFile(fileName string, v interface{}) error
	//根据字符串解析
	UnmarshalStr(str string, v interface{}) error
	//根据byte解析
	Unmarshal(b []byte, v interface{}) error
	//解析成字符串
	MarshalStr(b interface{}) (result string, err error)
	//解析成byte
	Marshal(b interface{}) (result []byte, err error)
	//解析io.Reader
	MarshalToReader(b interface{}) (result io.Reader, err error)
	//写入到文件中
	MarshalToFile(fileName string, v interface{}) error
	//写入到w中
	MarshalToWriter(w io.Writer, v interface{}) error
}
