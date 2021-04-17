package json

import (
	"fmt"
	jsoniter "github.com/json-iterator/go"

	"testing"
)

type Person struct {
	FirstName  string `snake:"first_name" camel:"firstName"`
	LastName   string `snake:"last_name" camel:"lastName"`
	CurrentAge int    `snake:"current_age" camel:"currentAge"`
}

func TestJSON(t *testing.T) {
	print()
}

func print() {

	snakeCaseJSON := jsoniter.Config{TagKey: "snake"}.Froze()
	camelCaseJSON := jsoniter.Config{TagKey: "camel"}.Froze()

	m := make(map[string]interface{})
	m["persionId"] = "vvv"
	m["persionAge"] = 12

	result, _ := jsoniter.Marshal(m)
	fmt.Println(string(result))

	p := &Person{"Pepito", "Perez", 32}

	result, _ = snakeCaseJSON.Marshal(p)
	fmt.Println(string(result))

	result, _ = camelCaseJSON.Marshal(p)
	fmt.Println(string(result))
}
