package sequential

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
			name: "sequential",
			args: args{
				arr: []int{5, 3, 6, 1, 4},
				v:   4,
			},
			want: 4,
		},
		{
			name: "sequential",
			args: args{
				arr: []int{5, 3, 6, 1, 4},
				v:   6,
			},
			want: 2,
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
