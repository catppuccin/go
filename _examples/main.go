package main

import (
	"fmt"

	catppuccin "github.com/catppuccin/go"
	"github.com/charmbracelet/lipgloss"
)

func main() {
	fmt.Println()
	for _, theme := range []catppuccin.Theme{
		catppuccin.Mocha(),
		catppuccin.Frappe(),
		catppuccin.Macchiato(),
		catppuccin.Latte(),
	} {

		fmt.Println(lipgloss.NewStyle().Bold(true).Render(theme.Name() + ":"))
		format("rosewater", theme.Rosewater(), theme.Base())
		format("flamingo", theme.Flamingo(), theme.Base())
		format("pink", theme.Pink(), theme.Base())
		format("mauve", theme.Mauve(), theme.Base())
		format("red", theme.Red(), theme.Base())
		fmt.Println()
		format("maroon", theme.Maroon(), theme.Base())
		format("peach", theme.Peach(), theme.Base())
		format("yellow", theme.Yellow(), theme.Base())
		format("green", theme.Green(), theme.Base())
		format("teal", theme.Teal(), theme.Base())
		fmt.Println()
		format("sky", theme.Sky(), theme.Base())
		format("sapphire", theme.Sapphire(), theme.Base())
		format("blue", theme.Blue(), theme.Base())
		format("lavender", theme.Lavender(), theme.Base())
		format("text", theme.Text(), theme.Base())
		fmt.Println()
		format("subtext1", theme.Subtext1(), theme.Base())
		format("subtext0", theme.Subtext0(), theme.Base())
		format("overlay2", theme.Overlay2(), theme.Base())
		format("overlay1", theme.Overlay1(), theme.Base())
		format("overlay0", theme.Overlay0(), theme.Text())
		fmt.Println()
		format("surface2", theme.Surface2(), theme.Text())
		format("surface1", theme.Surface1(), theme.Text())
		format("surface0", theme.Surface0(), theme.Text())
		format("crust", theme.Crust(), theme.Text())
		format("mantle", theme.Mantle(), theme.Text())
		fmt.Println()
		format("base", theme.Base(), theme.Text())
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
