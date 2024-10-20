package goinct

import (
	"internal/abi" // its $GOROOT/src/internal
	"log"
	// "runtime/rtin"
)

type Type = abi.Type
type FuncType = abi.FuncType

func init() {
	if false {
		var x *FuncType
		log.Println(x.NumIn())
	}
}
