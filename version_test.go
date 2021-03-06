package main

import (
	"reflect"
	"testing"
)

func TestDedupEnv(t *testing.T) {
	tests := []struct {
		noCase bool
		in     []string
		want   []string
	}{
		{
			noCase: true,
			in:     []string{"k1=v1", "k2=v2", "K1=v3"},
			want:   []string{"K1=v3", "k2=v2"},
		},
		{
			noCase: false,
			in:     []string{"k1=v1", "K1=V2", "k1=v3"},
			want:   []string{"k1=v3", "K1=V2"},
		},
	}
	for _, tt := range tests {
		got := dedupEnv(tt.noCase, tt.in)
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("Dedup(%v, %q) = %q; want %q", tt.noCase, tt.in, got, tt.want)
		}
	}
}
