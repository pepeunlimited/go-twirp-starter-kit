package fm

import (
	"google.golang.org/protobuf/types/known/fieldmaskpb"
	"testing"
)

func TestIsNilOrEmpty(t *testing.T) {
	fm0 := &fieldmaskpb.FieldMask{}
	if !IsNilOrPathEmpty(fm0) { // true
		t.FailNow()
	}
	if !IsNilOrPathEmpty(nil) { // true
		t.FailNow()
	}
	fm0.Paths = []string{}
	if !IsNilOrPathEmpty(fm0) { // true
		t.FailNow()
	}
	fm0.Paths = []string{"hola"}
	if IsNilOrPathEmpty(fm0) { // false
		t.FailNow()
	}
}
