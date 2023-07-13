package manager

import "testing"

func TestDefaultValidateURI(t *testing.T) {
	type args struct {
		baseURI     string
		redirectURI string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name:    "success",
			args:    args{baseURI: "https://example.com/", redirectURI: "https://example.com/oauth/callback"},
			wantErr: false,
		}, {
			name:    "failure",
			args:    args{baseURI: "https://example.com/", redirectURI: "https://example-error.com/oauth/callback"},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := DefaultValidateURI(tt.args.baseURI, tt.args.redirectURI); (err != nil) != tt.wantErr {
				t.Errorf("DefaultValidateURI() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
