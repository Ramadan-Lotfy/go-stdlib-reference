package reflectref_test

import (
	"fmt"
	"reflect"
	"strings"
)

type User struct {
	Name  string `json:"name" validate:"required"`
	Age   int    `json:"age"`
	email string
}

func (u User) Greet(greeting string) string { return greeting + ", " + u.Name }

// Example_typeAndKind separates a type's name from its kind.
func Example_typeAndKind() {
	u := User{"Ada", 36, ""}
	t := reflect.TypeOf(u)
	fmt.Println(t, t.Name(), t.Kind())

	pt := reflect.TypeOf(&u)
	fmt.Println(pt.Kind(), pt.Elem().Name())

	fmt.Println(reflect.TypeOf([]string{}).Elem(), reflect.TypeOf(map[string]int{}).Key())

	type UserID int
	id := reflect.TypeOf(UserID(1))
	fmt.Println(id.Name(), id.Kind())
	// Output:
	// reflectref_test.User User struct
	// ptr User
	// string string
	// UserID int
}

// Example_structFields lists fields, types and tags.
func Example_structFields() {
	t := reflect.TypeOf(User{})
	for i := range t.NumField() {
		f := t.Field(i)
		fmt.Println(f.Name, f.Type, f.Tag.Get("json"), f.IsExported())
	}
	f, _ := t.FieldByName("Name")
	fmt.Println(f.Tag.Get("validate"))
	// Output:
	// Name string name true
	// Age int age true
	// email string  false
	// required
}

// Example_setValues modifies a struct through a pointer.
func Example_setValues() {
	u := User{"Ada", 36, ""}
	v := reflect.ValueOf(&u).Elem() // pointer, then Elem: addressable
	v.FieldByName("Name").SetString("Bob")
	v.Field(1).SetInt(37)
	fmt.Println(u)

	fmt.Println(v.FieldByName("email").CanSet())      // unexported
	fmt.Println(reflect.ValueOf(u).Field(0).CanSet()) // not addressable
	// Output:
	// {Bob 37 }
	// false
	// false
}

// Example_methods calls a method by name.
func Example_methods() {
	u := User{Name: "Ada"}
	m := reflect.ValueOf(u).MethodByName("Greet")
	out := m.Call([]reflect.Value{reflect.ValueOf("Hello")})
	fmt.Println(out[0].String(), reflect.TypeOf(u).NumMethod())
	// Output: Hello, Ada 1
}

// Example_makeValues builds slices, maps and pointers at run time.
func Example_makeValues() {
	s := reflect.MakeSlice(reflect.TypeOf([]int{}), 0, 2)
	s = reflect.Append(s, reflect.ValueOf(1), reflect.ValueOf(2))
	fmt.Println(s.Interface(), s.Len())

	m := reflect.MakeMap(reflect.TypeOf(map[string]int{}))
	m.SetMapIndex(reflect.ValueOf("a"), reflect.ValueOf(1))
	fmt.Println(m.Interface())

	p := reflect.New(reflect.TypeOf(0)) // *int
	p.Elem().SetInt(7)
	fmt.Println(*(p.Interface().(*int)))
	// Output:
	// [1 2] 2
	// map[a:1]
	// 7
}

// describe walks any value by Kind. Real libraries (JSON, validators) do this.
func describe(v reflect.Value) string {
	switch v.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return fmt.Sprintf("int(%d)", v.Int())
	case reflect.String:
		return fmt.Sprintf("string(%q)", v.String())
	case reflect.Bool:
		return fmt.Sprintf("bool(%t)", v.Bool())
	case reflect.Slice:
		parts := make([]string, v.Len())
		for i := range parts {
			parts[i] = describe(v.Index(i))
		}
		return "[" + strings.Join(parts, ", ") + "]"
	case reflect.Struct:
		var parts []string
		for i := range v.NumField() {
			if f := v.Type().Field(i); f.IsExported() {
				parts = append(parts, f.Name+": "+describe(v.Field(i)))
			}
		}
		return "{" + strings.Join(parts, ", ") + "}"
	case reflect.Pointer, reflect.Interface:
		if v.IsNil() {
			return "nil"
		}
		return describe(v.Elem())
	default:
		return "unsupported " + v.Kind().String()
	}
}

// Example_walk handles unknown shapes by switching on Kind.
func Example_walk() {
	fmt.Println(describe(reflect.ValueOf([]int{1, 2})))
	fmt.Println(describe(reflect.ValueOf(User{"Ada", 36, "x@y"})))
	fmt.Println(describe(reflect.ValueOf(&User{Name: "Bob"})))
	fmt.Println(describe(reflect.ValueOf(3.5)))
	// Output:
	// [int(1), int(2)]
	// {Name: string("Ada"), Age: int(36)}
	// {Name: string("Bob"), Age: int(0)}
	// unsupported float64
}

// Example_typeFor checks interface implementation without a value.
func Example_typeFor() {
	stringer := reflect.TypeFor[fmt.Stringer]()
	fmt.Println(reflect.TypeFor[int]())
	fmt.Println(reflect.TypeOf(User{}).Implements(stringer))
	fmt.Println(reflect.TypeOf(reflect.TypeOf(0)).Implements(stringer)) // *reflect.rtype has String()
	// Output:
	// int
	// false
	// true
}

// Example_deepEqual compares structures; note nil vs empty.
func Example_deepEqual() {
	fmt.Println(reflect.DeepEqual(map[string][]int{"a": {1}}, map[string][]int{"a": {1}}))
	fmt.Println(reflect.DeepEqual([]int(nil), []int{}))
	// Output:
	// true
	// false
}

// Example_panics shows that reflection panics on misuse instead of returning errors.
func Example_panics() {
	defer func() { fmt.Println("recovered:", recover()) }()
	reflect.ValueOf(42).SetInt(1) // not addressable
	// Output: recovered: reflect: reflect.Value.SetInt using unaddressable value
}
