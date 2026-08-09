package barline

import (
	"encoding/xml"
	"strconv"
	"strings"
)

// EndingType says which end of the bracket a barline carries.
type EndingType uint8

const (
	// Start opens the bracket, on the left barline of its first bar.
	Start EndingType = iota
	// Stop closes it with a downward hook, on the right barline of its last bar.
	Stop
	// Discontinue closes it without a hook, used for the last ending of a
	// repeat where the music simply carries on.
	Discontinue
)

func (e EndingType) String() string {
	switch e {
	case Start:
		return "start"
	case Stop:
		return "stop"
	case Discontinue:
		return "discontinue"
	default:
		return ""
	}
}

// Ending is the bracket over the bars played on only some times through a
// repeat — a first or second time.
type Ending struct {
	XMLName xml.Name `xml:"ending"`

	// Number is the list of times through the repeat this bracket applies to,
	// comma separated. Usually a single "1" or "2", but a passage played on the
	// first and third time through is "1,3".
	Number string `xml:"number,attr"`
	Type   string `xml:"type,attr"`
}

// NewEnding builds a bracket end for the given times through the repeat.
func NewEnding(times []int, endingType EndingType) *Ending {
	return &Ending{
		XMLName: xml.Name{Local: "ending"},
		Number:  formatTimes(times),
		Type:    endingType.String(),
	}
}

func formatTimes(times []int) string {
	asText := make([]string, 0, len(times))
	for _, t := range times {
		asText = append(asText, strconv.Itoa(t))
	}

	return strings.Join(asText, ",")
}
