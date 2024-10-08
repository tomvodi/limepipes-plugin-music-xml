// Code generated by "enumer -transform=lower -type=NumberStyle"; DO NOT EDIT.

package tuplet

import (
	"fmt"
	"strings"
)

const _NumberStyleName = "invisiblenonebothactual"

var _NumberStyleIndex = [...]uint8{0, 9, 13, 17, 23}

const _NumberStyleLowerName = "invisiblenonebothactual"

func (i NumberStyle) String() string {
	if i >= NumberStyle(len(_NumberStyleIndex)-1) {
		return fmt.Sprintf("NumberStyle(%d)", i)
	}
	return _NumberStyleName[_NumberStyleIndex[i]:_NumberStyleIndex[i+1]]
}

// An "invalid array index" compiler error signifies that the constant values have changed.
// Re-run the stringer command to generate them again.
func _NumberStyleNoOp() {
	var x [1]struct{}
	_ = x[Invisible-(0)]
	_ = x[None-(1)]
	_ = x[Both-(2)]
	_ = x[Actual-(3)]
}

var _NumberStyleValues = []NumberStyle{Invisible, None, Both, Actual}

var _NumberStyleNameToValueMap = map[string]NumberStyle{
	_NumberStyleName[0:9]:        Invisible,
	_NumberStyleLowerName[0:9]:   Invisible,
	_NumberStyleName[9:13]:       None,
	_NumberStyleLowerName[9:13]:  None,
	_NumberStyleName[13:17]:      Both,
	_NumberStyleLowerName[13:17]: Both,
	_NumberStyleName[17:23]:      Actual,
	_NumberStyleLowerName[17:23]: Actual,
}

var _NumberStyleNames = []string{
	_NumberStyleName[0:9],
	_NumberStyleName[9:13],
	_NumberStyleName[13:17],
	_NumberStyleName[17:23],
}

// NumberStyleString retrieves an enum value from the enum constants string name.
// Throws an error if the param is not part of the enum.
func NumberStyleString(s string) (NumberStyle, error) {
	if val, ok := _NumberStyleNameToValueMap[s]; ok {
		return val, nil
	}

	if val, ok := _NumberStyleNameToValueMap[strings.ToLower(s)]; ok {
		return val, nil
	}
	return 0, fmt.Errorf("%s does not belong to NumberStyle values", s)
}

// NumberStyleValues returns all values of the enum
func NumberStyleValues() []NumberStyle {
	return _NumberStyleValues
}

// NumberStyleStrings returns a slice of all String values of the enum
func NumberStyleStrings() []string {
	strs := make([]string, len(_NumberStyleNames))
	copy(strs, _NumberStyleNames)
	return strs
}

// IsANumberStyle returns "true" if the value is listed in the enum definition. "false" otherwise
func (i NumberStyle) IsANumberStyle() bool {
	for _, v := range _NumberStyleValues {
		if i == v {
			return true
		}
	}
	return false
}
