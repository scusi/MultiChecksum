package main

import (
	"bufio"
	"bytes"
	"fmt"
	"github.com/andlabs/ui"
	"github.com/scusi/MultiChecksum"
	"io/ioutil"
)

func loadFile(filename string) (data []byte, err error) {
	data, err = ioutil.ReadFile(filename)
	if err != nil {
		return data, err
	}
	return data, nil
}

func main() {
	err := ui.Main(func() {
		name := ui.NewEntry()
		button := ui.NewButton("generate checksums")
		radioButton := ui.NewRadioButtons()
		radioButton.Append("all checksums")
		radioButton.Append("selective checksums")

		md5Check := ui.NewCheckbox("MD5")
		sha1Check := ui.NewCheckbox("SHA-1")
		sha2Check := ui.NewCheckbox("SHA-2")
		sha5Check := ui.NewCheckbox("SHA-5")
		blake2b2Check := ui.NewCheckbox("Blake2 256")
		blake2b5Check := ui.NewCheckbox("Blake2 512")

		typeBox := ui.NewHorizontalBox()
		typeBox.Append(radioButton, false)
		typeBox.Append(md5Check, false)
		typeBox.Append(sha1Check, false)
		typeBox.Append(sha2Check, false)
		typeBox.Append(sha5Check, false)
		typeBox.Append(blake2b2Check, false)
		typeBox.Append(blake2b5Check, false)

		greeting := ui.NewLabel("")
		box := ui.NewVerticalBox()
		box.Append(typeBox, false)
		box.Append(ui.NewLabel("Enter or Drag and Drop a file here:"), false)
		box.Append(name, false)
		box.Append(button, false)
		box.Append(greeting, false)

		window := ui.NewWindow("MultiChecksum", 600, 200, false)
		window.SetChild(box)
		button.OnClicked(func(*ui.Button) {
			filename := name.Text()
			// TODO:
			// - split pathes (if possible)
			// - strip non runes
			// - strip 'file://' prefix
			// - strip whitespaces in the beginning and at the end of file
			data, err := loadFile(filename)
			if err != nil {
				greeting.SetText("Error: " + err.Error())
			} else {
				chksums := multichecksum.CalcChecksums(filename, data)
				var outbuf bytes.Buffer
				w := bufio.NewWriter(&outbuf)
				fmt.Fprintf(w, "Checksums for '%s'\n", filename)
				for typ, sum := range *chksums {
					if typ == "Filename" {
						continue
					} else {
						fmt.Fprintf(w, "%s", sum)
					}
				}
				fmt.Fprintln(w, "")
				fmt.Fprintln(w, "")
				w.Flush()

				greeting.SetText(outbuf.String())
			}
		})
		window.OnClosing(func(*ui.Window) bool {
			ui.Quit()
			return true
		})
		window.Show()
	})
	if err != nil {
		panic(err)
	}
}
