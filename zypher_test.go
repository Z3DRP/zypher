package zypher

import "testing"

func TestAsciZyphNoIgnoreSpace(t *testing.T) {
	v := "Abc Z19"
	zy := NewZypher(WithIterCount(1))
	reslt, _ := zy.AsciZyph(v)
	want := "DefxC42"

	if reslt != want {
		t.Errorf("got %q, wanted %q", reslt, want)
	}
}

func TestAsciZyphIgnoreSpace(t *testing.T) {
	v := "Abc Z19"
	zy := NewZypher(WithIterCount(1), WithIgnoreSpace(true))
	reslt, _ := zy.AsciZyph(v)
	want := "Def C42"

	if reslt != want {
		t.Errorf("got %q, wanted %q", reslt, want)
	}
}
