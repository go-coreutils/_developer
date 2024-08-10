// Copyright (c) 2024  The Go-CoreLibs Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/urfave/cli/v2"
)

var (
	BuildVersion = "0.1.0"
	BuildRelease = "0000000000"
)

var (
	gApp = &cli.App{
		Name:            "tmplcmdname",
		Version:         BuildVersion + " (" + BuildRelease + ")",
		Usage:           "tmplsummary",
		UsageText:       "tmplcmdname [options] [arguments]",
		Description:     "tmpldescription",
		HideHelpCommand: true,
		Flags:           []cli.Flag{},
		Action: func(ctx *cli.Context) error {
			return action(ctx, fn)
		},
	}
)

func init() {
	cli.HelpFlag = &cli.BoolFlag{
		Category: "General",
		Name:     "help",
		Aliases:  []string{"h", "usage"},
	}
	cli.VersionFlag = &cli.BoolFlag{
		Category: "General",
		Name:     "version",
		Aliases:  []string{"V"},
	}
}

func main() {
	if err = app.Run(os.Args); err != nil {
		fatal(os.Stderr, "error: %v\n", err)
	}
}

func action(ctx *cli.Context, fn func(string) string) (err error) {
	fmt.Fprintf(os.Stdout, "Hello World\n")
	return
}

func stderr(format string, argv ...interface{}) {
	_, _ = fmt.Fprintf(os.Stderr, format, argv...)
}

func stdout(format string, argv ...interface{}) {
	_, _ = fmt.Fprintf(os.Stdout, format, argv...)
}

func fatal(format string, argv ...interface{}) {
	stderr(format, argv...)
	os.Exit(1)
}
