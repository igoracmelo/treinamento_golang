// NÃO MODIFIQUE ESSE ARQUIVO

package main

import (
	"strings"
	"testing"
	. "treinamento_golang/test_helpers"
)

var want = strings.TrimSpace(`
1
2
3
4
5
6
7
8
9
10
11
12
13
14
15
`)

func TestShouldPrintFrom1To15(t *testing.T) {
	got := SpyStdout(t, func() {
		main()
	})

	AssertEqual(t, got, want)
}
