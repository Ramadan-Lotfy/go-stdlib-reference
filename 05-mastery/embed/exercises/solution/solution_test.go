package solution

import (
	"slices"
	"testing"
)

func TestRender(t *testing.T) {
	got, err := Render("greeting", "Ada")
	if err != nil || got != "Hello, Ada!" {
		t.Errorf("Render(greeting) = %q, %v", got, err)
	}
	got, _ = Render("farewell", "Bob")
	if got != "Goodbye, Bob!" {
		t.Errorf("Render(farewell) = %q", got)
	}
	if _, err := Render("nope", "x"); err == nil {
		t.Error("unknown template should fail")
	}
}

func TestNames(t *testing.T) {
	if got := Names(); !slices.Equal(got, []string{"farewell", "greeting"}) {
		t.Errorf("Names() = %v", got)
	}
}
