package chat

import (
	"reflect"
	"testing"
)

func TestOutputIcon(t *testing.T) {
	tests := []struct {
		name    string
		want    []byte
		wantErr bool
	}{
		{"no file",nil,true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := OutputIcon()
			if (err != nil) != tt.wantErr {
				t.Errorf("OutputIcon() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("OutputIcon() = %v, want %v", got, tt.want)
			}
		})
	}
}
