package main

import (
	"fmt"

	"github.com/tidwall/gjson"
)

func main() {
	data := `{"nice":"cool"}`
	v := gjson.Get(data, "nice")
	fmt.Println(v)
}
