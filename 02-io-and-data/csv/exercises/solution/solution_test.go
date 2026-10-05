package solution

import (
	"reflect"
	"strings"
	"testing"
)

func TestParseUsers(t *testing.T) {
	got, err := ParseUsers(strings.NewReader("Age,NAME\n36,Ada\n41,\"Smith, J\"\n"))
	want := []User{{"Ada", 36}, {"Smith, J", 41}}
	if err != nil || !reflect.DeepEqual(got, want) {
		t.Errorf("ParseUsers() = %v, %v; want %v", got, err, want)
	}
}

func TestParseUsersErrors(t *testing.T) {
	_, err := ParseUsers(strings.NewReader("name,age\nAda,36\nBob,x\n"))
	if err == nil || !strings.Contains(err.Error(), "line 3") {
		t.Errorf("want error mentioning line 3, got %v", err)
	}
	if _, err := ParseUsers(strings.NewReader("name,city\nAda,Cairo\n")); err == nil {
		t.Error("missing age column should fail")
	}
	if _, err := ParseUsers(strings.NewReader("")); err == nil {
		t.Error("empty input should fail")
	}
}
