package codility

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
		{
			name: "1",
			args: args{A: []int{2, 1, 3, 5, 4}},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SolutionBulb(tt.args.A); got != tt.want {
				t.Errorf("SolutionBulb() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSolutionNumeratorDenominator(t *testing.T) {
	type args struct {
		nums []int
		dens []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "1",
			args: args{
				nums: []int{1, 2, 3, 4, 90, 1000},
				dens: []int{2, 3, 6, 8, 4, 2000},
			},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SolutionNumeratorDenominator(tt.args.nums, tt.args.dens); got != tt.want {
				t.Errorf("SolutionNumeratorDenominator() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkSolutionNumeratorDenominator(b *testing.B) {
	type args struct {
		nums []int
		dens []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "1",
			args: args{
				nums: []int{100000000000000000, 223423423, 333333333333, 4000000000000, 9075543434232332, 100000000000000000, 100000000000000000},
				dens: []int{200000000000000000, 323343244, 666666666666, 8000000000000, 4232342342342323, 200000000000000000, 1100000000000000000},
			},
			want: 4,
		},
	}
	for _, tt := range tests {
		_ = SolutionNumeratorDenominator(tt.args.nums, tt.args.dens)
	}
}

func TestSolutionPositiveNegative(t *testing.T) {
	type args struct {
		A []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "1",
			args: args{A: []int{1, 2, 3, -4, 3, -3, 0, 7}},
			want: 3,
		},
		{
			name: "2",
			args: args{A: []int{11, 2, 3, -4, 3, 4, 3, 0, 17}},
			want: 4,
		},
		{
			name: "3",
			args: args{A: []int{1, -4, -3, -4, 2, 4, 3, -4, 3, -3, 7, 7}},
			want: 4,
		},
		{
			name: "4",
			args: args{A: []int{1, 2, 3}},
			want: 0,
		},
		{
			name: "5",
			args: args{A: []int{1, 4, -2, 3, -3, 4}},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SolutionPositiveNegative(tt.args.A); got != tt.want {
				t.Errorf("SolutionPositiveNegative() = %v, want %v", got, tt.want)
			}
		})
	}
}
