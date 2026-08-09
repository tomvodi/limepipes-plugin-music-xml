# limepipes-plugin-music-xml

A [LimePipes](https://github.com/tomvodi/limepipes) plugin that exports tunes to
the [MusicXML](https://www.musicxml.com) format, so that music parsed from
Bagpipe Music Writer files can be opened in notation programs such as
[MuseScore](https://musescore.org), Sibelius or Finale.

The plugin is export only. It advertises itself as `PluginType_OUT` and rejects
attempts to parse MusicXML: `Export` and `ExportToFile` do the work, `Parse` and
`ParseFromFile` return an error. MusicXML holds a single tune per document, so
exporting more than one tune at a time is rejected as well.

Like the other LimePipes plugins it talks to the host over
[gRPC](https://grpc.io/) using [Hashicorp's Go Plugin System](https://github.com/hashicorp/go-plugin),
and takes its music model from the
[LimePipes plugin API](https://github.com/tomvodi/limepipes-plugin-api).

## Time lines

MuseScore ships a `bww2mxml` tool that already converts basic tunes. What it
does not handle, and what this plugin exists for, are the time lines bagpipe
music uses for repeats.

Two of them need rewriting before MusicXML can express them:

- A bracket may open or close in the middle of a bar, which happens whenever the
  first and second endings have different upbeats. A MusicXML ending has to
  begin and end at a barline, so such a bar is cut into several.
- `'22`, `'24` and the rest of the "second of N" family are written in one part
  but played in another — the bracket is drawn wherever the engraver had room.
  MusicXML has no way to refer to a passage somewhere else, so the music is
  moved into the part it belongs to and becomes an ordinary second ending there.

Both rewrites live in `internal/timelines`; start at `doc.go` in that package,
which explains the notation before the code.

## Directory structure

`cmd/limepipes-plugin-music-xml`

The plugin executable. It wires the pieces together and serves them over gRPC.

`internal/model`

The MusicXML document model — `score-partwise` and everything under it. These
structs exist to be marshalled, so they mirror the format rather than the music.

`internal/timelines`

Rewrites a tune's time lines into a form MusicXML can express.

`internal/musicmodel/expander`

Expands bagpipe embellishments into the runs of grace notes they are played as.
A doubling or a taorluath is one symbol in the music model but several notes on
the page.

`internal/pluginimplementation`

The plugin API implementation.

## Build

`make build` produces the `limepipes-plugin-music-xml` executable.

The LimePipes application discovers plugins by looking for these executables;
see the [LimePipes README](https://github.com/tomvodi/limepipes) for where to
put them.

## Develop

### Prerequisites

- [golangci-lint](https://golangci-lint.run/) for `make lint`
- [mockery](https://vektra.github.io/mockery/latest/installation) for `make mocks`
- [enumer](https://github.com/dmarkham/enumer) for the `go:generate` directives
  on the model enums

### Make targets

- `make test` runs the tests
- `make lint` runs golangci-lint
- `make check-coverage` runs the tests and checks the thresholds in
  `.testcoverage.yaml`
- `make cover-html` opens the coverage report in a browser

### Test fixtures

The tests work from three kinds of file in `internal/testfiles` and
`internal/timelines/testfiles`:

- `*.bww` — the tune as written, kept for provenance
- `*.yaml` — the music model the bww plugin parses that file into, which is what
  this plugin is actually handed
- `*.musicxml` — the expected export

The `.yaml` fixtures are produced by running the `.bww` through the current
[bww plugin](https://github.com/tomvodi/limepipes-plugin-bww) parser, so they
have to be regenerated whenever that parser's output changes.

To regenerate an expected `.musicxml`, uncomment the `exportToMusicXML` call in
the spec that reads it, run the suite once, then comment it out again. Check the
diff before keeping it — the file becomes the expectation that spec compares
against from then on.
