package main

import (
	"github.com/andlabs/ui"
)

func main() {
	err := ui.Main(func() {
		pbar := ui.NewProgressBar()
		slider := ui.NewSlider(0, 100)
		spinbox := ui.NewSpinbox(0, 100)
		controls := [][]ui.Control{
			{ui.NewLabel("TextField"), ui.NewEntry()},
			{ui.NewLabel("Button"), ui.NewButton("Button")},
			{ui.NewLabel("Label"), ui.NewLabel("LabelText")},
			{ui.NewLabel("CheckBox"), ui.NewCheckbox("Checkbox")},
			{ui.NewLabel("Combobox"), ui.NewCombobox()},
			{ui.NewLabel("DataTimePicker"), ui.NewDateTimePicker()},
			{ui.NewLabel("ProgressBar"), pbar},
			{ui.NewLabel("RadioButton"), ui.NewRadioButtons()},
			{ui.NewLabel("HorizontalSeparator"), ui.NewHorizontalSeparator()},
			{ui.NewLabel("Slider"), slider},
			{ui.NewLabel("SpinBox"), spinbox},
			{ui.NewLabel("Tab"), ui.NewTab()},
		}
		slider.OnChanged(func(*ui.Slider) {
			spinbox.SetValue(slider.Value())
			pbar.SetValue(slider.Value())
		})
		spinbox.OnChanged(func(*ui.Spinbox) {
			slider.SetValue(spinbox.Value())
			pbar.SetValue(spinbox.Value())
		})
		window := ui.NewWindow("ComponentText", 200, 100, false)
		window.SetChild(createVerticalBox(controls))
		window.OnClosing(func(*ui.Window) bool {
			ui.Quit()
			return true
		})
		window.Show()
	})
	if err != nil {
		panic(err)
	}
}

func createVerticalBox(controls [][]ui.Control) ui.Control {
	box := ui.NewVerticalBox();
	for i := 0; i < len(controls); i++ {
		h := ui.NewHorizontalBox();
		h.Append(controls[i][0], false);
		h.Append(ui.NewHorizontalSeparator(), false)
		h.Append(controls[i][1], false);
		box.Append(h, true)
	}
	return box
}
