package main

import (
	. "mygin/src/classes"
	"mygin/src/goft"
)

func main() {
	goft.
	Ignite().
	Mount("v1", NewIndexClass(), NewUserClass()).
	Launch()
}
