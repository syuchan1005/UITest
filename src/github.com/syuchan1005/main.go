package main

import (
	"github.com/andlabs/ui"
)

func main() {
	err := ui.Main(func() {
		tab := ui.NewTab()
		pbar := ui.NewProgressBar()
		slider := ui.NewSlider(0, 100)
		spinbox := ui.NewSpinbox(0, 100)
		combobox := ui.NewCombobox()
		radios := ui.NewRadioButtons()
		controls := [][]ui.Control{
			{ui.NewLabel("TextField"), ui.NewEntry()},
			{ui.NewLabel("Button"), ui.NewButton("Button")},
			{ui.NewLabel("Label"), ui.NewLabel("LabelText")},
			{ui.NewLabel("CheckBox"), ui.NewCheckbox("Checkbox")},
			{ui.NewLabel("Combobox"), combobox},
			{ui.NewLabel("DataTimePicker"), ui.NewDateTimePicker()},
			{ui.NewLabel("RadioButton"), radios},
			{ui.NewLabel("ProgressBar"), pbar},
			{ui.NewLabel("SpinBox"), spinbox},
			{ui.NewLabel("Slider"), slider},
		}
		slider.OnChanged(func(*ui.Slider) {
			spinbox.SetValue(slider.Value())
			pbar.SetValue(slider.Value())
		})
		spinbox.OnChanged(func(*ui.Spinbox) {
			slider.SetValue(spinbox.Value())
			pbar.SetValue(spinbox.Value())
		})
		combobox.Append("Test1")
		combobox.Append("Test2")
		combobox.Append("Test3")
		radios.Append("Test1")
		radios.Append("Test2")
		radios.Append("Test3")
		tab.Append("main", createVerticalBox(controls))
		tab.Append("sample", ui.NewLabel("test"))
		window := ui.NewWindow("ComponentText", 800, 600, false)
		window.SetChild(tab)
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
		for k := 0; k < len(controls[i]); k++ {
			h.Append(controls[i][k], true);
		}
		box.Append(h, false)
	}
	return box
}
