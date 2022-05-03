package main

import (
	. "mygin/classes"
	"mygin/goft"
	. "mygin/middlewares"
)

func main() {
	goft.
		Ignite().
		Beans(goft.NewGormAdapter()).
		Attach(NewUserMid()).
		Mount("v1", NewIndexClass()).
		Mount("v2", NewUserClass()).
		Launch()
}
