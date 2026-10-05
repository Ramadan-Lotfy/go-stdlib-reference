package solution

import (
	"reflect"
	"slices"
	"testing"
)

type signup struct {
	Name   string `json:"name" validate:"required"`
	Email  string `json:"email,omitempty" validate:"required"`
	Age    int    `json:"age"`
	Nick   string
	secret string
	Skip   string `json:"-"`
}

func TestStructToMap(t *testing.T) {
	in := signup{Name: "Ada", Email: "a@x.com", Age: 36, Nick: "ada", secret: "s", Skip: "z"}
	want := map[string]any{"name": "Ada", "email": "a@x.com", "age": 36, "Nick": "ada"}
	for _, v := range []any{in, &in} {
		if got := StructToMap(v); !reflect.DeepEqual(got, want) {
			t.Errorf("StructToMap(%T) = %v, want %v", v, got, want)
		}
	}
	if StructToMap(42) != nil || StructToMap(nil) != nil || StructToMap((*signup)(nil)) != nil {
		t.Error("non-structs and nil should give nil")
	}
}

func TestMissingRequired(t *testing.T) {
	got := MissingRequired(signup{Name: "Ada"})
	if want := []string{"Email"}; !slices.Equal(got, want) {
		t.Errorf("MissingRequired() = %v, want %v", got, want)
	}
	if got := MissingRequired(&signup{}); !slices.Equal(got, []string{"Name", "Email"}) {
		t.Errorf("MissingRequired(empty) = %v", got)
	}
	if got := MissingRequired(signup{Name: "A", Email: "b"}); len(got) != 0 {
		t.Errorf("complete struct reported %v", got)
	}
}
