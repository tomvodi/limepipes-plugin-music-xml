package timelines

import "github.com/tomvodi/limepipes-plugin-api/musicmodel/v1/measure"

// splitIntoParts groups a tune's measures into its parts.
//
// Parts are delimited by heavy barlines: a part runs up to and including the
// measure carrying a heavy right barline. Measures after the last heavy barline
// still form a group of their own so that no measure is dropped.
//
// The measures are not copied. Each group holds the same pointers as the tune,
// so editing a measure in a part edits it in the tune.
func splitIntoParts(measures []*measure.Measure) [][]*measure.Measure {
	var parts [][]*measure.Measure
	var current []*measure.Measure

	for _, m := range measures {
		current = append(current, m)

		if endsPart(m.RightBarline) {
			parts = append(parts, current)
			current = nil
		}
	}

	if len(current) > 0 {
		parts = append(parts, current)
	}

	return parts
}

// joinParts lays the parts back out as a single run of measures.
func joinParts(parts [][]*measure.Measure) []*measure.Measure {
	var measures []*measure.Measure
	for _, p := range parts {
		measures = append(measures, p...)
	}

	return measures
}
