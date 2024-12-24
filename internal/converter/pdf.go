package converter

import (
	"os"
	"path/filepath"
	"strings"

	md "github.com/JohannesKaufmann/html-to-markdown"

	"github.com/gen2brain/go-fitz"
	"github.com/gogf/gf/v2/os/gfile"
)

type PDFConverter struct {
}

var pdfConv *PDFConverter = new(PDFConverter)

func GetPDFConverter() *PDFConverter {
	return pdfConv
}

func (c *PDFConverter) Convert(localPath string, opt Options) (result DocumentConvertResult, err error) {

	err = convertPDFToMDWithStyle(localPath, opt.OutputPath)
	var fileName string
	var outputFileName string
	if opt.OutputPath == "" {
		outputFileName = filepath.Base(localPath)
	} else {
		outputFileName = filepath.Base(opt.OutputPath)
	}
	fileName = strings.TrimSuffix(outputFileName, filepath.Ext(outputFileName))
	result = DocumentConvertResult{
		Title:   fileName,
		Content: gfile.GetContents(opt.OutputPath),
	}
	return
}

func convertPDFToMDWithStyle(pdfPath, mdPath string) error {
	doc, err := fitz.New(pdfPath)
	if err != nil {
		return err
	}
	defer doc.Close()

	numPages := doc.NumPage()
	var mdContent string

	for i := 0; i < numPages; i++ {
		html, err := doc.HTML(i, true)
		if err != nil {
			return err
		}

		converter := md.NewConverter("", true, nil)
		text, err := converter.ConvertString(html)
		if err != nil {
			return err
		}

		mdContent += text + "\n\n"
	}

	err = os.WriteFile(mdPath, []byte(mdContent), 0644)
	if err != nil {
		return err
	}

	return nil
}