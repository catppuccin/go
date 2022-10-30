package catppuccingo

// frappe variant.
type frappe struct{}

var _ Theme = frappe{}

func (frappe) Name() string     { return "frappe" }
func (frappe) Pink() string     { return "#f4b8e4" }
func (frappe) Mauve() string    { return "#ca9ee6" }
func (frappe) Red() string      { return "#e78284" }
func (frappe) Maroon() string   { return "#ea999c" }
func (frappe) Peach() string    { return "#ef9f76" }
func (frappe) Yellow() string   { return "#e5c890" }
func (frappe) Green() string    { return "#a6d189" }
func (frappe) Sky() string      { return "#99d1db" }
func (frappe) Blue() string     { return "#8caaee" }
func (frappe) Lavender() string { return "#babbf1" }
func (frappe) Text() string     { return "#c6d0f5" }
func (frappe) Overlay0() string { return "#737994" }
func (frappe) Surface2() string { return "#626880" }
func (frappe) Surface0() string { return "#414559" }
func (frappe) Base() string     { return "#303446" }
