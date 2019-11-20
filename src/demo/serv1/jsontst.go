package main

import (
	"buster/utils"
	"bytes"
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	//utils.GlobalObject
	buf, err := json.Marshal(utils.GlobalObject)
	if err != nil {
		fmt.Println("json marshal error:", err)
	}
	var out bytes.Buffer
	json.Indent(&out, buf, "", "	")
	file, err := os.OpenFile("test2.json", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("open file error:", err)
	}
	_, err = out.WriteTo(file)
	if err != nil {
		fmt.Println("buffer write error:", err)
	}
}
