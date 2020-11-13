package main

import (
	"github.com/sentiger/goft-gin/goft"
	. "github.com/sentiger/goft-gin/tests/controllers"
)

func main() {
	goft.Ignite().
		Mount(NewIndexController()).
		Launch()
}
