package compiler

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

type CompilationOptions struct {
	IncludeToC bool
}

func CompileWithToC(dir, mainFile string) (string, error) {
	return compile(dir, mainFile, 2)
}

func OneTimeCompile(dir, mainFile string) (string, error) {
	return compile(dir, mainFile, 1)
}

func prepareDir(dir string) (string, error) {
	currDir, err := os.Getwd()
	if err != nil {
		return "", err
	}

	fmt.Printf("Current dir: %s\n", currDir)
	if err := os.Chdir(dir); err != nil {
		return "", err
	}

	fmt.Printf("switching to: %s", dir)

	return currDir, nil
}

func clean(dir string) error {
	err := os.Chdir(dir)
	if err != nil {
		return err
	}
	return nil
}

func compile(dir, mainFile string, rep int) (string, error) {
	currDir, err := prepareDir(dir)
	if err != nil {
		return "", fmt.Errorf("could not prepare for compilation: %v", err)
	}

	jn := strings.Split(mainFile, ".")[0]
	for k := 0; k < rep; k++ {
		cmd := exec.Command("pdflatex", mainFile)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		err = cmd.Run()
		if err != nil {
			return "", fmt.Errorf("LaTeX compilation failed with %s\n", err)
		}
	}

	if err := clean(currDir); err != nil {
		return "", fmt.Errorf("failed while cleaning: %v", err)
	}

	return jn + ".pdf", nil
}
