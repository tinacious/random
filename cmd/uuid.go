/*
Copyright © 2024 Tina Holly <info@tinaciousdesign.com>
*/
package cmd

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/spf13/cobra"
	"github.com/spf13/cobra/doc"
)

var uuidCmd = &cobra.Command{
	Use:   "uuid",
	Short: "Generates an RFC 4122-compliant UUID",
	Run: func(cmd *cobra.Command, args []string) {
		id := uuid.NewString()
		fmt.Println(id)
	},
}

func init() {
	rootCmd.AddCommand(uuidCmd)
	doc.GenMarkdownTree(rootCmd, "./docs")
}
