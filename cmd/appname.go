/*
Copyright © 2024 Tina Holly <info@tinaciousdesign.com>
*/
package cmd

import (
	"fmt"
	"time"

	"github.com/spf13/cobra"
	"github.com/yelinaung/go-haikunator"
)

var appnameCmd = &cobra.Command{
	Use:   "appname",
	Short: "Generate a name for your new app",
	Run: func(cmd *cobra.Command, args []string) {
		h := haikunator.New(time.Now().UTC().UnixNano())
		fmt.Println(h.Haikunate())
	},
}

func init() {
	rootCmd.AddCommand(appnameCmd)
}
