package coursera

import (
	"testing"
)

func Test_multiply(t *testing.T) {
	type args struct {
		a string
		b string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "1",
			args: args{
				a: "10",
				b: "10",
			},
			want: "100",
		},
		{
			name: "2",
			args: args{
				a: "435",
				b: "100",
			},
			want: "43500",
		},
		{
			name: "3",
			args: args{
				a: "251",
				b: "351",
			},
			want: "88101",
		},
		{
			name: "4",
			args: args{
				a: "2555",
				b: "3333",
			},
			want: "8515815",
		},
		{
			name: "5",
			args: args{
				a: "3141592653589793238462643383279502884197169399375105820974944592",
				b: "2718281828459045235360287471352662497757247093699959574966967627",
			},
			want: "8539734222673567065463550869546574495034888535765114961879601127067743044893204848617875072216249073013374895871952806582723184",
		},
		{
			name: "6",
			args: args{
				a: "0",
				b: "0",
			},
			want: "0",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := multiply(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("multiply() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_sum(t *testing.T) {
	type args struct {
		arrays []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "1",
			args: args{arrays: []string{"10", "10"}},
			want: "20",
		},
		{
			name: "2",
			args: args{arrays: []string{"75", "75"}},
			want: "150",
		},
		{
			name: "3",
			args: args{arrays: []string{"99", "199", "1"}},
			want: "299",
		},
		{
			name: "4",
			args: args{arrays: []string{"945", "65", "9"}},
			want: "1019",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sum(tt.args.arrays...); got != tt.want {
				t.Errorf("sum() = %v, want %v", got, tt.want)
			}
		})
	}
}
