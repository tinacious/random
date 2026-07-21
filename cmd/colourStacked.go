/*
Copyright © 2026 Tina Holly <info@tinaciousdesign.com>
*/
package cmd

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
	"github.com/tinacious/random/utils"
)

var colourStackedCmd = &cobra.Command{
	Use:   "stacked",
	Short: "Colour contrast calculations for Web Content Accessibility Guidelines",
	Long: `This command allows you to pass a list of colours with alpha channel to resolve
a final colour of the stacked colours, starting from the bottom.
This can be useful for calculating the actual colour of a background when there
are semi-transparent layers like overlays and text shadows.
You can use this resolved colour as your background when calculating contrast against WCAG requirements.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("colourStacked called")

		scn := bufio.NewScanner(os.Stdin)

		fmt.Print("🔴🟢🔵 Enter colours, 1 on each line, with alpha value between 0 and 1, then press Ctrl+] and Enter, e.g.:\n\n#FFFFFF 1\n#111111 0.4\n\n")
		var lines []string
		for scn.Scan() {
			line := scn.Text()
			if len(line) == 1 {
				// Group Separator (GS ^]): ctrl-]
				if line[0] == '\x1D' {
					break
				}
			}
			lines = append(lines, line)
		}

		allLines := strings.Join(lines, "\n")
		allLines = strings.TrimSpace(allLines)

		if allLines == "" {
			log.Fatal("no colour data")
		}

		list := strings.Split(allLines, "\n")

		var layers = make([]utils.ColorLayer, len(list))

		for i, line := range list {
			parts := strings.Split(line, " ")

			if len(parts) != 2 {
				log.Fatalf("invalid colour data: %s", line)
			}

			n, err := strconv.ParseFloat(parts[1], 8)
			if err != nil {
				log.Fatalf("not a number: %s", parts[1])
			}

			layer := utils.ColorLayer{
				Hex:   parts[0],
				Alpha: n,
			}

			layers[i] = layer
		}

		resolved, err := utils.GetStackedResolvedColour(layers)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("✅ The resolved colour is: %s\n", resolved)

		// todo: take a list of colours, e.g. background, layer 1, layer 2, text-shadow, foreground
	},
}

func init() {
	coloursCmd.AddCommand(colourStackedCmd)
	colourCmd.AddCommand(colourStackedCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// colourContrastCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// colourContrastCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
