package server

import (
	"net/http"
	"testing"
)

func TestClientBasicHandler(t *testing.T) {
	type args struct {
		r *http.Request
	}
	tests := []struct {
		name    string
		args    args
		want    string
		want1   string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, err := ClientBasicHandler(tt.args.r)
			if (err != nil) != tt.wantErr {
				t.Errorf("ClientBasicHandler() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ClientBasicHandler() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("ClientBasicHandler() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestClientBodyHandler(t *testing.T) {
	type args struct {
		r *http.Request
	}
	tests := []struct {
		name    string
		args    args
		want    string
		want1   string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, err := ClientBodyHandler(tt.args.r)
			if (err != nil) != tt.wantErr {
				t.Errorf("ClientBodyHandler() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ClientBodyHandler() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("ClientBodyHandler() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
