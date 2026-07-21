package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"

	"github.com/scusi/MultiChecksum"
)

func main() {
	myApp := app.NewWithID("github.com.scusi.MultiChecksum")
	myWindow := myApp.NewWindow("MultiChecksum")

	// Create UI elements
	fileEntry := widget.NewEntry()
	fileEntry.SetPlaceHolder("Enter file path or drag and drop a file here...")

	// Hash selection checkboxes
	md5Check := widget.NewCheck("MD5", func(checked bool) {})
	sha1Check := widget.NewCheck("SHA-1", func(checked bool) {})
	sha256Check := widget.NewCheck("SHA-256", func(checked bool) {})
	sha512Check := widget.NewCheck("SHA-512", func(checked bool) {})
	sha3256Check := widget.NewCheck("SHA3-256", func(checked bool) {})
	sha3512Check := widget.NewCheck("SHA3-512", func(checked bool) {})
	blake2sCheck := widget.NewCheck("Blake2s", func(checked bool) {})
	blake2bCheck := widget.NewCheck("Blake2b-256", func(checked bool) {})
	blake3Check := widget.NewCheck("Blake3-256", func(checked bool) {})
	blake2b512Check := widget.NewCheck("Blake2b-512", func(checked bool) {})

	// Select all checkbox by default
	md5Check.SetChecked(true)
	sha1Check.SetChecked(true)
	sha256Check.SetChecked(true)
	sha512Check.SetChecked(true)
	sha3256Check.SetChecked(true)
	sha3512Check.SetChecked(true)
	blake2sCheck.SetChecked(true)
	blake2bCheck.SetChecked(true)
	blake3Check.SetChecked(true)
	blake2b512Check.SetChecked(true)

	// Select all / Deselect all buttons
	selectAllBtn := widget.NewButton("Select All", func() {
		md5Check.SetChecked(true)
		sha1Check.SetChecked(true)
		sha256Check.SetChecked(true)
		sha512Check.SetChecked(true)
		sha3256Check.SetChecked(true)
		sha3512Check.SetChecked(true)
		blake2sCheck.SetChecked(true)
		blake2bCheck.SetChecked(true)
		blake3Check.SetChecked(true)
		blake2b512Check.SetChecked(true)
	})

	selectNoneBtn := widget.NewButton("Select None", func() {
		md5Check.SetChecked(false)
		sha1Check.SetChecked(false)
		sha256Check.SetChecked(false)
		sha512Check.SetChecked(false)
		sha3256Check.SetChecked(false)
		sha3512Check.SetChecked(false)
		blake2sCheck.SetChecked(false)
		blake2bCheck.SetChecked(false)
		blake3Check.SetChecked(false)
		blake2b512Check.SetChecked(false)
	})

	// Result display - use TextGrid for proper multi-line display
	resultText := widget.NewTextGrid()
	resultText.SetText("Checksums will appear here...")

	// File selection button
	browseBtn := widget.NewButton("Browse...", func() {
		dialog.ShowFileOpen(func(reader fyne.URIReadCloser, err error) {
			if err != nil {
				dialog.ShowError(err, myWindow)
				return
			}
			if reader == nil {
				return
			}
			defer reader.Close()

			// Convert URI path to absolute file path
			filePath := reader.URI().Path()

			// For file:// URIs, we need to handle them properly
			if strings.HasPrefix(filePath, "file://") {
				filePath = strings.TrimPrefix(filePath, "file://")
			}

			fileEntry.SetText(filePath)
		}, myWindow)
	})

	// Calculate button
	calculateBtn := widget.NewButton("Calculate Checksums", func() {
		filename := strings.TrimSpace(fileEntry.Text)

		if filename == "" {
			dialog.ShowInformation("Error", "Please select a file first", myWindow)
			return
		}

		// Check if file exists
		if _, err := os.Stat(filename); os.IsNotExist(err) {
			dialog.ShowError(fmt.Errorf("file does not exist: %s", filename), myWindow)
			return
		}

		// Show loading state
		resultText.SetText("Calculating checksums...")
		resultText.Refresh()

		// Load file
		data, err := os.ReadFile(filename)
		if err != nil {
			dialog.ShowError(err, myWindow)
			return
		}

		// Calculate checksums
		checksums, err := multichecksum.CalcChecksums(filename, data)
		if err != nil {
			dialog.ShowError(err, myWindow)
			return
		}

		// Build result string based on selected hashes
		var result strings.Builder
		result.WriteString(fmt.Sprintf("Checksums for: %s\n\n", filepath.Base(filename)))

		// Helper function to check if a hash should be included
		shouldInclude := func(hashName string) bool {
			switch hashName {
			case "MD5":
				return md5Check.Checked
			case "SHA1":
				return sha1Check.Checked
			case "SHA256":
				return sha256Check.Checked
			case "SHA512":
				return sha512Check.Checked
			case "SHA3-256":
				return sha3256Check.Checked
			case "SHA3-512":
				return sha3512Check.Checked
			case "Blake2s":
				return blake2sCheck.Checked
			case "Blake2b":
				return blake2bCheck.Checked
			case "Blake3-256":
				return blake3Check.Checked
			case "Blake2b-512":
				return blake2b512Check.Checked
			default:
				return true
			}
		}

		// Add selected checksums to result
		for _, h := range checksums.Hashes {
			if shouldInclude(h.HashName) {
				result.WriteString(fmt.Sprintf("%s: %x\n", h.HashName, h.Hash))
			}
		}

		// Update the result text grid
		resultText.SetText(result.String())
		resultText.Refresh()
	})

	// Copy to clipboard button
	copyBtn := widget.NewButton("Copy to Clipboard", func() {
		if resultText.Text() == "Checksums will appear here..." {
			dialog.ShowInformation("Info", "No checksums to copy", myWindow)
			return
		}
		myWindow.Clipboard().SetContent(resultText.Text())
		dialog.ShowInformation("Success", "Checksums copied to clipboard!", myWindow)
	})

	// Layout
	fileRow := container.NewHBox(
		fileEntry,
		browseBtn,
	)

	hashGrid := container.NewGridWithColumns(2,
		md5Check,
		sha1Check,
		sha256Check,
		sha512Check,
		sha3256Check,
		sha3512Check,
		blake2sCheck,
		blake2bCheck,
		blake3Check,
		blake2b512Check,
	)

	buttonRow := container.NewHBox(
		selectAllBtn,
		selectNoneBtn,
		layout.NewSpacer(),
		calculateBtn,
		copyBtn,
	)

	// Main content
	content := container.NewVBox(
		widget.NewLabel("MultiChecksum - Calculate multiple checksums at once"),
		widget.NewSeparator(),
		fileRow,
		hashGrid,
		buttonRow,
		widget.NewSeparator(),
		container.NewScroll(resultText),
	)

	myWindow.SetContent(content)
	myWindow.Resize(fyne.NewSize(800, 600))
	myWindow.ShowAndRun()
}
