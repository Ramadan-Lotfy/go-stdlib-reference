package solution

import "testing"

func TestAddQuery(t *testing.T) {
	got, err := AddQuery("https://x.com/p?a=1", map[string]string{"b": "2 3"})
	if err != nil || got != "https://x.com/p?a=1&b=2+3" {
		t.Errorf("AddQuery() = %q, %v", got, err)
	}

	got, _ = AddQuery("https://x.com/p?a=1", map[string]string{"a": "9"})
	if got != "https://x.com/p?a=9" {
		t.Errorf("replace: %q", got)
	}

	got, _ = AddQuery("https://x.com/p#top", map[string]string{"q": "go"})
	if got != "https://x.com/p?q=go#top" {
		t.Errorf("fragment kept: %q", got)
	}

	if _, err := AddQuery("x.com/p", nil); err == nil {
		t.Error("URL without scheme should fail")
	}
}

func TestDomain(t *testing.T) {
	got, err := Domain("https://user:pw@example.com:8443/path")
	if err != nil || got != "example.com" {
		t.Errorf("Domain() = %q, %v", got, err)
	}
	if _, err := Domain("/just/a/path"); err == nil {
		t.Error("no host should fail")
	}
}
