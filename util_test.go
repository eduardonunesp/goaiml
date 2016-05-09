package goaiml

import "testing"

func Test_Util_StringMin(t *testing.T) {
	in := " a        b  c "
	out := stringMinifier(in)

	if out != " a b c " {
		t.Error("String min result error")
	}
}

func Test_Util_Pre_Format(t *testing.T) {
	in := `
    A B
    C
    `

	out := preFormatInput(in)

	if out != " A B    C " {
		t.Error("String pre format result error:", out)
	}
}

func Test_Util_Post_Format(t *testing.T) {
	in := `
        A
        B C
        D
    `

	out := postFormatInput(in)

	if out != "A B C D" {
		t.Error("String post format result error:", out)
	}
}
