package solution

import (
	"slices"
	"testing"
)

func TestFeedTitles(t *testing.T) {
	data := []byte(`<feed>
  <entry><title>One</title></entry>
  <entry><title>Two &amp; Three</title></entry>
</feed>`)
	got, err := FeedTitles(data)
	want := []string{"One", "Two & Three"}
	if err != nil || !slices.Equal(got, want) {
		t.Errorf("FeedTitles() = %v, %v; want %v", got, err, want)
	}

	got, err = FeedTitles([]byte(`<feed></feed>`))
	if err != nil || len(got) != 0 {
		t.Errorf("empty feed: %v, %v", got, err)
	}

	if _, err := FeedTitles([]byte(`<feed><entry>`)); err == nil {
		t.Error("malformed XML should fail")
	}
}
