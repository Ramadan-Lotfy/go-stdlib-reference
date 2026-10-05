package solution

import "testing"

func TestRenderComments(t *testing.T) {
	got, err := RenderComments([]string{"nice", "<script>x</script>", "a & b"})
	want := "<ul><li>nice</li><li>&lt;script&gt;x&lt;/script&gt;</li><li>a &amp; b</li></ul>"
	if err != nil || got != want {
		t.Errorf("RenderComments() = %q, %v\nwant %q", got, err, want)
	}

	got, _ = RenderComments(nil)
	if got != "<ul></ul>" {
		t.Errorf("RenderComments(nil) = %q", got)
	}
}
