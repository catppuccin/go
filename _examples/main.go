package main

import (
	"fmt"

	catppuccin "github.com/catppuccin/go"
	"github.com/charmbracelet/lipgloss"
)

func main() {
	fmt.Println()
	for _, flavour := range []catppuccin.Flavour{
		catppuccin.Mocha(),
		catppuccin.Frappe(),
		catppuccin.Macchiato(),
		catppuccin.Latte(),
	} {

		fmt.Println(lipgloss.NewStyle().Bold(true).Render(flavour.Name() + ":"))
		format("rosewater", flavour.Rosewater(), flavour.Base())
		format("flamingo", flavour.Flamingo(), flavour.Base())
		format("pink", flavour.Pink(), flavour.Base())
		format("mauve", flavour.Mauve(), flavour.Base())
		format("red", flavour.Red(), flavour.Base())
		fmt.Println()
		format("maroon", flavour.Maroon(), flavour.Base())
		format("peach", flavour.Peach(), flavour.Base())
		format("yellow", flavour.Yellow(), flavour.Base())
		format("green", flavour.Green(), flavour.Base())
		format("teal", flavour.Teal(), flavour.Base())
		fmt.Println()
		format("sky", flavour.Sky(), flavour.Base())
		format("sapphire", flavour.Sapphire(), flavour.Base())
		format("blue", flavour.Blue(), flavour.Base())
		format("lavender", flavour.Lavender(), flavour.Base())
		format("text", flavour.Text(), flavour.Base())
		fmt.Println()
		format("subtext1", flavour.Subtext1(), flavour.Base())
		format("subtext0", flavour.Subtext0(), flavour.Base())
		format("overlay2", flavour.Overlay2(), flavour.Base())
		format("overlay1", flavour.Overlay1(), flavour.Base())
		format("overlay0", flavour.Overlay0(), flavour.Text())
		fmt.Println()
		format("surface2", flavour.Surface2(), flavour.Text())
		format("surface1", flavour.Surface1(), flavour.Text())
		format("surface0", flavour.Surface0(), flavour.Text())
		format("crust", flavour.Crust(), flavour.Text())
		format("mantle", flavour.Mantle(), flavour.Text())
		fmt.Println()
		format("base", flavour.Base(), flavour.Text())
		fmt.Println()
		fmt.Println()
	}
}

func format(s string, c, txt catppuccin.Color) {
	fmt.Print(lipgloss.NewStyle().
		Background(lipgloss.Color(c.Hex)).
		Foreground(lipgloss.Color(txt.Hex)).
		Align(lipgloss.Center).
		Width(22).
		Render(s))
}
