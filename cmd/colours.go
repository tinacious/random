/*
Copyright © 2024 Tina Holly <info@tinaciousdesign.com>
*/
package cmd

import (
	"bytes"
	b64 "encoding/base64"
	"fmt"
	"log"
	"strconv"

	"github.com/spf13/cobra"
	"github.com/tinacious/random/tmpl"
	"github.com/tinacious/random/utils"
)

var coloursCmd = &cobra.Command{
	Use:     "colours",
	Aliases: []string{"colors"},
	Short:   "Generates random colours as a base64 HTML string you can copy/paste into a browser",
	Long: `Generates 10 colours by default.
If you want a 20 colours, run the following command on macOS to copy the base64 URL to your clipboard:

	random colours 20 | pbcopy
`,
	Run: func(cmd *cobra.Command, args []string) {
		var err error
		var colourCount int = 10

		if len(args) == 1 {
			colourCount, err = strconv.Atoi(args[0])
			if err != nil {
				log.Fatalf("invalid number: %s", args[0])
			}
		}

		colours := utils.RandomColours(colourCount)
		page := tmpl.ColourPage{
			Colours: colours,
		}

		var doc bytes.Buffer
		err = page.Render(&doc)
		if err != nil {
			log.Fatal(err)
		}

		htmlString := doc.String()
		base64str := b64.StdEncoding.EncodeToString([]byte(htmlString))
		base64StringWithPrefix := fmt.Sprintf("data:text/html;base64,%s", base64str)

		fmt.Println(base64StringWithPrefix)
	},
}

func init() {
	rootCmd.AddCommand(coloursCmd)
}
