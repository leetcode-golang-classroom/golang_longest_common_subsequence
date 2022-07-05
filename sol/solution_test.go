package sol

import "testing"

func BenchmarkTest(b *testing.B) {
	text1, text2 := "abcde", "ace"
	for idx := 0; idx < b.N; idx++ {
		longestCommonSubsequence(text1, text2)
	}
}
func Test_longestCommonSubsequence(t *testing.T) {
	type args struct {
		text1 string
		text2 string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "text1 = \"abcde\", text2 = \"ace\"",
			args: args{text1: "abcde", text2: "ace"},
			want: 3,
		},
		{
			name: "text1 = \"abc\", text2 = \"abc\"",
			args: args{text1: "abc", text2: "abc"},
			want: 3,
		},
		{
			name: "text1 = \"abc\", text2 = \"def\"",
			args: args{text1: "abc", text2: "def"},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestCommonSubsequence(tt.args.text1, tt.args.text2); got != tt.want {
				t.Errorf("longestCommonSubsequence() = %v, want %v", got, tt.want)
			}
		})
	}
}
