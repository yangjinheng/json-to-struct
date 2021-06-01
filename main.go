package main

import (
	"log"
	"strings"

	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
)

var sample = `{
  "name": "Json To Go Struct",
  "author": "yangjinheng",
  "feat": [
    "Name",
    "Suffix",
    "Omitempty",
    "Not Exported",
    "Example"
  ],
  "reference": [
    {
      "name": "json2struct",
      "link": "https://github.com/yudppp/json2struct"
    },
    {
      "name": "walk",
      "link": "https://github.com/lxn/walk"
    }
  ],
  "version": 1.0
}
`

func onchange(opt Options, input, output *walk.TextEdit) {
	in := input.Text()
	parsed, err := Parse(strings.NewReader(in), opt)
	if err != nil {
		output.SetText("invalid json string")
		return
	}
	output.SetText(strings.Replace(parsed, "\n", "\r\n", -1))
}

func main() {
	var (
		mainWindow *walk.MainWindow
		db         *walk.DataBinder
		input      *walk.TextEdit
		output     *walk.TextEdit
	)

	var options = new(Options)

	err := MainWindow{
		AssignTo: &mainWindow,
		Title:    "JSON To Go Struct",
		Size:     Size{Width: 1000, Height: 600},
		MinSize:  Size{Width: 1000, Height: 600},
		Layout:   VBox{Margins: Margins{Left: 5, Right: 0, Top: 10, Bottom: 10}},
		DataBinder: DataBinder{
			AssignTo:       &db,
			Name:           "options",
			DataSource:     options,
			ErrorPresenter: ToolTipErrorPresenter{},
		},
		Children: []Widget{
			Composite{
				Layout:  Grid{Rows: 1},
				MaxSize: Size{Width: 600},
				Children: []Widget{
					Label{
						Text: "Name:",
					},
					LineEdit{
						MaxSize: Size{Width: 100},
						MinSize: Size{Width: 100},
						Text:    Bind("Name"),
						OnKeyUp: func(key walk.Key) {
							db.Submit()
							onchange(*options, input, output)
						},
					},
					Label{
						Text: "Suffix:",
					},
					LineEdit{
						MaxSize: Size{Width: 100},
						MinSize: Size{Width: 100},
						Text:    Bind("Suffix"),
						OnKeyUp: func(key walk.Key) {
							db.Submit()
							onchange(*options, input, output)
						},
					},
					CheckBox{
						Checked: Bind("UseShortStruct"),
						OnCheckedChanged: func() {
							db.Submit()
							onchange(*options, input, output)
						},
						AlwaysConsumeSpace: true,
						Text:               "Short Name",
					},
					HSpacer{
						GreedyLocallyOnly: true,
					},
					CheckBox{
						Checked: Bind("UseOmitempty"),
						OnCheckedChanged: func() {
							db.Submit()
							onchange(*options, input, output)
						},
						AlwaysConsumeSpace: true,
						Text:               "Omitempty",
					},
					CheckBox{
						Checked: Bind("UseExample"),
						OnCheckedChanged: func() {
							db.Submit()
							onchange(*options, input, output)
						},
						AlwaysConsumeSpace: true,
						Text:               "Example",
					},
					CheckBox{
						Checked: Bind("UseLocal"),
						OnCheckedChanged: func() {
							db.Submit()
							onchange(*options, input, output)
						},
						AlwaysConsumeSpace: true,
						Text:               "Not Exported",
					},
				},
			},
			Composite{
				Layout: HBox{MarginsZero: true, SpacingZero: true},
				Children: []Widget{
					TextEdit{
						AssignTo: &input,
						Text:     strings.Replace(sample, "\n", "\r\n", -1),
						Font: Font{
							Family:    "Consolas",
							PointSize: 10,
						},
						OnTextChanged: func() {
							db.Submit()
							onchange(*options, input, output)
						},
						VScroll: true,
					},
					TextEdit{
						AssignTo: &output,
						Font: Font{
							Family:    "Consolas",
							PointSize: 10,
						},
						VScroll: true,
					},
				},
			},
		},
	}.Create()
	if err != nil {
		log.Fatal(err)
	}
	mainWindow.Run()
}
