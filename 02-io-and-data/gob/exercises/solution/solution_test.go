package solution

import (
	"reflect"
	"testing"
)

type Order struct {
	ID    int
	Items []string
	Meta  map[string]string
}

func TestDeepCopy(t *testing.T) {
	orig := Order{ID: 1, Items: []string{"a", "b"}, Meta: map[string]string{"k": "v"}}
	cp, err := DeepCopy(orig)
	if err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(orig, cp) {
		t.Fatalf("copy = %+v, want %+v", cp, orig)
	}

	cp.Items[0] = "changed"
	cp.Meta["k"] = "changed"
	if orig.Items[0] != "a" || orig.Meta["k"] != "v" {
		t.Errorf("original was modified: %+v", orig)
	}
}
