package catppuccingo

import (
	"image/color"
	"testing"
)

func TestVariant(t *testing.T) {
	for s, ss := range map[string]string{
		"moCha":     Mocha.Name(),
		"Mocha":     Mocha.Name(),
		"MOCHA":     Mocha.Name(),
		"mocha":     Mocha.Name(),
		"Frappe":    Frappe.Name(),
		"frappe":    Frappe.Name(),
		"MaccHiaTo": Macchiato.Name(),
		"latte":     Latte.Name(),
	} {
		t.Run(s, func(t *testing.T) {
			if Variant(s).Name() != ss {
				t.Errorf("expected %q, got %q", ss, Variant(s).Name())
			}
		})
	}
	t.Run("invalid", func(t *testing.T) {
		v := Variant("invalid")
		if v != nil {
			t.Errorf("expected nil, got %v", v)
		}
	})
}

func TestColor(t *testing.T) {
	c := Mocha.Base()
	r := c.RGB[0]
	g := c.RGB[1]
	b := c.RGB[2]
	if 30 != r || 30 != g || b != 46 {
		t.Fatalf("expected rgba=%d %d %d, got %d %d %d", 30, 30, 46, r, g, b)
	}
}

func TestStdColor(t *testing.T) {
	ourColor := Mocha.Base()
	var otherColor color.RGBA
	{
		otherColor.R = ourColor.RGB[0]
		otherColor.G = ourColor.RGB[1]
		otherColor.B = ourColor.RGB[2]
		otherColor.A = 0xff
	}

	r, g, b, a := ourColor.RGBA()
	otherR, otherG, otherB, otherA := otherColor.RGBA()

	if r != otherR || g != otherG || b != otherB || a != otherA {
		t.Fatalf(
			"image.color.RGBA gave %d %d %d %d, but we got %d %d %d %d",
			otherR, otherG, otherB, otherA, r, g, b, a,
		)
	}
}

func TestTypeAlias(t *testing.T) {
	var mocha Theme = Mocha
	if mocha.Name() != "mocha" {
		t.Errorf("expected Mocha, got %q", mocha.Name())
	}
}

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
