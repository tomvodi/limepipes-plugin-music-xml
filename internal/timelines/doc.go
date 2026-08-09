// Package timelines rewrites a tune's time lines into a form MusicXML can
// express.
//
// # Time lines in bagpipe music
//
// A time line is the bracket drawn over a passage that is only played on some
// times through a repeat. The familiar case is a first and second ending:
//
//	B  ⌐1———¬  ⌐2———¬
//	   | D  |  | E  |   played D the first time, E the second
//
// bww writes those as '1 D_4 _' and '2 E_4 _'. MusicXML writes them as
// <ending number="1"> elements hanging off barlines, so the two notations line
// up and the conversion is mechanical.
//
// # The two things MusicXML cannot take literally
//
// First, bww lets a bracket open and close in the middle of a bar. That happens
// all the time when the first and second endings have different upbeats. A
// MusicXML ending has to begin and end at a barline, so such a bar is cut into
// several bars at the bracket boundaries. See splitMeasuresAtTimelines.
//
// Second, and this is the interesting one, bww has time lines that are written
// in one part but played in another. '22 means "second time of part 2", '24
// means "second time of part 4", and so on. The bracket is drawn wherever the
// engraver had room — usually in an earlier part — but the passage belongs to
// the repeat of the part named in the symbol. MusicXML has no way to say
// "somewhere else", so before converting, the passage is moved to the part it
// is actually played in. See distributeTimelines.
//
// This second case is what the bww2mxml tool that ships with MuseScore does not
// handle, and it is the reason this package exists.
//
// # What Normalize does
//
//	'22 written in part 1          part 1 loses the bracket, keeps the notes
//	part 2 has a '1 bracket   -->  part 2 gains a '2 bracket holding those notes
//	                               every bracket now spans whole bars
//
// The result is a tune whose time lines map one-to-one onto MusicXML endings.
package timelines
