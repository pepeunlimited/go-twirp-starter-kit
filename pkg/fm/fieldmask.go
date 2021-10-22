package fm

import "google.golang.org/protobuf/types/known/fieldmaskpb"

func IsNilOrPathEmpty(fieldMask *fieldmaskpb.FieldMask) bool {
	return fieldMask == nil || len(fieldMask.Paths) == 0
}