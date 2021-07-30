package config

import (
	"fmt"
	"testing"
)

type JsonFile struct {
	CompilerOptions struct {
		Module    string `json:"module"`
		Target    string `json:"target"`
		SourceMap bool   `json:"sourceMap"`
	} `json:"compilerOptions"`

	Exclude []string `json:"exclude"`
}

func TestUnmarshalFile(t *testing.T) {
	data := new(JsonFile)
	err := UnmarshalFile("./tsconfig.json", data)
	fmt.Println(err, data)
}
