package main

import (
	"encoding/json"
	"fmt"
)

// Dump proints propperly formated JSON
func Dump(obj interface{}) {
	result, _ := json.MarshalIndent(obj, "", "\t")
	fmt.Println(string(result))
}
