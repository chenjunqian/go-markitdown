package converter

import (
	"testing"
)

func TestXlsxConverter_Convert(t *testing.T) {
	type args struct {
		localPath string
		opt       Options
	}
	tests := []struct {
		name       string
		c          *XlsxConverter
		args       args
		wantErr    bool
	}{
		{
			name: "Test XlsxConverter Convert Xlsx to Markdown",
			args: args{
				localPath: "./testdata/test_xlsx_content.xlsx",
				opt: Options{
					FileExtention: "xlsx",
					OutputPath: "./testdata/test_xlsx_content.md",
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &XlsxConverter{}
			gotResult, err := c.Convert(tt.args.localPath, tt.args.opt)
			if (err != nil) != tt.wantErr {
				t.Errorf("XlsxConverter.Convert() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if gotResult.Content == "" {
				t.Errorf("XlsxConverter.Convert() gotResult.Content is empty")
			}
		})
	}
}
