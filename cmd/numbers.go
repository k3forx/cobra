/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	countFlagNumbers int
	rangeFlagNumbers []string
)

// numbersCmd represents the numbers command
var numbersCmd = &cobra.Command{
	Use:   "numbers",
	Short: "Returns random numbers",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("numbers mode")
		fmt.Println("--count: ", countFlagNumbers)
		fmt.Println("--range: ", rangeFlagNumbers)
		// fmt.Println("--verbose: ", verbose)
	},
}

func init() {
	rootCmd.AddCommand(numbersCmd)

	// Declare the `--count` flag
	numbersCmd.Flags().IntVarP(
		&countFlagNumbers, "count", "c", 0,
		"A count of random numbers",
	)

	// Make this flag required
	numbersCmd.MarkFlagRequired("count")

	// Declare the `--range` flag
	numbersCmd.Flags().StringSliceVarP(
		&rangeFlagNumbers, "range", "r", []string{"1:100"},
		"Range of numbers. Optional",
	)
}
