package converter

import (
	"path/filepath"
	"strings"

	"github.com/gogf/gf/v2/os/gfile"
	"github.com/mandolyte/mdtopdf"
)

type PDFConverter struct {
}

var pdfConv *PDFConverter = new(PDFConverter)

func GetPDFConverter() *PDFConverter {
	return pdfConv
}

func (c *PDFConverter) Convert(localPath string, opt Options) (result DocumentConvertResult, err error) {
	var render *mdtopdf.PdfRenderer
	var opts []mdtopdf.RenderOption
	var pdfContent = gfile.GetContents(localPath)
	render = mdtopdf.NewPdfRenderer("", "", opt.OutputPath, "", opts, mdtopdf.LIGHT)
	err = render.Process([]byte(pdfContent))
	if err != nil {
		return
	}

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
