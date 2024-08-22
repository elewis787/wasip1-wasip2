//go:build !wasip1

package main

import (
	"fmt"

	"github.com/elewis787/wasip1-wasip2/gen/go/wasip2/elewis787/example/bar"
	"github.com/elewis787/wasip1-wasip2/wasip2"
)

func init() {
	bar.Exports.Baz = wasip2.Baz
}

func main() {
	v := bar.Exports.Baz("")
	fmt.Println(v)
}
