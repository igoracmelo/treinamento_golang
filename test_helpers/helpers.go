package test_helpers

import (
	"fmt"
	"io"
	"os"
	"strings"
	"testing"

	"github.com/sergi/go-diff/diffmatchpatch"
)

func SpyStdout(t *testing.T, fn func()) string {
	t.Helper()

	r, w, err := os.Pipe()
	AssertNoError(t, err)

	stdout := os.Stdout
	os.Stdout = w
	defer func() {
		os.Stdout = stdout
	}()

	fn()
	w.Close()

	b, err := io.ReadAll(r)
	AssertNoError(t, err)
	return strings.TrimSpace(string(b))
}

func AssertNoError(t *testing.T, err error) {
	t.Helper()
	if err != nil {
		t.Fatal(err)
	}
}

func AssertEqual[T comparable](t *testing.T, got T, want T) {
	t.Helper()

	if got == want {
		return
	}

	sGot := fmt.Sprintf("%v", got)
	sWant := fmt.Sprintf("%v", want)

	differ := diffmatchpatch.New()
	diff := differ.DiffPrettyText(differ.DiffMain(sGot, sWant, true))

	t.Logf("\nEsperado:\n\"%s\"\n\nObtido:\n\"%s\"\n\nDiferença:\n\"%s\"\n\n", sWant, sGot, diff)
	t.FailNow()
}
