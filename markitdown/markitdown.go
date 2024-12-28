package markitdown

import (
	"github.com/chenjunqian/go-markitdown/internal/converter"
	"github.com/gogf/gf/v2/text/gstr"
)

type ConverResult struct {
	Title   string
	Content string
	Path    string
}

var converters map[string]converter.Converter

func init() {
	converters = map[string]converter.Converter{
		string(converter.PDF): converter.GetPDFConverter(),
		string(converter.HTML): converter.GetHTMLConverter(),
		string(converter.DOCX): converter.GetDocxConverter(),
	}
}

func Convert(source, outputPath string) (result ConverResult, err error) {

	if gstr.HasPrefix(source, "http://") || gstr.HasPrefix(source, "https://") || gstr.HasPrefix(source, "file://") {

	} else {
		var docResult converter.DocumentConvertResult
		docResult, err = converLocal(source, outputPath)
		if err != nil {
			return
		}

		result = ConverResult{
			Title:   docResult.Title,
			Content: docResult.Content,
			Path:    outputPath,
		}
	}

	return
}

func converLocal(sourcePath, outputPath string) (result converter.DocumentConvertResult, err error) {

	contentType := contentTypeDetector(sourcePath)
	opt := converter.Options{
		FileExtention: string(contentType),
		Url:           sourcePath,
		OutputPath:    outputPath,
	}

	if converter, ok := converters[string(contentType)]; ok {
		result, err = converter.Convert(sourcePath, opt)
		if err != nil {
			return
		}
	}

	return
}

func contentTypeDetector(sourcePath string) (contentType converter.SourceContentType) {

	if gstr.HasSuffix(sourcePath, ".pdf") {
		contentType = converter.PDF
	}

	return
}
