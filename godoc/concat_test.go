package godoc

import (
	"testing"
)

func TestConcat(t *testing.T){
	// Note: you *must* have at least one unit test or examples will not render
	// this is possibly a bug.
	a := "foo"
	b := "bar"
	result := Concat(a,b)
	expect := "foobar"
	if result != expect {
          t.Errorf("expected: %s, got: %s",expect,result)
	}
}

func ExampleConcat() {
	// This will render in your godoc documentation and will be listed in the "Examples"
	a := "foo"
	b := "bar"
	result := Concat(a,b)
	expect := "foobar"

	fmt.Printf("expected: %s, got: %s\n",expect,result)
	// Output: "expected: foobar, got: foobar"
}


