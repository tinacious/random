/*
Copyright © 2025 Tina Holly <info@tinaciousdesign.com>
*/
package cmd

import (
	_ "embed"
	"fmt"

	"github.com/spf13/cobra"
	"github.com/tinacious/random/utils"
)

var nameCmd = &cobra.Command{
	Use:   "name",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		service := utils.NewNameService()
		// names, err = utils.RandomName()

		f := service.GetFirstNamesNonBinary()
		fmt.Printf("Names: %v (%d) last = '%s'", f, len(f), f[len(f)-1])
		// service.RandomName()
	},
}

func init() {
	rootCmd.AddCommand(nameCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// nameCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// nameCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
