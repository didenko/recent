// Copyright © 2017 Vlad Didenko <business@didenko.com>
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

package cmd

import (
	"fmt"
	"log"
	"os"
	"path"
	"path/filepath"

	"github.com/didenko/recent/internal/pt"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

var (
	limit    int
	withdate bool
	withmill bool
	withtime bool
	root     = "."
	store    pt.PTslice
	tsFormat string
)

var rootCmd = &cobra.Command{
	Use:   "recent <path>",
	Short: "List a number of recently modified files ",
	Long: `The recent utility collects a number of most recently modified
files from the specified directory tree. Current directory is used by default.`,

	Args:   cobra.MaximumNArgs(1),
	PreRun: setup,
	Run:    collect,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().IntVarP(&limit, "limit", "n", 10, "a number of recent files to list")
	rootCmd.PersistentFlags().BoolVarP(&withdate, "date", "d", false, "include date in file timestamps output")
	rootCmd.PersistentFlags().BoolVarP(&withtime, "time", "t", false, "include time in file timestamps output")
	rootCmd.PersistentFlags().BoolVarP(&withmill, "mill", "m", false, "include milliseconts in file timestamps output - only if time included as well")
}

func setup(cmd *cobra.Command, args []string) {
	if len(args) > 0 {
		root = path.Clean(args[0])
	}

	tsFormat = pt.PickFormat(withdate, withtime, withmill)

	store = make(pt.PTslice, 0, limit+1)
}

func collect(cmd *cobra.Command, args []string) {
	err := filepath.Walk(root, consider)
	if err != nil {
		log.Fatal(errors.Wrap(err, "Scanning incomplete"))
	}

	for _, pt := range store {
		fmt.Println(pt.ToString(tsFormat))
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func consider(path string, info os.FileInfo, err error) error {
	if err != nil {
		return errors.Wrapf(err, "Error while scanning %q", path)
	}
	if !info.IsDir() {
		store = pt.Insert(store, pt.New(path, info.ModTime()))
		store = store[:min(limit, len(store))]
	}
	return nil
}
