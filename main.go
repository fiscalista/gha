package main

import (
	"fmt"
	"time"

	"github.com/tidwall/gjson"
)

func main() {
	fmt.Println(PrettyNice())
}

func PrettyNice() string {
	data := `{"nice":"cool"}`
	v := gjson.Get(data, "nice")
	time.Sleep(time.Duration(1 * time.Nanosecond))
	return v.String()
}
