package coursera

import (
	"reflect"
	"testing"
)

// *************** week1 *******************

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

// *************** week2 *******************

func Test_countSplitInv(t *testing.T) {
	type args struct {
		arr1 []int
		arr2 []int
	}
	tests := []struct {
		name  string
		args  args
		want  []int
		want1 int
	}{
		{
			name: "1",
			args: args{
				arr1: []int{1, 3, 5},
				arr2: []int{2, 4, 6},
			},
			want:  []int{1, 2, 3, 4, 5, 6},
			want1: 3,
		},
		{
			name: "2",
			args: args{
				arr1: []int{1, 2, 3},
				arr2: []int{4, 5, 6, 7},
			},
			want:  []int{1, 2, 3, 4, 5, 6, 7},
			want1: 0,
		},
		{
			name: "3",
			args: args{
				arr1: []int{1, 3, 5, 6},
				arr2: []int{2, 4},
			},
			want:  []int{1, 2, 3, 4, 5, 6},
			want1: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := countSplitInv(tt.args.arr1, tt.args.arr2)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("countSplitInv() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("countSplitInv() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func Test_countSplitInversions(t *testing.T) {

	// path := "https://d3c33hcgiwev3.cloudfront.net/_bcb5c6658381416d19b01bfc1d3993b5_IntegerArray.txt?Expires=1606176000&Signature=Ecoaj6dddKxGM6DfuSCf1prN1fVD2khs3G-0N50q6LBz7-qZYU7Reptb78p82Nf0vZrRE~-Rij1REQ6IzEC2Umhduy8Mm6jCi9iIOmk8f1~PoMD7LSXIHzZ7mcvBkYjRwEb6E4HOfDXokWn462k3jqfzrKf4F8LSPO0yYuuVpUU_&Key-Pair-Id=APKAJLTNE6QMUY6HBC5A"
	// res, _ := http.Get(path)
	// b, _ := ioutil.ReadAll(res.Body)
	// nums := strings.Split(string(b), "\r\n")
	// nums = nums[:len(nums)-1]
	// arr := make([]int, len(nums), len(nums))

	// var err error
	// for i := 0; i < len(nums); i++ {
	// 	arr[i], err = strconv.Atoi(nums[i])
	// 	if err != nil {
	// 		fmt.Println(err, i)
	// 	}
	// }
	// arrCopy := make([]int, len(nums), len(nums))
	// copy(arrCopy, arr)
	// sort.Ints(arrCopy)

	type args struct {
		arr []int
	}
	tests := []struct {
		name  string
		args  args
		want  []int
		want1 int
	}{
		{name: "1", args: args{arr: []int{1, 3, 5, 2, 4, 6}}, want: []int{1, 2, 3, 4, 5, 6}, want1: 3},
		{name: "2", args: args{arr: []int{1, 2, 3, 4, 5, 6}}, want: []int{1, 2, 3, 4, 5, 6}, want1: 0},
		{name: "3", args: args{arr: []int{6, 5, 4, 3, 2, 1}}, want: []int{1, 2, 3, 4, 5, 6}, want1: 15},
		{name: "4", args: args{arr: []int{1, 2, 3}}, want: []int{1, 2, 3}, want1: 0},
		//{name: "coursera", args: args{arr: arr}, want: arrCopy, want1: 2407905288},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := countSplitInversions(tt.args.arr)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("countSplitInversions() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("countSplitInversions() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}

}
