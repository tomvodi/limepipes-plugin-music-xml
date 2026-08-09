package timelines

import (
	"github.com/tomvodi/limepipes-plugin-api/musicmodel/v1/tune"
	"google.golang.org/protobuf/proto"
)

// Normalize returns a copy of the tune whose time lines can be written as
// MusicXML endings.
//
// Two things are rewritten, in this order:
//
//  1. Every "second of N" bracket is moved to the part it is played in, since
//     MusicXML cannot refer to a passage that lives somewhere else.
//  2. Every bracket is made to span whole bars, since a MusicXML ending begins
//     and ends at a barline.
//
// The tune passed in is not modified. Plugins are handed their tunes over gRPC
// and the host may keep using them afterwards.
func Normalize(t *tune.Tune) *tune.Tune {
	normalized, ok := proto.Clone(t).(*tune.Tune)
	if !ok {
		// proto.Clone returns the same concrete type it was given, so this is
		// unreachable; falling back to the original keeps the export working.
		return t
	}

	parts := splitIntoParts(normalized.Measures)
	parts = distributeTimelines(parts)

	normalized.Measures = splitMeasuresAtTimelines(joinParts(parts))

	return normalized
}
