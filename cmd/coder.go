package main

import (
	"flag"
	"fmt"

	"github.com/fullstackly/coder/engine"
)

var (
	text2Encode = flag.String("encode", "", "encodes given string")
	text2Decode = flag.String("decode", "", "decodes given string")
)

func main() {
	flag.Parse()
	process()
}

func process() {
	if len(*text2Encode) != 0 {
		fmt.Println(engine.Encode(*text2Encode))
	} else if len(*text2Decode) != 0 {
		fmt.Println(engine.Decode(*text2Decode))
	} else {
		fmt.Println("Enter command and text. Enter -h for help")
	}
}
