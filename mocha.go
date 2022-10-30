package catppuccingo

type mocha struct{}

var _ Theme = mocha{}

func (mocha) Name() string     { return "mocha" }
func (mocha) Pink() string     { return "#f5c2e7" }
func (mocha) Mauve() string    { return "#cba6f7" }
func (mocha) Red() string      { return "#f38ba8" }
func (mocha) Maroon() string   { return "#eba0ac" }
func (mocha) Peach() string    { return "#fab387" }
func (mocha) Yellow() string   { return "#f9e2af" }
func (mocha) Green() string    { return "#a6e3a1" }
func (mocha) Sky() string      { return "#89dceb" }
func (mocha) Blue() string     { return "#89b4fa" }
func (mocha) Lavender() string { return "#b4befe" }
func (mocha) Text() string     { return "#cdd6f4" }
func (mocha) Overlay0() string { return "#6c7086" }
func (mocha) Surface2() string { return "#585b70" }
func (mocha) Surface0() string { return "#313244" }
func (mocha) Base() string     { return "#1e1e2e" }
