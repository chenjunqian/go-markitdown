package converter

import (
	"testing"
)

func TestHTMLConverter_Convert(t *testing.T) {
	type args struct {
		localPath string
		opt       Options
	}
	tests := []struct {
		name       string
		c          *HTMLConverter
		args       args
		wantErr    bool
	}{
		{
			name: "Test HTMLConverter Convert HTML to Markdown",
			args: args{
				localPath: "./testdata/test_html_content.html",
				opt: Options{
					FileExtention: "html",
					OutputPath: "./testdata/test_html_content.md",
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &HTMLConverter{}
			gotResult, err := c.Convert(tt.args.localPath, tt.args.opt)
			if (err != nil) != tt.wantErr {
				t.Errorf("HTMLConverter.Convert() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if gotResult.Content == "" {
				t.Errorf("HTMLConverter.Convert() gotResult.Content = %v, wantErr %v", gotResult.Content, tt.wantErr)
			}
		})
	}
}
