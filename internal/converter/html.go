package converter

import (
	"path/filepath"
	"strings"

	md "github.com/JohannesKaufmann/html-to-markdown"
	"github.com/gogf/gf/v2/os/gfile"
)

type HTMLConverter struct {
}

var htmlConv *HTMLConverter

func GetHTMLConverter() *HTMLConverter {
	return htmlConv
}

func (c *HTMLConverter) Convert(localPath string, opt Options) (result DocumentConvertResult, err error) {

	htmlContent := gfile.GetContents(localPath)
	converter := md.NewConverter("", true, nil)
	text, err := converter.ConvertString(htmlContent)
	if err != nil {
		return
	}

	if opt.OutputPath != "" {
		err = gfile.PutContents(opt.OutputPath, text)
		if err != nil {
			return
		}
	}

	result = DocumentConvertResult{
		Title:   strings.TrimSuffix(filepath.Base(localPath), filepath.Ext(localPath)),
		Content: text,
	}
	return
}
