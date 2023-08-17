package errorgen

import (
	_ "embed"
	"gopkg.in/yaml.v3"
	"testing"
)

//go:embed errx.yaml
var errxConfRaw []byte

func TestGenCode(t *testing.T) {
	type errConfStruct struct {
		Errx map[int32]map[string]string `yaml:"errx"`
	}
	var errxConfStruct errConfStruct
	if err := yaml.Unmarshal(errxConfRaw, &errxConfStruct); err != nil {
		panic(err)
	}
	err := GenFile(errxConfStruct.Errx, "errorgen", "errx.go", true)
	if err != nil {
		t.Fatal(err)
	}
}
