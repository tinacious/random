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

var bignumCmd = &cobra.Command{
	Use:   "bignum",
	Short: "Generate a random large number with comma as the delimiter",
	Long:  "Do you ever feel like adding a giant random number to your sentences? You can get one with `random bignum`",

	Run: func(cmd *cobra.Command, args []string) {
		delimiter, err := cmd.Flags().GetString("delimiter")
		if err != nil {
			log.Fatal("invalid flag: delimiter")
		}

		n := utils.RandomNumberBetweenRange(9999999, 1000000000)
		formattedNumber := utils.FormatNumberWithDelimiter(n, delimiter)

		fmt.Println(formattedNumber)
	},
}

func init() {
	bignumCmd.Flags().StringP("delimiter", "d", ",", "Thousands delimiter")

	rootCmd.AddCommand(bignumCmd)
}
