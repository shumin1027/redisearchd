package json

import (
	"fmt"
	"github.com/RediSearch/redisearch-go/redisearch"
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

	jsoniter.RegisterExtension(NewJSONStyleExtension(true, KebabCase))

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


var data  = `
{
    "Raw": "*",
    "Paging": {
        "Offset": 0,
        "Num": 10
    },
    "Flags": 0,
    "Slop": 0,
    "Filters": null,
    "InKeys": null,
    "InFields": null,
    "ReturnFields": null,
    "Language": "",
    "Expander": "",
    "Scorer": "",
    "Payload": null,
    "SortBy": null,
    "HighlightOpts": null,
    "SummarizeOpts": null
}
`
func TestUn(t *testing.T) {
	var query = new(redisearch.Query)
	bytes := []byte(data)
	jsoniter.Unmarshal(bytes,query)
	println(query.Raw)
}
