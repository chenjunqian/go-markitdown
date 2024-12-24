package converter

import (
	"path/filepath"
	"strings"

	md "github.com/JohannesKaufmann/html-to-markdown"
	"github.com/gogf/gf/v2/os/gfile"

	"github.com/gen2brain/go-fitz"
)

type PDFConverter struct {
}

var pdfConv *PDFConverter = new(PDFConverter)

func GetPDFConverter() *PDFConverter {
	return pdfConv
}

func (c *PDFConverter) Convert(localPath string, opt Options) (result DocumentConvertResult, err error) {

	var mdContent string
	mdContent, err = convertPDFToMDWithStyle(localPath)
	if opt.OutputPath != "" {
		err = gfile.PutContents(opt.OutputPath, mdContent)
		if err != nil {
			return
		}
	}
	result = DocumentConvertResult{
		Title:   strings.TrimSuffix(filepath.Base(localPath), filepath.Ext(localPath)),
		Content: mdContent,
	}
	return
}

func convertPDFToMDWithStyle(pdfPath string) (mdContent string, err error) {
	doc, err := fitz.New(pdfPath)
	if err != nil {
		return
	}
	defer doc.Close()

	numPages := doc.NumPage()

	for i := 0; i < numPages; i++ {
		var html string
		html, err = doc.HTML(i, true)
		if err != nil {
			return
		}

		converter := md.NewConverter("", true, nil)
		var text string
		text, err = converter.ConvertString(html)
		if err != nil {
			return
		}

		mdContent += text + "\n\n"
	}

	return
}