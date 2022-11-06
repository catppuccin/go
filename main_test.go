package catppuccingo

import (
	"testing"
)

func TestThemes(t *testing.T) {
	for _, theme := range []Theme{
		Mocha,
		Frappe,
		Macchiato,
		Latte,
	} {
		t.Run(theme.Name(), func(t *testing.T) {
			requireNotEmpty(t, theme.Rosewater())
			requireNotEmpty(t, theme.Flamingo())
			requireNotEmpty(t, theme.Pink())
			requireNotEmpty(t, theme.Mauve())
			requireNotEmpty(t, theme.Red())
			requireNotEmpty(t, theme.Maroon())
			requireNotEmpty(t, theme.Peach())
			requireNotEmpty(t, theme.Yellow())
			requireNotEmpty(t, theme.Green())
			requireNotEmpty(t, theme.Teal())
			requireNotEmpty(t, theme.Sky())
			requireNotEmpty(t, theme.Sapphire())
			requireNotEmpty(t, theme.Blue())
			requireNotEmpty(t, theme.Lavender())
			requireNotEmpty(t, theme.Text())
			requireNotEmpty(t, theme.Subtext1())
			requireNotEmpty(t, theme.Subtext0())
			requireNotEmpty(t, theme.Overlay2())
			requireNotEmpty(t, theme.Overlay1())
			requireNotEmpty(t, theme.Overlay0())
			requireNotEmpty(t, theme.Surface2())
			requireNotEmpty(t, theme.Surface1())
			requireNotEmpty(t, theme.Surface0())
			requireNotEmpty(t, theme.Crust())
			requireNotEmpty(t, theme.Mantle())
			requireNotEmpty(t, theme.Base())
		})
	}
}

func requireNotEmpty(tb testing.TB, s Color) {
	tb.Helper()
	if s.Hex == "" {
		tb.Fatalf("should be empty, was %q", s.Hex)
	}
}
