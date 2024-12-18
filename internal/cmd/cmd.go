package cmd

import (
	"context"
	"errors"
	"fmt"
	"os"

	"github.com/chenjunqian/go-markitdown/markitdown"
	"github.com/gogf/gf/v2/os/gcmd"
)

var (
	Main = gcmd.Command{
		Name:  "main",
		Usage: "main",
		Brief: "Convert various file formats to markdown.",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			var (
				fileSource string
				outputPath string
			)
			fileSourceOpt := parser.GetOpt("f")
			outputPathOpt := parser.GetOpt("o")
			if fileSourceOpt.IsNil() {
				err = errors.New("Please input the path of the file to be converted with the -f parameter.")
			}

			if outputPathOpt.IsNil() {
				outputPath, err = os.Getwd()
				if err != nil {
					return err
				}
				outputPath = outputPath + "/markitdown_output.md"
			}

			fileSource = fileSourceOpt.String()
			result ,err := markitdown.Convert(fileSource, outputPath)

			title := result.Title
			content := result.Content

			if err == nil {
				fmt.Printf("Title: %s\nContent: %s\n", title, content)
			}

			return err
		},
	}
)