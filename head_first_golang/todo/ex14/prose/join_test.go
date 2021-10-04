package prose

import "testing"

type testData struct {
	list []string
	want string
}

func TestJoinWithCommas(t *testing.T) {
	tests := []testData{
		{
			list: []string{},
			want: "",
		},
		{
			list: []string{"apple"},
			want: "apple",
		},
		{
			list: []string{"apple", "orange"},
			want: "apple and orange",
		},
		{
			list: []string{"apple", "orange", "pear"},
			want: "apple, orange, and pear",
		},
	}
	for _, test := range tests {
		got := JoinWithCommas(test.list)
		if got != test.want {
			errorString(t, test.list, got, test.want)
		}
	}
}

func errorString(t *testing.T, list []string, got, want string) {
	t.Errorf("JoinWithCommas(%#v) = \"%s\", want \"%s\"", list, got, want)
}
