// NÃO MODIFIQUE ESSE ARQUIVO

package main

import (
	"testing"
	. "treinamento_golang/test_helpers"
)

func TestShouldPrintHelloWorld(t *testing.T) {
	got := SpyStdout(t, func() {
		main()
	})

	want := "Hello, World!"
	AssertEqual(t, got, want)
}
