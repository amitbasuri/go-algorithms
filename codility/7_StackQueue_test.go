package codility

import "testing"

func TestFishSolution(t *testing.T) {
	type args struct {
		sizes      []int
		directions []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "1",
			args: args{
				sizes:      []int{4, 3, 2, 1, 5},
				directions: []int{0, 1, 0, 0, 0},
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FishSolution(tt.args.sizes, tt.args.directions); got != tt.want {
				t.Errorf("FishSolution() = %v, want %v", got, tt.want)
			}
		})
	}
}
