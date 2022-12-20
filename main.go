package main

import (
	"fmt"

	"github.com/tidwall/gjson"
)

func main() {
	fmt.Println(PrettyNice())
}

func PrettyNice() string {
	data := `{"nice":"cool"}`
	v := gjson.Get(data, "nice")
	return v.String()
}
