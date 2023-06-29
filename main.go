package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path"

	"github.com/ms-henglu/pal/formatter"
	"github.com/ms-henglu/pal/trace"
)

const version = "0.2.0"

var showHelp = flag.Bool("help", false, "Show help")
var showVersion = flag.Bool("version", false, "Show version")

func main() {
	input := ""
	output := ""
	mode := ""

	flag.StringVar(&input, "i", "", "Input terraform log file")
	flag.StringVar(&output, "o", "", "Output directory")
	flag.StringVar(&mode, "m", "markdown", "Output format, allowed values are `markdown` and `oav`")

	// backward compatibility, the first argument is the input file
	if len(os.Args) == 2 {
		if _, err := os.Stat(os.Args[1]); err == nil {
			input = os.Args[1]
			mode = "markdown"
		}
	}
	if input == "" {
		flag.Parse()
		if *showHelp {
			flag.Usage()
			os.Exit(0)
		}
		if *showVersion {
			fmt.Println(version)
			os.Exit(0)
		}
	}
	if input == "" {
		flag.Usage()
		log.Fatalf("[ERROR] input file is required")
	}

	if output == "" {
		output = path.Dir(input)
	}

	log.Printf("[INFO] input file: %s", input)
	log.Printf("[INFO] output directory: %s", output)
	log.Printf("[INFO] output format: %s", mode)

	traces, err := trace.RequestTracesFromFile(input)
	if err != nil {
		log.Fatalf("[ERROR] failed to parse request traces: %v", err)
	}

	switch mode {
	case "oav":
		format := formatter.OavTrafficFormatter{}
		files, err := os.ReadDir(output)
		if err != nil {
			log.Fatalf("[ERROR] failed to read output directory: %v", err)
		}
		index := len(files)
		for _, t := range traces {
			out := format.Format(t)
			index = index + 1
			outputPath := path.Join(output, fmt.Sprintf("trace-%d.json", index))
			if err := os.WriteFile(outputPath, []byte(out), 0644); err != nil {
				log.Fatalf("[ERROR] failed to write file: %v", err)
			}
			log.Printf("[INFO] output file: %s", outputPath)
		}
	case "markdown":
		content := `<!--
Tips:

1. Use Markdown preview mode to get a better reading experience.
2. If you want to select some of the request traces, in VSCode, use shortcut "Ctrl + K, 0" to fold all blocks.

-->

`
		format := formatter.MarkdownFormatter{}
		for _, t := range traces {
			content += format.Format(t)
		}
		if err := os.WriteFile(path.Join(output, "output.md"), []byte(content), 0644); err != nil {
			log.Fatalf("[ERROR] failed to write file: %v", err)
		}
		log.Printf("[INFO] output file: %s", path.Clean(path.Join(output, "output.md")))
	}

}
