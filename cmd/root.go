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
package cmd

import (
	"github.com/Melbeck777/noq/internal/application/usecase"
	"github.com/goark/gocli/exitcode"
	"github.com/goark/gocli/rwi"
	"github.com/spf13/cobra"
)

func newRootCmd(ui *rwi.RWI, args []string, configUseCase usecase.ConfigUseCase) *cobra.Command {
	rootCmd := &cobra.Command{
		Use:   "noq",
		Short: "noq function",
		Long:  "noq function (detail)",
	}
	rootCmd.SilenceUsage = true
	rootCmd.SetArgs(args)
	rootCmd.SetIn(ui.Reader())
	rootCmd.SetOut(ui.ErrorWriter())
	rootCmd.SetErr(ui.ErrorWriter())
	rootCmd.AddCommand(
		newInitCmd(ui, configUseCase),
	)
	return rootCmd
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute(ui *rwi.RWI, args []string, configUseCase usecase.ConfigUseCase) exitcode.ExitCode {
	if err := newRootCmd(ui, args, configUseCase).Execute(); err != nil {
		return exitcode.Abnormal
	}
	return exitcode.Normal
}
