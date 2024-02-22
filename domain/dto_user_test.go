package domain

import (
	"strings"
	"testing"
)

func TestIsValidPhoneNumber(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  bool
	}{
		{name: "PASS - 유효한 번호, 하이픈 없음", input: "01012345678", want: true},
		{name: "PASS - 유효한 번호, 하이픈 있음", input: "010-1234-5678", want: true},
		{name: "FAIL - 유효하지 않은 번호, 하이픈이 잘못됨", input: "010-12345678", want: false},
		{name: "FAIL - 유효하지 않은 번호, 너무 짧음", input: "0101234", want: false},
		{name: "FAIL - 유효하지 않은 번호, 잘못된 문자 포함", input: "010-1234-abcd", want: false},
		{name: "FAIL - 빈 문자열", input: "", want: false},
		{name: "FAIL - 유효하지 않은 번호, 잘못된 하이픈", input: "0101234-5678", want: false},
	}

	for _, test := range tests {
		actualValid := isValidMobileID(test.input)
		if actualValid != test.want {
			t.Errorf("Expected validity for input %s to be %t, but got %t", test.input, test.want, actualValid)
		}
	}
}

func Test_isValidPassword(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  bool
	}{
		{
			name:  "PASS - 영어 소문자 한글자",
			input: "p",
			want:  true,
		},
		{
			name:  "PASS - 영어 대문자 한글자",
			input: "P",
			want:  true,
		},
		{
			name:  "PASS - 숫자 한글자",
			input: "5",
			want:  true,
		},
		{
			name:  "PASS - 특수 기호 한글자",
			input: "@",
			want:  true,
		},
		{
			name:  "PASS - 255자 패스워드",
			input: "payhere" + strings.Repeat("x", 248),
			want:  true,
		},
		{
			name:  "FAIL – 0자 패스워드",
			input: "",
			want:  false,
		},
		{
			name:  "FAIL - 256자 패스워드",
			input: "payhere" + strings.Repeat("x", 249),
			want:  false,
		},
	}
	for _, test := range tests {
		actualValid := isValidPassword(test.input)
		if actualValid != test.want {
			t.Errorf("Expected validity for input %s to be %t, but got %t", test.input, test.want, actualValid)
		}
	}
}
