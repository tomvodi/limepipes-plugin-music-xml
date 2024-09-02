// Code generated by "enumer -json -yaml -transform=kebab -type=Style"; DO NOT EDIT.

package barline

import (
	"encoding/json"
	"fmt"
	"strings"
)

const _StyleName = "noneregulardasheddottedheavyheavy-heavyheavy-lightlight-heavylight-lightshorttick"

var _StyleIndex = [...]uint8{0, 4, 11, 17, 23, 28, 39, 50, 61, 72, 77, 81}

const _StyleLowerName = "noneregulardasheddottedheavyheavy-heavyheavy-lightlight-heavylight-lightshorttick"

func (i Style) String() string {
	if i >= Style(len(_StyleIndex)-1) {
		return fmt.Sprintf("Style(%d)", i)
	}
	return _StyleName[_StyleIndex[i]:_StyleIndex[i+1]]
}

// An "invalid array index" compiler error signifies that the constant values have changed.
// Re-run the stringer command to generate them again.
func _StyleNoOp() {
	var x [1]struct{}
	_ = x[None-(0)]
	_ = x[Regular-(1)]
	_ = x[Dashed-(2)]
	_ = x[Dotted-(3)]
	_ = x[Heavy-(4)]
	_ = x[HeavyHeavy-(5)]
	_ = x[HeavyLight-(6)]
	_ = x[LightHeavy-(7)]
	_ = x[LightLight-(8)]
	_ = x[Short-(9)]
	_ = x[Tick-(10)]
}

var _StyleValues = []Style{None, Regular, Dashed, Dotted, Heavy, HeavyHeavy, HeavyLight, LightHeavy, LightLight, Short, Tick}

var _StyleNameToValueMap = map[string]Style{
	_StyleName[0:4]:        None,
	_StyleLowerName[0:4]:   None,
	_StyleName[4:11]:       Regular,
	_StyleLowerName[4:11]:  Regular,
	_StyleName[11:17]:      Dashed,
	_StyleLowerName[11:17]: Dashed,
	_StyleName[17:23]:      Dotted,
	_StyleLowerName[17:23]: Dotted,
	_StyleName[23:28]:      Heavy,
	_StyleLowerName[23:28]: Heavy,
	_StyleName[28:39]:      HeavyHeavy,
	_StyleLowerName[28:39]: HeavyHeavy,
	_StyleName[39:50]:      HeavyLight,
	_StyleLowerName[39:50]: HeavyLight,
	_StyleName[50:61]:      LightHeavy,
	_StyleLowerName[50:61]: LightHeavy,
	_StyleName[61:72]:      LightLight,
	_StyleLowerName[61:72]: LightLight,
	_StyleName[72:77]:      Short,
	_StyleLowerName[72:77]: Short,
	_StyleName[77:81]:      Tick,
	_StyleLowerName[77:81]: Tick,
}

var _StyleNames = []string{
	_StyleName[0:4],
	_StyleName[4:11],
	_StyleName[11:17],
	_StyleName[17:23],
	_StyleName[23:28],
	_StyleName[28:39],
	_StyleName[39:50],
	_StyleName[50:61],
	_StyleName[61:72],
	_StyleName[72:77],
	_StyleName[77:81],
}

// StyleString retrieves an enum value from the enum constants string name.
// Throws an error if the param is not part of the enum.
func StyleString(s string) (Style, error) {
	if val, ok := _StyleNameToValueMap[s]; ok {
		return val, nil
	}

	if val, ok := _StyleNameToValueMap[strings.ToLower(s)]; ok {
		return val, nil
	}
	return 0, fmt.Errorf("%s does not belong to Style values", s)
}

// StyleValues returns all values of the enum
func StyleValues() []Style {
	return _StyleValues
}

// StyleStrings returns a slice of all String values of the enum
func StyleStrings() []string {
	strs := make([]string, len(_StyleNames))
	copy(strs, _StyleNames)
	return strs
}

// IsAStyle returns "true" if the value is listed in the enum definition. "false" otherwise
func (i Style) IsAStyle() bool {
	for _, v := range _StyleValues {
		if i == v {
			return true
		}
	}
	return false
}

// MarshalJSON implements the json.Marshaler interface for Style
func (i Style) MarshalJSON() ([]byte, error) {
	return json.Marshal(i.String())
}

// UnmarshalJSON implements the json.Unmarshaler interface for Style
func (i *Style) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return fmt.Errorf("Style should be a string, got %s", data)
	}

	var err error
	*i, err = StyleString(s)
	return err
}

// MarshalYAML implements a YAML Marshaler for Style
func (i Style) MarshalYAML() (interface{}, error) {
	return i.String(), nil
}

// UnmarshalYAML implements a YAML Unmarshaler for Style
func (i *Style) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var s string
	if err := unmarshal(&s); err != nil {
		return err
	}

	var err error
	*i, err = StyleString(s)
	return err
}
