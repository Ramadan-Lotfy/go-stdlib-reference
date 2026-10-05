package solution

import "testing"

func TestParseLogLine(t *testing.T) {
	f, ok := ParseLogLine("2026-10-03 ERROR [db] connection lost")
	if !ok {
		t.Fatal("line should match")
	}
	want := map[string]string{"date": "2026-10-03", "level": "ERROR", "comp": "db", "msg": "connection lost"}
	for k, v := range want {
		if f[k] != v {
			t.Errorf("field %s = %q, want %q", k, f[k], v)
		}
	}
	if _, ok := ParseLogLine("garbage line"); ok {
		t.Error("garbage should not match")
	}
}

func TestRedactDigits(t *testing.T) {
	tests := []struct{ in, want string }{
		{"card 1234567890 id 42", "card ********** id 42"},
		{"pin 123", "pin 123"},
		{"a1111b22222c", "a****b*****c"},
		{"", ""},
	}
	for _, tt := range tests {
		if got := RedactDigits(tt.in); got != tt.want {
			t.Errorf("RedactDigits(%q) = %q, want %q", tt.in, got, tt.want)
		}
	}
}
