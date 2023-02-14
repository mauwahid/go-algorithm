package binary

import "testing"

func TestSearch(t *testing.T) {
	type args struct {
		arr []int
		v   int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "found",
			args: args{
				arr: []int{1, 2, 3, 4, 5, 6, 7},
				v:   5,
			},
			want: 4,
		},
		{
			name: "found",
			args: args{
				arr: []int{1, 2, 3, 4, 54, 61, 74},
				v:   61,
			},
			want: 5,
		},

		{
			name: "notfound",
			args: args{
				arr: []int{1, 2, 3, 4, 54, 61, 74},
				v:   62,
			},
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Search(tt.args.arr, tt.args.v); got != tt.want {
				t.Errorf("Search() = %v, want %v", got, tt.want)
			}
		})
	}
}
