package converter

import (
	"testing"
)

func TestPDFConverter_Convert(t *testing.T) {
	type args struct {
		localPath string
		opt       Options
	}
	tests := []struct {
		name       string
		c          *PDFConverter
		args       args
		wantErr    bool
	}{
		{
			name: "Test PDFConverter Convert PDF to Markdown",
			args: args{
				localPath: "./testdata/test_pdf_code_block.pdf",
				opt: Options{
					FileExtention: "pdf",
					OutputPath: "./testdata/test_pdf_code_block.md",
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &PDFConverter{}
			gotResult, err := c.Convert(tt.args.localPath, tt.args.opt)
			if (err != nil) != tt.wantErr {
				t.Errorf("PDFConverter.Convert() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if gotResult.Content == "" {
				t.Errorf("PDFConverter.Convert() gotResult.Content = %v, wantErr %v", gotResult.Content, tt.wantErr)
			}
		})
	}
}
