package solution

import "testing"

func TestParseConfig(t *testing.T) {
	cfg, err := ParseConfig([]byte(`{"debug": true}`))
	if err != nil || cfg != (Config{"localhost", 8080, true}) {
		t.Errorf("defaults: %+v, %v", cfg, err)
	}

	cfg, err = ParseConfig([]byte(`{"host":"example.com","port":9000}`))
	if err != nil || cfg.Host != "example.com" || cfg.Port != 9000 {
		t.Errorf("explicit: %+v, %v", cfg, err)
	}
}

func TestParseConfigErrors(t *testing.T) {
	bad := []string{
		`{"hots":"typo"}`,      // unknown field
		`{"port":0}`,           // out of range
		`{"port":70000}`,       // out of range
		`{"port":"80"}`,        // wrong type
		`{"host":"a"} {"x":1}`, // trailing data
		`not json`,
		``,
	}
	for _, in := range bad {
		if _, err := ParseConfig([]byte(in)); err == nil {
			t.Errorf("ParseConfig(%q) should fail", in)
		}
	}
}

func TestPrettyJSON(t *testing.T) {
	got, err := PrettyJSON([]byte(`{"a":[1,2]}`))
	want := "{\n  \"a\": [\n    1,\n    2\n  ]\n}"
	if err != nil || got != want {
		t.Errorf("PrettyJSON() = %q, %v", got, err)
	}
	if _, err := PrettyJSON([]byte(`{`)); err == nil {
		t.Error("invalid JSON should fail")
	}
}
