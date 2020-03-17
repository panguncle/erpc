package main

import (
	"context"
	"fmt"
	"reflect"
)

func main() {
	var x context.Context
	var a = reflect.TypeOf(x)

	fmt.Printf("%+#v\n", a)
	fmt.Printf("%+#v\n", reflect.TypeOf((*context.Context)(nil)).Elem())
	fmt.Printf("%+#v\n", reflect.TypeOf(new(context.Context)))
}
