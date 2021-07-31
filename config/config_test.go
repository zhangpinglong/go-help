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

type YamlFile struct {
	Server struct {
		HostPort   string `yaml:"hostPort"`
		BkPort     string `yaml:"bkPort"`
		WebPort    string `yaml:"webPort"`
		WebsslPort string `yaml:"websslPort"`
		AgentPort  string `yaml:"agentPort"`
		CenterPort string `yaml:"centerPort"`
		PnsslPort  string `yaml:"pnsslPort"`
		PnPort     string `yaml:"pnPort"`
		Debug      string `yaml:"debug"`
		AccessLog  string `yaml:"accessLog"`
		Doc        bool   `yaml:"doc"`
	} `yaml:"server"`
}

func TestUnmarshalFile(t *testing.T) {
	data := new(JsonFile)
	err := DecodeFile("./tsconfig.json", data, "json")
	fmt.Println(err, data)

	yamlData := new(YamlFile)
	err = DecodeFile("./testyaml.yaml", yamlData, "yaml")
	fmt.Println(err, yamlData)
}

func BenchmarkStdJsonUnmarshalFile(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		data := new(JsonFile)
		_ = DecodeFile("./tsconfig.json", data, "json")
	}
}
