/*
Copyright © 2026 Tina Holly <info@tinaciousdesign.com>
*/
package cmd

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
	"github.com/tinacious/random/utils"
)

var colourContrastCmd = &cobra.Command{
	Use:   "contrast",
	Short: "Test a foreground and background colour for WCAG AA and AAA contrast requirements",
	Long: `Given 2 solid colours, one as foreground, the other as background,
it will test the contrast and generate a report based on WCAG AA and AAA compliance.
If you are working with semi-transparent colours, use the 'colours stacked' command to
resolve the painted colour.`,
	Run: func(cmd *cobra.Command, args []string) {
		foreground, err := cmd.Flags().GetString("fg")
		if err != nil || "" == foreground {
			log.Fatal("invalid flag: fg")
		}

		background, err := cmd.Flags().GetString("bg")
		if err != nil || "" == background {
			log.Fatal("invalid flag: bg")
		}

		ratio := utils.CalculateContrastRatio(foreground, background)
		report := fmt.Sprintf(`
===========================================
foreground %s vs. background %s

Ratio: %.2f:1

WCAG AA:
Standard text (smaller than 24px) must have a contrast ratio of 4.5:1 to pass WCAG AA requirements.
Large text (24px+) must have a contrast ratio of 3:1 to pass WCAG AA requirements.

WCAG AAA:
Standard text (smaller than 24px) must have a contrast ratio of 7:1 to pass WCAG AAA requirements.
Large text (24px+) must have a contrast ratio of 4.5:1 to pass WCAG AAA requirements.

Normal-sized text (23px or less):
WCAG AA: %s
WCAG AAA: %s

Large-sized text (24px or larger):
WCAG AA: %s
WCAG AAA: %s

More info: https://www.w3.org/WAI/WCAG21/Understanding/contrast-minimum.html
===========================================
`,
			foreground,
			background,
			ratio,
			utils.GetGradeForWCAGAA(ratio),
			utils.GetGradeForWCAGAAA(ratio),
			utils.GetGradeForWCAGAALarge(ratio),
			utils.GetGradeForWCAGAAALarge(ratio),
		)

		fmt.Println(report)

	},
}

func init() {
	colourCmd.AddCommand(colourContrastCmd)
	coloursCmd.AddCommand(colourContrastCmd)

	colourContrastCmd.Flags().StringP("fg", "f", "", "The foreground colour, e.g. #FFFFFF")
	colourContrastCmd.Flags().StringP("bg", "b", "", "The background colour, e.g. #000000")
}
