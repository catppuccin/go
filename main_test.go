package catppuccingo

import "testing"

func TestGet(t *testing.T) {
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
			if Get(s).Name() != ss {
				t.Errorf("expected %q, got %q", ss, Get(s).Name())
			}
		})
	}
	t.Run("invalid", func(t *testing.T) {
		v := Get("invalid")
		if v != nil {
			t.Errorf("expected nil, got %v", v)
		}
	})
}

func TestColor(t *testing.T) {
	r, g, b, a := Mocha.Base().RGBA()
	if 30 != r || 30 != g || b != 46 || a != 1 {
		t.Fatalf("expected rgba=%d %d %d %d, got %d %d %d %d", 30, 30, 46, 1, r, g, b, a)
	}
}

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
