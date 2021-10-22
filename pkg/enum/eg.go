package enum

import "strings"

type ExampleEvent uint8

const (
	Created ExampleEvent = 1
	Locked  ExampleEvent = 2
)

var enumerations = map[ExampleEvent]string{
	Created: "CREATED_EVENT",
	Locked:  "LOCKED_EVENT",
}

func (event ExampleEvent) String() string {
	return enumerations[event]
}

func (event ExampleEvent) UInt8() uint8 {
	return uint8(event)
}

func (event ExampleEvent) Int32() int32 {
	return int32(event)
}

func (event ExampleEvent) Int() int {
	return int(event)
}

func ExampleEventFromString(enum string) ExampleEvent {
	enum = strings.ToUpper(enum)
	for i, s := range enumerations {
		if s == enum {
			return ExampleEvent(i)
		}
	}
	return 0
}
