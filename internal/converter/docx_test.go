package converter

import (
	"testing"
)

func TestDocxConverter_Convert(t *testing.T) {
	type args struct {
		localPath string
		opt       Options
	}
	tests := []struct {
		name       string
		c          *DocxConverter
		args       args
		wantErr    bool
	}{
		{
			name: "Test DocxConverter Convert Docx to Markdown",
			args: args{
				localPath: "./testdata/test_docx_content.docx",
				opt: Options{
					FileExtention: "docx",
					OutputPath: "./testdata/test_docx_content.md",
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &DocxConverter{}
			gotResult, err := c.Convert(tt.args.localPath, tt.args.opt)
			if (err != nil) != tt.wantErr {
				t.Errorf("DocxConverter.Convert() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if gotResult.Content == "" {
				t.Errorf("DocxConverter.Convert() gotResult.Content is empty")
			}
		})
	}
}
