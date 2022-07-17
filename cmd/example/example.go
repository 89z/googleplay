package main

import (
	"github.com/89z/googleplay"
	"os"
)

func main() {
	var head googleplay.Header
	head.Open_Device("x86.bin")
	head.Open_Auth("auth.txt")
	head.Auth.Exchange()
	detail, err := head.Details("com.app.xt")
	if err != nil {
		panic(err)
	}
	text, err := detail.MarshalText()
	if err != nil {
		panic(err)
	}
	os.Stdout.Write(text)
}
