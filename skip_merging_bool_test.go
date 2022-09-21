package mergo

import "testing"

type testStruct struct {
	BoolField bool
}

func TestMergeBool(t *testing.T) {
	dst := testStruct{
		BoolField: false,
	}
	src := testStruct{
		BoolField: true,
	}

	if err := Merge(&dst, src); err != nil {
		t.Error(err)
	}

	if dst.BoolField != src.BoolField {
		t.Error("dst is supposed to be the same as src")
	}
}

func TestSkipMergeBool(t *testing.T) {
	dst := testStruct{
		BoolField: false,
	}
	src := testStruct{
		BoolField: true,
	}

	if err := Merge(&dst, src, WithSkipMergingBool); err != nil {
		t.Error(err)
	}

	if dst.BoolField == src.BoolField {
		t.Error("dst is supposed to be different than src")
	}
}
