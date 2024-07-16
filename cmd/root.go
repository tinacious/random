/*
Copyright © 2024 Tina Holly <info@tinaciousdesign.com>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/cobra/doc"
)

var rootCmd = &cobra.Command{
	Use:   "random",
	Short: "⁉️ Generate random things!",
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}

	if os.Getenv("GENERATE_RANDOM_DOCS") == "1" {
		fmt.Println("📑 Generating documentation in ./docs")
		doc.GenMarkdownTree(rootCmd, "./docs")
	}
}

func init() {
	rootCmd.Root().CompletionOptions.DisableDefaultCmd = true
}
