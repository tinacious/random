package cmd

import (
	"bufio"
	"fmt"
	"image/png"
	"log"
	"os"
	"strings"
	"time"

	"github.com/makiuchi-d/gozxing"
	"github.com/makiuchi-d/gozxing/qrcode"
	"github.com/spf13/cobra"
	"github.com/tinacious/random/utils"
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

	enc := qrcode.NewQRCodeWriter()

	now := time.Now()
	fileName := fmt.Sprintf("qr-code_%s.png", now.Format("2006-01-02_15.04.05"))

	img, err := enc.Encode(allLines, gozxing.BarcodeFormat_QR_CODE, size, size, nil)
	if err != nil {
		log.Fatal("failed to encode data")
	}

	file, err := os.Create(fileName)
	if err != nil {
		log.Fatal("failed to create QR code file")
	}
	defer file.Close()

	// *BitMatrix implements the image.Image interface,
	// so it is able to be passed to png.Encode directly.
	err = png.Encode(file, img)
	if err != nil {
		log.Fatal("failed to create QR code file")
	}

	fmt.Printf("✅ Successfully created %s\n", fileName)

	if !shouldOpenFile {
		return
	}

	err = utils.OpenFile(fileName)
	if err != nil {
		log.Fatal("failed to open file")
	}
}

func init() {
	qrcodeCmd.Flags().IntP("size", "s", 1440, "How wide in pixels?")
	qrcodeCmd.Flags().Bool("open", false, "Open the file after generating?")

	rootCmd.AddCommand(qrcodeCmd)
}
