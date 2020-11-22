package others

import "testing"

func TestSolutionBulb(t *testing.T) {
	type args struct {
		A []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "1",
			args: args{A: []int{2, 1, 3, 5, 4}},
			want: 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SolutionBulb(tt.args.A); got != tt.want {
				t.Errorf("SolutionBulb() = %v, want %v", got, tt.want)
			}
		})
	}
}
