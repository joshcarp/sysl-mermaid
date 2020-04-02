package main

import (
	"flag"
	"fmt"
	"path/filepath"
	"strings"

	"github.com/spf13/afero"

	"github.com/joshcarp/sysl-mermaid/mermaid"
)

func main() {
	var output string
	flag.StringVar(&output, "o", "", "Output file of the svg")
	flag.Parse()
	filename := flag.Arg(0)
	if filename == "" {
		fmt.Println("Error, no filename specified")
	}
	if output == "" {
		output = strings.TrimSuffix(filename, filepath.Ext(filename)) + ".svg"
	}
	fmt.Println("iutfiuyifyf", output, filename)
	fs := afero.NewOsFs()
	file, err := afero.ReadFile(fs, filename)
	if err != nil {
		fmt.Println("Error reading input file")
	}
	result := mermaid.Execute(string(file))

	outfile, err := fs.Create(output)
	if err != nil {
		fmt.Println("Error creating output file")
	}
	outfile.Write([]byte(result))

}
