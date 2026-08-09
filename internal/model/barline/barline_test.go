package barline

import (
	"testing"

	. "github.com/onsi/gomega"
	"github.com/tomvodi/limepipes-plugin-api/musicmodel/v1/barline"
	"github.com/tomvodi/limepipes-plugin-music-xml/internal/utils"
)

func Test_convertBarlineType(t *testing.T) {
	utils.SetupConsoleLogger()
	g := NewGomegaWithT(t)
	type fields struct {
		barlineType barline.Type
		want        Style
	}
	tests := []struct {
		name    string
		prepare func(f *fields)
	}{
		{
			name: "Regular",
			prepare: func(f *fields) {
				f.barlineType = barline.Type_Regular
				f.want = Regular
			},
		},
		{
			name: "Heavy",
			prepare: func(f *fields) {
				f.barlineType = barline.Type_Heavy
				f.want = Heavy
			},
		},
		{
			name: "HeavyHeavy",
			prepare: func(f *fields) {
				f.barlineType = barline.Type_HeavyHeavy
				f.want = HeavyHeavy
			},
		},
		{
			name: "LightHeavy",
			prepare: func(f *fields) {
				f.barlineType = barline.Type_LightHeavy
				f.want = LightHeavy
			},
		},
		{
			name: "HeavyLight",
			prepare: func(f *fields) {
				f.barlineType = barline.Type_HeavyLight
				f.want = HeavyLight
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(_ *testing.T) {
			f := &fields{}

			if tt.prepare != nil {
				tt.prepare(f)
			}

			got := convertBarlineType(f.barlineType)
			g.Expect(got).To(Equal(f.want))
		})
	}
}
