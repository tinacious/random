/*
Copyright © 2024 Tina Holly <info@tinaciousdesign.com>
*/
package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
	"github.com/tinacious/random-cli/utils"
)

var portCmd = &cobra.Command{
	Use:   "port",
	Short: "Generates a random port number",
	Long:  "Generates a random port between 1024 and 65535 by default. You can provide a range, e.g. random port 2000-3000 to generate a port between 2000 and 3000.",
	Run:   runPortCmd,
}

func init() {
	rootCmd.AddCommand(portCmd)
}

func runPortCmd(cmd *cobra.Command, args []string) {
	var start int = 1024
	var end int = 66535

	if len(args) > 0 {
		s, e, err := utils.IntRangeFromString(args[0], 1024, 66535)
		if err != nil {
			log.Fatalf("invalid int range: %s", err.Error())
		}

		start = s
		end = e
	}

	randomPort := utils.RandomNumberBetweenRange(start, end)
	fmt.Println(randomPort)
	os.Exit(0)
}
