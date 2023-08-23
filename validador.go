package main

import (
	"log"
	"os"
	"os/exec"
	"path"
)

func main() {
	log.SetFlags(0)

	challengesDir := "desafios"
	challenges, err := os.ReadDir(challengesDir)
	if err != nil {
		log.Printf("falha ao ler desafios na pasta %s", challengesDir)
		os.Exit(1)
	}

	for _, challenge := range challenges {
		id := challenge.Name()
		dir := "./" + path.Join(challengesDir, id) // HACK: path.Join remove o ./
		cmd := exec.Command("go", "test", "./"+dir)
		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		err := cmd.Run()
		if err != nil {
			os.Exit(1)
		}
	}
}
