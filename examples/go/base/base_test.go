package base

import (
	"fmt"
	"reflect"
	"testing"
)

func TestSplit(t *testing.T) {
	tests := []struct {
		n    int
		by   int
		want []Idx
	}{
		{
			n: 1, by: 10,
			want: []Idx{
				{Start: 1, End: 1},
			},
		},
		{
			n: 5, by: 5,
			want: []Idx{
				{Start: 1, End: 5},
			},
		},
		{
			n: 10, by: 3,
			want: []Idx{
				{Start: 1, End: 3},
				{Start: 4, End: 6},
				{Start: 7, End: 9},
				{Start: 10, End: 10},
			},
		},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("%d by %d", tt.n, tt.by), func(t *testing.T) {
			got := split(tt.n, tt.by)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("want: %v, got: %v", tt.want, got)
			}
		})
	}
}
