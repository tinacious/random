/*
Copyright © 2024 Tina Holly <info@tinaciousdesign.com>
*/
package cmd

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
	"github.com/tinacious/random/utils"
)

var passwordCmd = &cobra.Command{
	Use:   "password",
	Short: "Generate a random password",

	Run: func(cmd *cobra.Command, args []string) {
		length, err := cmd.Flags().GetInt("length")
		if err != nil {
			log.Fatal("invalid flag: length")
		}

		pw := utils.GeneratePassword(length)

		fmt.Println(pw)
	},
}

func init() {
	passwordCmd.Flags().IntP("length", "l", 37, "How many characters?")

	rootCmd.AddCommand(passwordCmd)
}
