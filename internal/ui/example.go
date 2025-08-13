package ui

import (
	"strings"
	"punkdoku/internal/theme"
)

// ExampleRenderSample returns a sample rendering string for docs/tests.
func ExampleRenderSample() string {
	th := theme.Dark()
	styles := BuildStyles(th)
	var b strings.Builder
	b.WriteString(styles.CellSelected.Render("5"))
	b.WriteString(" ")
	b.WriteString(styles.CellDuplicate.Render("5"))
	b.WriteString(" ")
	b.WriteString(styles.CellConflict.Render("3"))
	b.WriteString("\n")
	return styles.App.Render(b.String())
}
