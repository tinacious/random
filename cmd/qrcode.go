package cmd

import (
	"bufio"
	"fmt"
	"image"
	"image/png"
	"log"
	"os"
	"strings"
	"time"

	"github.com/makiuchi-d/gozxing"
	"github.com/makiuchi-d/gozxing/qrcode"
	"github.com/spf13/cobra"
	"github.com/tinacious/random/utils"
	"github.com/xyproto/png2svg"
)

var qrcodeCmd = &cobra.Command{
	Use:   "qrcode",
	Short: "Generate a QR code with the provided text",
	Long:  "This can be a URL or any text, including useful formats like vCards: https://en.wikipedia.org/wiki/VCard",
	Run:   qrCodeCmd,
}

func qrCodeCmd(cmd *cobra.Command, args []string) {
	size, err := cmd.Flags().GetInt("size")
	if err != nil {
		log.Fatal("invalid flag: size")
	}

	shouldOpenFile, err := cmd.Flags().GetBool("open")
	if err != nil {
		log.Fatal("invalid flag: open")
	}

	fileFormat, err := cmd.Flags().GetString("format")
	if err != nil {
		log.Fatal("invalid flag: format")
	}
	if fileFormat != "png" && fileFormat != "svg" {
		log.Fatal("invalid flag: format - must be one of: svg, png")
	}

	scn := bufio.NewScanner(os.Stdin)

	fmt.Print("📝 Enter QR code data then press Ctrl+] and Enter:\n\n")
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
		log.Fatal("no data to encode")
	}

	if fileFormat == "png" {
		fmt.Println("✅ Got the input. Converting to PNG. Should be quick.")
	} else {
		fmt.Println("✅ Got the input. Converting to SVG. This might take a bit. If it's taking too long, try reducing the file size.")
	}

	enc := qrcode.NewQRCodeWriter()

	now := time.Now()

	// Filename using the file type `fileFormat` and date, e.g. qr-code_2025-07-20_21.53.51.svg
	fileName := fmt.Sprintf("qr-code_%s.%s", now.Format("2006-01-02_15.04.05"), fileFormat)

	img, err := enc.Encode(allLines, gozxing.BarcodeFormat_QR_CODE, size, size, nil)
	if err != nil {
		log.Fatal("failed to encode data")
	}

	file, err := os.Create(fileName)
	if err != nil {
		log.Fatal("failed to create QR code file")
	}
	defer file.Close()

	createImageFile(img, file, fileName, fileFormat)

	fmt.Printf("✅ Successfully created %s\n", fileName)

	if !shouldOpenFile {
		return
	}

	err = utils.OpenFile(fileName)
	if err != nil {
		log.Fatal("failed to open file")
	}
}

// createImageFile Checks the fileFormat and uses the respective library for generating an image
// the png2svg library does not return an error, assuming it will cause a fatal exception but i haven't seen it fail 🤷🏻‍♀️
func createImageFile(img image.Image, file *os.File, fileName string, fileFormat string) {
	if fileFormat == "png" {
		err := createPng(img, file)
		if err != nil {
			log.Fatal("failed to create QR code file as PNG")
		}
	} else {
		createSvg(img, fileName)
	}
}

func createPng(img image.Image, file *os.File) error {
	// *BitMatrix implements the image.Image interface,
	// so it is able to be passed to png.Encode directly.
	return png.Encode(file, img)
}

// uses png2svg to convert the image.Image interface into SVG
// ref: https://github.com/xyproto/png2svg/blob/48dc9089026fdc81c70bf46e560791d81dcafea7/cmd/png2svg/main.go#L146-L170
func createSvg(img image.Image, fileName string) {
	var (
		box          *png2svg.Box
		x, y         int
		expanded     bool
		lastx, lasty int
		lastLine     int // one message per line / y coordinate
		done         bool
	)
	var optimizeColors bool = true
	var verbose bool = false // todo: maybe make this configurable

	pi := png2svg.NewPixelImage(img, verbose)
	pi.SetColorOptimize(true)

	height := img.Bounds().Max.Y - img.Bounds().Min.Y

	percentage := 0
	lastPercentage := 0

	for !done {
		// Select the first uncovered pixel, searching from the given coordinate
		x, y = pi.FirstUncovered(lastx, lasty)

		if verbose && y != lastLine {
			lastPercentage = percentage
			percentage = int((float64(y) / float64(height)) * 100.0)
			png2svg.Erase(len(fmt.Sprintf("%d%%", lastPercentage)))
			fmt.Printf("%d%%", percentage)
			lastLine = y
		}

		// Create a box at that location
		box = pi.CreateBox(x, y)
		// Expand the box to the right and downwards, until it can not expand anymore
		expanded = pi.Expand(box)

		// Use the expanded box. Color pink if it is > 1x1, and colorPink is true
		pi.CoverBox(box, expanded && false, optimizeColors)

		// Check if we are done, searching from the current x,y
		done = pi.Done(x, y)
	}

	pi.WriteSVG(fileName)
}

func init() {
	qrcodeCmd.Flags().IntP("size", "s", 200, "How wide in pixels?")
	qrcodeCmd.Flags().Bool("open", true, "Open the file after generating?")
	qrcodeCmd.Flags().StringP("format", "f", "svg", "File format. Options include: svg, png")

	rootCmd.AddCommand(qrcodeCmd)
}
