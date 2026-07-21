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
	Short: "Resolve a final colour from a stack of semi-transparent colours on a solid background",
	Long: `This command allows you to pass a list of colours with alpha channel to resolve
a final colour of the stacked colours, starting from the bottom.
This can be useful for calculating the actual colour of a background when there
are semi-transparent layers like overlays and text shadows.
You can use this resolved colour as your background when calculating contrast against WCAG requirements.`,
	Run: func(cmd *cobra.Command, args []string) {
		scn := bufio.NewScanner(os.Stdin)

		fmt.Print("🔴🟢🔵 Enter colours, 1 on each line, with alpha value between 0 and 1, then press Ctrl+] and Enter.\nStart from the bottom, e.g. if my background is white and I have a dark grey overlay, I could write:\n\n#FFFFFF 1\n#111111 0.4\n\n------------------------\n\n")
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
	},
}

func init() {
	coloursCmd.AddCommand(colourStackedCmd)
	colourCmd.AddCommand(colourStackedCmd)
}
