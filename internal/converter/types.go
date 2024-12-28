package converter

type Converter interface {
	Convert(localPath string, opt Options) (result DocumentConvertResult, err error)
}

type Options struct {
	FileExtention string
	Url           string
	OutputPath    string
}

type DocumentConvertResult struct {
	Title   string
	Content string
}

type SourceContentType string

const (
	PDF  SourceContentType = "pdf"
	HTML SourceContentType = "html"
	DOCX SourceContentType = "docx"
)
