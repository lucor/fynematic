package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("Fynematic example")
	w.Resize(fyne.NewSize(200, 200))
	w.SetContent(container.NewBorder(
		nil,
		container.NewGridWithColumns(2,
			widget.NewButton("Dark", func() {
				a.Settings().SetTheme(theme.DarkTheme())
			}),
			widget.NewButton("Light", func() {
				a.Settings().SetTheme(theme.LightTheme())
			}),
		),
		nil,
		makeAppTabs(),
	))

	w.ShowAndRun()
}

func makeAppTabs() *container.AppTabs {
	return container.NewAppTabs(
		container.NewTabItem("Outlined", widget.NewIcon(SentimentSatisfiedOutlinedIconThemed)),
		container.NewTabItem("Filled", widget.NewIcon(SentimentSatisfiedFilledIconThemed)),
		container.NewTabItem("Round", widget.NewIcon(SentimentSatisfiedRoundIconThemed)),
		container.NewTabItem("Sharp", widget.NewIcon(SentimentSatisfiedSharpIconThemed)),
		container.NewTabItem("Twotone", widget.NewIcon(SentimentSatisfiedTwotoneIconThemed)),
	)
}
