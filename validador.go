package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path"
	"strings"

	"github.com/sergi/go-diff/diffmatchpatch"
)

func main() {
	challengesDir := "./desafios"
	challenges, err := os.ReadDir(challengesDir)
	if err != nil {
		log.Printf("falha ao ler desafios na pasta %s", challengesDir)
		os.Exit(1)
	}

	for _, challenge := range challenges {
		challengeID := challenge.Name()
		challengeDir := path.Join(challengesDir, challengeID)
		testsFolder := path.Join(challengeDir, "testes")

		testFiles, err := os.ReadDir(testsFolder)
		if err != nil {
			log.Printf("erro: não foi possível listar os arquivos em %s", testsFolder)
			return
		}

		prev := ""
		fmt.Printf("%s\n", challengeID)
		for _, f := range testFiles {
			testID := f.Name()
			testID = strings.TrimSuffix(testID, path.Ext(testID))
			if testID == prev {
				continue
			}
			prev = testID

			passed, err := runTestCase(challengeDir, testID)
			if !passed || err != nil {
				os.Exit(1)
			}
		}
	}
}

func sanitizeText(s string) string {
	s = strings.TrimSpace(s)
	s = strings.ReplaceAll(s, "\r", "")	
	return s
}

func runTestCase(challengeDir string, testID string) (bool, error) {
	inPath := path.Join(challengeDir, "testes", testID+".in")
	outPath := path.Join(challengeDir, "testes", testID+".out")

	outWriter := &bytes.Buffer{}
	cmd := exec.Command("go", "run", ".")
	cmd.Dir = challengeDir
	cmd.Stdin, _ = os.Open(inPath)
	cmd.Stdout = outWriter
	cmd.Stderr = outWriter

	err := cmd.Run()
	if err != nil {
		log.Printf("erro: não foi possível executar a sua solução do exercício %s: %v", challengeDir, err)
		log.Print(outWriter.String())
		return false, err
	}

	got := sanitizeText(outWriter.String())
	b, err := os.ReadFile(outPath)
	if err != nil {
		log.Printf("erro: não foi possível ler o gabarito do teste %s (arquivo %s): %v", testID, outPath, err)
		return false, err
	}

	want := sanitizeText(string(b))

	const green = "\033[0;32m"
	const red = "\033[0;31m"
	const noColor = "\033[0m"

	if got == want {
		fmt.Printf("  %s[V]%s teste %s PASSOU\n", green, noColor, testID)
		return true, nil
	} else {
		fmt.Printf("  %s[X]%s teste %s FALHOU.\n", red, noColor, testID)
		fmt.Printf("Esperado:\n'%s'\n\n", want)
		fmt.Printf("Obtido:\n'%s'\n\n", got)
		fmt.Println("Diferença:")
		differ := diffmatchpatch.New()
		fmt.Println(differ.DiffPrettyText(differ.DiffMain(got, want, true)))
		fmt.Println()
		return false, nil
	}

}
