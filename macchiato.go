package catppuccingo

// macchiato variant.
type macchiato struct{}

var _ Theme = macchiato{}

func (macchiato) Name() string     { return "macchiato" }
func (macchiato) Pink() string     { return "#f5bde6" }
func (macchiato) Mauve() string    { return "#c6a0f6" }
func (macchiato) Red() string      { return "#ed8796" }
func (macchiato) Maroon() string   { return "#ee99a0" }
func (macchiato) Peach() string    { return "#f5a97f" }
func (macchiato) Yellow() string   { return "#eed49f" }
func (macchiato) Green() string    { return "#a6da95" }
func (macchiato) Sky() string      { return "#91d7e3" }
func (macchiato) Blue() string     { return "#8aadf4" }
func (macchiato) Lavender() string { return "#b7bdf8" }
func (macchiato) Text() string     { return "#cad3f5" }
func (macchiato) Overlay0() string { return "#6e738d" }
func (macchiato) Surface2() string { return "#5b6078" }
func (macchiato) Surface0() string { return "#363a4f" }
func (macchiato) Base() string     { return "#24273a" }
