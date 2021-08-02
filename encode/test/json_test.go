package test

import (
	"fmt"
	"go-help/encode/types"
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

func TestMarshal(t *testing.T) {
	result, err := types.Json.Marshal(struct {
		Test  string
		TestP string
		Testb string `json:"testb"`
	}{
		Test:  "你是谁",
		TestP: "我到底是谁",
	})
	fmt.Println(string(result), err)

	r1, err1 := types.Json.MarshalStr(struct {
		Test  string
		TestP string
		Testb string `json:"testb"`
	}{
		Test:  "你是谁",
		TestP: "我到底是谁",
	})
	fmt.Println(r1, err1)
	data := &JsonFile{}
	err2 := types.Json.MarshalToFile("./tsconfig2.json", data)
	fmt.Println(err2)
}

func TestUnmarshal(t *testing.T) {
	j := new(JsonFile)
	err := types.Json.UnmarshalFromFile("./tsconfig.json", j)
	fmt.Println(err, j)
}
