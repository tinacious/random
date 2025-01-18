/*
Copyright © 2025 Tina Holly <info@tinaciousdesign.com>
*/
package cmd

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
	"github.com/tinacious/random/utils"
)

var pinCmd = &cobra.Command{
	Use:   "pin",
	Short: "Generate a random PIN or small number",
	Long:  `Generate a small number suitable for combination locks and real estate lock boxes. Pipe to pbcopy to avoid outputting to standard out`,
	Run:   runPinCmd,
}

func init() {
	pinCmd.Flags().IntP("length", "l", 4, "How many digits?")

	rootCmd.AddCommand(pinCmd)
}

func runPinCmd(cmd *cobra.Command, args []string) {
	length, err := cmd.Flags().GetInt("length")
	if err != nil {
		log.Fatal("invalid flag: length")
	}

	var pin string
	for i := 1; i <= length; i++ {
		n := utils.RandomNumberBetweenRange(0, 9)
		pin = fmt.Sprintf("%s%d", pin, n)
	}

	fmt.Println(pin)
}
