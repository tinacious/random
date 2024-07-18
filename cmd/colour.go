/*
Copyright © 2024 Tina Holly <info@tinaciousdesign.com>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/tinacious/random/utils"
)

var colourCmd = &cobra.Command{
	Use:     "colour",
	Aliases: []string{"color"},
	Short:   "Generate a single colour. For more, use `random colours n`",
	Run: func(cmd *cobra.Command, args []string) {
		c := utils.RandomColour()
		coolorsUrl := fmt.Sprintf("https://coolors.co/%s", c)

		fmt.Println(coolorsUrl)
	},
}

func init() {
	rootCmd.AddCommand(colourCmd)
}
