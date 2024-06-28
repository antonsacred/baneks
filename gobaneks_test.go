package main

import (
	"testing"
)

func TestRandomBAnek(t *testing.T) {
	tests := []struct {
		name    string
		wantErr bool
	}{
		{
			name:    "Should return a banek",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := RandomBAnek()
			if (err != nil) != tt.wantErr {
				t.Errorf("RandomBAnek() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if len(got) == 0 {
				t.Errorf("RandomBAnek() = %v, want something", got)
			}
		})
	}
}
