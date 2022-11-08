package catppuccingo

import (
	"testing"
)

func TestFlavours(t *testing.T) {
	for _, flavour := range []Flavour{
		Mocha,
		Frappe,
		Macchiato,
		Latte,
	} {
		t.Run(flavour.Name(), func(t *testing.T) {
			requireNotEmpty(t, flavour.Rosewater())
			requireNotEmpty(t, flavour.Flamingo())
			requireNotEmpty(t, flavour.Pink())
			requireNotEmpty(t, flavour.Mauve())
			requireNotEmpty(t, flavour.Red())
			requireNotEmpty(t, flavour.Maroon())
			requireNotEmpty(t, flavour.Peach())
			requireNotEmpty(t, flavour.Yellow())
			requireNotEmpty(t, flavour.Green())
			requireNotEmpty(t, flavour.Teal())
			requireNotEmpty(t, flavour.Sky())
			requireNotEmpty(t, flavour.Sapphire())
			requireNotEmpty(t, flavour.Blue())
			requireNotEmpty(t, flavour.Lavender())
			requireNotEmpty(t, flavour.Text())
			requireNotEmpty(t, flavour.Subtext1())
			requireNotEmpty(t, flavour.Subtext0())
			requireNotEmpty(t, flavour.Overlay2())
			requireNotEmpty(t, flavour.Overlay1())
			requireNotEmpty(t, flavour.Overlay0())
			requireNotEmpty(t, flavour.Surface2())
			requireNotEmpty(t, flavour.Surface1())
			requireNotEmpty(t, flavour.Surface0())
			requireNotEmpty(t, flavour.Crust())
			requireNotEmpty(t, flavour.Mantle())
			requireNotEmpty(t, flavour.Base())
		})
	}
}

func requireNotEmpty(tb testing.TB, s Color) {
	tb.Helper()
	if s.Hex == "" {
		tb.Fatalf("should be empty, was %q", s.Hex)
	}
}
