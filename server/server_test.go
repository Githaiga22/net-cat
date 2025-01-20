// server.go
package server

import (
	"net"
	"testing"
)

func TestIsNameUnique(t *testing.T) {
	type args struct {
		clientName string
	}

	// add name to
	var conn net.Conn

	clients[conn] = "munene"
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"Name is Absent", args{"fred"}, false},
		{"Name is Present", args{"munene"}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsNameUnique(tt.args.clientName); got != tt.want {
				t.Errorf("IsNameUnique() = %v, want %v", got, tt.want)
			}
		})
	}
}
