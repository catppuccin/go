package catppuccingo

// Theme is an interface implemented by all Catppuccin variations.
type Theme interface {
	Pink() string
	Mauve() string
	Red() string
	Maroon() string
	Peach() string
	Yellow() string
	Green() string
	Sky() string
	Blue() string
	Lavender() string
	Text() string
	Overlay0() string
	Surface2() string
	Surface0() string
	Base() string
	Name() string
}

// Mocha variant.
func Mocha() Theme { return mocha{} }

// Frappe variant.
func Frappe() Theme { return frappe{} }

// Macchiato variant.
func Macchiato() Theme { return macchiato{} }

// Latte variant.
func Latte() Theme { return latte{} }
