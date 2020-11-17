package main

import (
	"fmt"
	"testing"
)

func TestType(t *testing.T) {
	f := func() string {
		return "Hello"
	}
	MyFunc(f).Response()
}

type MyFunc func() string

func (str MyFunc) Response() {
	fmt.Println(str())
}
