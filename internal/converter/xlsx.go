package converter

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/gogf/gf/v2/os/gfile"
	"github.com/tealeg/xlsx"
)



type XlsxConverter struct {
}

var excelConv *XlsxConverter = new(XlsxConverter)

func GetXlsxConverter() *XlsxConverter {
	return excelConv
}

func (c *XlsxConverter) Convert(localPath string, opt Options) (result DocumentConvertResult, err error) {

	mdContents, err := readWriteSheet(localPath)
	if err != nil {
		return
	}

	if opt.OutputPath != "" {
		err = gfile.PutContents(opt.OutputPath, strings.Join(mdContents, ""))
		if err != nil {
			return
		}
	}

	result = DocumentConvertResult{
		Title:   strings.TrimSuffix(filepath.Base(localPath), filepath.Ext(localPath)),
		Content: strings.Join(mdContents, "\n"),
	}

	return
}

func readWriteSheet(inputFilePath string) (mdContents []string, err error) {
	xlFile, err := xlsx.OpenFile(inputFilePath)
	if err != nil {
		return
	}

	for _, sheet := range xlFile.Sheets {

		hyou := false

		for rowIdx, row := range sheet.Rows {

			if rowIdx == 0 {
				mdContents = append(mdContents, "# ")
			}

			text := ""
			for _, cell := range row.Cells {
				text += cell.Value
			}

			if len(text) == 0 {
				hyou = false
				mdContents = append(mdContents, "\n")
				mdContents = append(mdContents, "## ")
				continue
			}

			if len(row.Cells) >= 2 && len(row.Cells[0].Value) == 0 {
				mdContents = append(mdContents, "- ")
				idx := 1
				mdContents = append(mdContents, row.Cells[idx].Value)

			} else if len(row.Cells) >= 2 {

				for _, cell := range row.Cells {
					mdContents = append(mdContents, "|")
					mdContents = append(mdContents, cell.Value)
				}
				mdContents = append(mdContents, "|")

				if !hyou {
					mdContents = append(mdContents, "\n")
					mdContents = append(mdContents, strings.Repeat("| --- ", len(row.Cells)))
					mdContents = append(mdContents, "|")
					hyou = true
				}

			} else if strings.HasPrefix(row.Cells[0].Value, "http") {
				mdContents = append(mdContents, fmt.Sprintf("![%s](%s)", row.Cells[0].Value, row.Cells[0].Value))
				mdContents = append(mdContents, "\n")
			} else {
				mdContents = append(mdContents, text)
				mdContents = append(mdContents, "\n")
			}
			mdContents = append(mdContents, "\n")
		}
	}

	return
}