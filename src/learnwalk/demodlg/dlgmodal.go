package demodlg

import (
	// "bufio"
	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
	//	"github.com/lxn/win"
	//	"strings"
	"log"
	"time"
)

type Animal struct {
	Name          string
	ArrivalDate   time.Time
	SpeciesId     int
	Speed         int
	Sex           Sex
	Weight        float64
	PreferredFood string
	Domesticated  bool
	Remarks       string
	Patience      time.Duration
}

func (a *Animal) PatienceField() *DurationField {
	return &DurationField{&a.Patience}
}

type DurationField struct {
	p *time.Duration
}

func (*DurationField) CanSet() bool {
	return true
}
func (f *DurationField) Get() interface{} {
	return f.p.String()
}
func (f *DurationField) Set(v interface{}) error {
	x, err := time.ParseDuration(v.(string))
	if err == nil {
		*f.p = x
	}
	return err
}
func (f *DurationField) Zero() interface{} {
	return ""
}

type Sex byte

const (
	SexMale Sex = 1 + iota
	SexFemale
	SexHermaphrodite
)

type Species struct {
	Id   int
	Name string
}

func KnownSpecies() []*Species {
	return []*Species{
		{1, "Dog"},
		{2, "Cat"},
		{3, "Bird"},
		{4, "Fish"},
		{5, "Elephant"},
	}
}

func RunModalDialog(owner walk.Form, animal *Animal) (int, error) {
	var dlg *walk.Dialog
	var nameEdit *walk.LineEdit
	var db *walk.DataBinder

	return Dialog{
		AssignTo: &dlg,
		Title:    Bind("'Animal Details' + (dog.Name == '' ? '' : ' - ' + dog.Name)"),
		MinSize:  Size{300, 500},
		Layout:   VBox{},
		DataBinder: DataBinder{
			AssignTo:       &db,
			Name:           "dog",
			DataSource:     animal,
			ErrorPresenter: ToolTipErrorPresenter{},
		},
		Children: []Widget{
			Composite{
				Layout: Grid{Columns: 2},
				Children: []Widget{
					Label{
						Text: "Name:",
					},
					LineEdit{
						AssignTo: &nameEdit,
						Text:     Bind("Name"),
						//	Text: animal.Name,
					},
					Label{
						Text: "Arrival Date:",
					},
					DateEdit{
						Date: Bind("ArrivalDate"),
					},
					Label{
						Text: "Species:",
					},
					ComboBox{
						//	Value:         Bind("SpeciesId", SelRequired{}),
						Value:         Bind("SpeciesId", SelRequired{}),
						BindingMember: "Id",
						DisplayMember: "Name",
						Model:         KnownSpecies(),
					},
					Label{
						Text: "Speed:",
					},
					Slider{
						Value: Bind("Speed"),
					},
					RadioButtonGroupBox{
						ColumnSpan: 2,
						Title:      "Sex",
						Layout:     HBox{},
						DataMember: "Sex",
						Buttons: []RadioButton{
							{Text: "Male", Value: SexMale},
							{Text: "Female", Value: SexFemale},
							{Text: "Hermaphrodite", Value: SexHermaphrodite},
						},
					},
					Label{
						Text: "Weight:",
					},
					NumberEdit{
						Value:    Bind("Weight", Range{0.01, 9999.99}),
						Suffix:   " kg",
						Decimals: 2,
					},
					Label{
						Text: "Preferred Food:",
					},
					ComboBox{
						Editable: false,
						Value:    Bind("PreferredFood"),
						Model:    []string{"Fruit", "Grass", "Fish", "Meat"},
					},
					Label{
						Text: "Domesticated:",
					},
					CheckBox{
						Checked: Bind("Domesticated"),
					},
					VSpacer{
						ColumnSpan: 2,
						Size:       8,
					},
					Label{
						ColumnSpan: 2,
						Text:       "Remarks:",
					},
					/*
						GroupBox{
							//	ColumnSpan: 2,
							Title:  "Color1",
							Layout: Grid{Columns: 2},
							Children: []Widget{
								Label{Text: "Red:"},
								Slider{Name: "c1RedSld", Tracking: true, MaxValue: 255, Value: 95},
								Label{Text: "Green:"},
								Slider{Name: "c1GreenSld", Tracking: true, MaxValue: 255, Value: 191},
								Label{Text: "Blue:"},
								Slider{Name: "c1BlueSld", Tracking: true, MaxValue: 255, Value: 255},
							},
						},
					*/
					TextEdit{
						ColumnSpan: 2,
						MinSize:    Size{100, 50},
						Text:       Bind("Remarks"),
					},
					Label{
						Text: "Patience:",
					},
					LineEdit{
						Text: Bind("PatienceField"),
					},
				},
			},

			Composite{
				Layout: HBox{},
				Children: []Widget{
					HSpacer{},
					PushButton{
						//	AssignTo: &acceptPB,
						Text: "OK",
						OnClicked: func() {
							if err := db.Submit(); err != nil {
								log.Print(err)
								return
							}

							dlg.Accept()
						},
					},
					PushButton{
						//	AssignTo:  &cancelPB,
						Text:      "Cancel",
						OnClicked: func() { dlg.Cancel() },
					},
				},
			},
		},
	}.Run(owner)

}
