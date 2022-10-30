package catppuccingo

import "testing"

func TestThemes(t *testing.T) {
	for _, theme := range []Theme{
		Mocha(),
		Frappe(),
		Macchiato(),
		Latte(),
	} {
		t.Run(theme.Name(), func(t *testing.T) {
			requireNotEmpty(t, theme.Pink())
			requireNotEmpty(t, theme.Mauve())
			requireNotEmpty(t, theme.Red())
			requireNotEmpty(t, theme.Maroon())
			requireNotEmpty(t, theme.Peach())
			requireNotEmpty(t, theme.Yellow())
			requireNotEmpty(t, theme.Green())
			requireNotEmpty(t, theme.Sky())
			requireNotEmpty(t, theme.Blue())
			requireNotEmpty(t, theme.Lavender())
			requireNotEmpty(t, theme.Text())
			requireNotEmpty(t, theme.Overlay0())
			requireNotEmpty(t, theme.Surface2())
			requireNotEmpty(t, theme.Surface0())
			requireNotEmpty(t, theme.Base())
		})
	}
}

func requireNotEmpty(tb testing.TB, s string) {
	tb.Helper()
	if s == "" {
		tb.Fatalf("should be empty, was %q", s)
	}
}
