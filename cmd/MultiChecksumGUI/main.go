package main

import (
	"fmt"
	"net/url"
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
	myWindow.Resize(fyne.NewSize(800, 600))

	// Icon direkt im Code laden und für das Fenster setzen
	if iconBytes, err := os.ReadFile("Certainity_Icon.png"); err == nil {
		myWindow.SetIcon(fyne.NewStaticResource("Certainity_Icon.png", iconBytes))
	}
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

			filePath := reader.URI().Path()
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

		if _, err := os.Stat(filename); os.IsNotExist(err) {
			dialog.ShowError(fmt.Errorf("file does not exist: %s", filename), myWindow)
			return
		}

		resultText.SetText("Calculating checksums...")
		resultText.Refresh()

		data, err := os.ReadFile(filename)
		if err != nil {
			dialog.ShowError(err, myWindow)
			return
		}

		checksums, err := multichecksum.CalcChecksums(filename, data)
		if err != nil {
			dialog.ShowError(err, myWindow)
			return
		}

		var result strings.Builder
		result.WriteString(fmt.Sprintf("Checksums for: %s\n\n", filepath.Base(filename)))

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

		for _, h := range checksums.Hashes {
			if shouldInclude(h.HashName) {
				result.WriteString(fmt.Sprintf("%s: %x\n", h.HashName, h.Hash))
			}
		}

		resultText.SetText(result.String())
		resultText.Refresh()
	})

	// Copy to clipboard button
	copyBtn := widget.NewButton("Copy to Clipboard", func() {
		if resultText.Text() == "Checksums will_ appear here..." {
			dialog.ShowInformation("Info", "No checksums to copy", myWindow)
			return
		}
		myWindow.Clipboard().SetContent(resultText.Text())
		dialog.ShowInformation("Success", "Checksums copied to clipboard!", myWindow)
	})

	// Drag-and-Drop Handler für das Fenster einrichten
	myWindow.SetOnDropped(func(_ fyne.Position, uris []fyne.URI) {
		if len(uris) > 0 {
			filePath := uris[0].Path()
			if strings.HasPrefix(filePath, "file://") {
				filePath = strings.TrimPrefix(filePath, "file://")
			}
			fileEntry.SetText(filePath)
		}
	})

	// URL für das Repository parsen
	repoURL, _ := url.Parse("https://github.com/scusi/MultiChecksum")
	link := widget.NewHyperlink("github.com/scusi/MultiChecksum", repoURL)

	// Kopfzeile mit Beschreibung und Link nebeneinander
	headerRow := container.NewHBox(
		widget.NewLabel("MultiChecksum - Calculate multiple checksums at once"),
		layout.NewSpacer(),
		link,
	)

	// Layout for file selection
	fileRow := container.NewBorder(nil, nil, nil, browseBtn, fileEntry)

	// Kompaktes Grid mit 4 Spalten für die Checkboxen
	hashGrid := container.NewGridWithColumns(4,
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


	// Control panel (top part)
	controls := container.NewVBox(
		headerRow,
		widget.NewSeparator(),
		fileRow,
		hashGrid,
		buttonRow,
		widget.NewSeparator(),
	)

	// Result panel (bottom part)
	resultScroll := container.NewScroll(resultText)

	// Main layout
	content := container.NewBorder(controls, nil, nil, nil, resultScroll)

	myWindow.SetContent(content)
	myWindow.ShowAndRun()
}
