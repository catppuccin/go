package catppuccingo

// Theme is an interface implemented by all Catppuccin variations.
type Theme interface {
	Rosewater() Color
	Flamingo() Color
	Pink() Color
	Mauve() Color
	Red() Color
	Maroon() Color
	Peach() Color
	Yellow() Color
	Green() Color
	Teal() Color
	Sky() Color
	Sapphire() Color
	Blue() Color
	Lavender() Color
	Text() Color
	Subtext1() Color
	Subtext0() Color
	Overlay2() Color
	Overlay1() Color
	Overlay0() Color
	Surface2() Color
	Surface1() Color
	Surface0() Color
	Crust() Color
	Mantle() Color
	Base() Color
	Name() string
}

type Color struct {
	Hex string
	RGB [3]int
	HSL [3]float32
}

// Mocha variant.
func Mocha() Theme { return mocha{} }

// Frappe variant.
func Frappe() Theme { return frappe{} }

// Macchiato variant.
func Macchiato() Theme { return macchiato{} }

// Latte variant.
func Latte() Theme { return latte{} }
