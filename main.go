/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
package main

import (
	"os"
	"path/filepath"

	"github.com/Melbeck777/noq/cmd"
	"github.com/Melbeck777/noq/infrastructure/repository"
	"github.com/Melbeck777/noq/internal/application/usecase"
	"github.com/goark/gocli/exitcode"
	"github.com/goark/gocli/rwi"
)

func run(ui *rwi.RWI) exitcode.ExitCode {
	ui.Outputln("Hello world")
	return exitcode.Normal
}

func main() {
	home, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}
	configPath := filepath.Join(home, ".noq", "config.yml")
	configRepo := repository.NewConfigRepositoryImpl(configPath)
	configUseCase := usecase.NewConfigUseCase(configRepo)
	cmd.Execute(
		rwi.New(
			rwi.WithReader(os.Stdin),
			rwi.WithWriter(os.Stdout),
			rwi.WithErrorWriter(os.Stderr),
		),
		os.Args[1:],
		configUseCase,
	).Exit()
}
