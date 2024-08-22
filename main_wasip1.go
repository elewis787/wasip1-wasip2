//go:build wasip1

package main

import (
	"github.com/elewis787/wasip1-wasip2/gen/go/wasip1/elewis787/example/bar"
	"github.com/elewis787/wasip1-wasip2/wasip1"
)

func init() {
	bar.Exports.Baz = wasip1.Baz
}

func main() {
}
