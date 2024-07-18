/*
Copyright © 2024 Tina Holly <info@tinaciousdesign.com>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/tinacious/random/utils"
)

var flagThousandsDelimiter string

var bignumCmd = &cobra.Command{
	Use:   "bignum",
	Short: "Generate a random large number with comma as the delimiter",
	Long:  "Do you ever feel like adding a giant random number to your sentences? You can get one with `random bignum`",

	Run: func(cmd *cobra.Command, args []string) {
		n := utils.RandomNumberBetweenRange(9999999, 1000000000000000000)
		formattedNumber := utils.FormatNumberWithDelimiter(n, flagThousandsDelimiter)

		fmt.Println(formattedNumber)
	},
}

func init() {
	rootCmd.AddCommand(bignumCmd)

	bignumCmd.Flags().StringVarP(&flagThousandsDelimiter, "delimiter", "d", ",", "Thousands delimiter")
}
