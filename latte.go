package catppuccingo

// latte variant.
type latte struct{}

var _ Theme = latte{}

func (latte) Name() string     { return "latte" }
func (latte) Pink() string     { return "#ea76cb" }
func (latte) Mauve() string    { return "#8839ef" }
func (latte) Red() string      { return "#d20f39" }
func (latte) Maroon() string   { return "#e64553" }
func (latte) Peach() string    { return "#fe640b" }
func (latte) Yellow() string   { return "#df8e1d" }
func (latte) Green() string    { return "#40a02b" }
func (latte) Sky() string      { return "#04a5e5" }
func (latte) Blue() string     { return "#1e66f5" }
func (latte) Lavender() string { return "#7287fd" }
func (latte) Text() string     { return "#4c4f69" }
func (latte) Overlay0() string { return "#9ca0b0" }
func (latte) Surface2() string { return "#acb0be" }
func (latte) Surface0() string { return "#ccd0da" }
func (latte) Base() string     { return "#eff1f5" }
