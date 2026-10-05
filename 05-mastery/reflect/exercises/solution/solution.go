// Package solution holds the reference solutions for the reflect exercises.
package solution

import (
	"reflect"
	"strings"
)

// StructToMap converts a struct (or pointer to struct) to a map keyed by the
// json tag name (or the field name if there is no tag). Unexported fields and
// fields tagged `json:"-"` are skipped. Other kinds return nil.
func StructToMap(v any) map[string]any {
	rv := reflect.ValueOf(v)
	if rv.Kind() == reflect.Pointer {
		if rv.IsNil() {
			return nil
		}
		rv = rv.Elem()
	}
	if rv.Kind() != reflect.Struct {
		return nil
	}
	rt := rv.Type()
	out := make(map[string]any)
	for i := range rt.NumField() {
		f := rt.Field(i)
		if !f.IsExported() {
			continue
		}
		name, _, _ := strings.Cut(f.Tag.Get("json"), ",")
		if name == "-" {
			continue
		}
		if name == "" {
			name = f.Name
		}
		out[name] = rv.Field(i).Interface()
	}
	return out
}

// MissingRequired returns the names of exported fields tagged
// `validate:"required"` that still have their zero value.
func MissingRequired(v any) []string {
	rv := reflect.ValueOf(v)
	if rv.Kind() == reflect.Pointer {
		if rv.IsNil() {
			return nil
		}
		rv = rv.Elem()
	}
	if rv.Kind() != reflect.Struct {
		return nil
	}
	var missing []string
	rt := rv.Type()
	for i := range rt.NumField() {
		f := rt.Field(i)
		if f.IsExported() && f.Tag.Get("validate") == "required" && rv.Field(i).IsZero() {
			missing = append(missing, f.Name)
		}
	}
	return missing
}
