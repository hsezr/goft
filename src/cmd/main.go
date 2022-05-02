package main

import (
	. "mygin/src/classes"
	"mygin/src/goft"
	. "mygin/src/middlewares"
)

func main() {
	goft.
		Ignite().
		Attach(NewUserMid()).
		Mount("v1", NewIndexClass()).
		Mount("v2", NewUserClass()).
		Launch()
}
