package windowtype

import "fmt"

type WindowType int

const (
	Managed WindowType = iota
	Floating
)

var wtToString = map[WindowType]string{
	Managed:  "managed",
	Floating: "floating",
}

var wtToID = map[string]WindowType{
	wtToString[Managed]:  Managed,
	wtToString[Floating]: Floating,
}

func (wt WindowType) String() string {
	return wtToString[wt]
}

func (wt WindowType) MarshalYAML() ([]byte, error) {
	return []byte(wtToString[wt]), nil
}

func (wt *WindowType) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var val string
	if err := unmarshal(&val); err != nil {
		return err
	}
	if conv, ok := wtToID[val]; ok {
		// Get the memory address of wt (dereference it)
		// Then stick the value in that memory. (Yes, I still have to stop and think about how pointers work.)
		*wt = conv
		return nil
	}
	return fmt.Errorf("window type must be one of: ['managed', 'floating'], not: %s", val)
}
