package main

import (
    "github.com/andlabs/ui"
	"github.com/scusi/MultiChecksum"
    "io/ioutil"
    "fmt"
    "bytes"
    "bufio"
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
        button := ui.NewButton("MultiChecksum")
        greeting := ui.NewLabel("")
        box := ui.NewVerticalBox()
        box.Append(ui.NewLabel("Enter or Drag and Drop a file here:"), false)
        box.Append(name, false)
        box.Append(button, false)
        box.Append(greeting, false)
        window := ui.NewWindow("MultiChecksum", 600, 200, false)
        window.SetChild(box)
        button.OnClicked(func(*ui.Button) {
            filename := name.Text()
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
