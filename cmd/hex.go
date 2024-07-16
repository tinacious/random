/*
Copyright © 2024 Tina Holly <info@tinaciousdesign.com>
*/
package cmd

import (
	"fmt"
	"log"
	"strconv"

	"github.com/spf13/cobra"
	"github.com/tinacious/random/utils"
)

var hexCmd = &cobra.Command{
	Use:   "hex",
	Short: "A random hex value",
	Long:  "A random hex value that is 32 characters long by default. Pass an integer as an argument to change the length to n * 2, e.g. for 10 characters, do `random hex 5`",
	Run: func(cmd *cobra.Command, args []string) {
		var err error
		var length int = 16

		if len(args) > 0 {
			length, err = strconv.Atoi(args[0])
			if err != nil {
				log.Fatalf("invalid number: %s", args[0])
			}
		}

		hex, err := utils.RandomHex(length)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(hex)
	},
}

func init() {
	rootCmd.AddCommand(hexCmd)
}
