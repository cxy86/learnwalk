package main

import (
	// "bufio"
	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
	"github.com/lxn/win"
	"log"
	// "strings"
	"learnwalk/demodlg"
	"time"
)

type MyMainWindow struct {
	*walk.MainWindow
	hWnd     win.HWND
	chkClose *walk.CheckBox
	ni       *walk.NotifyIcon
	dlg      *walk.Dialog
	animal   *demodlg.Animal
}

func main() {
	/*
		walk.AppendToWalkInit(func() {
				walk.FocusEffect, _ = walk.NewBorderGlowEffect(walk.RGB(255, 0, 0))
				walk.InteractionEffect, _ = walk.NewDropShadowEffect(walk.RGB(63, 63, 63))
				walk.ValidationErrorEffect, _ = walk.NewBorderGlowEffect(walk.RGB(0, 255, 0))
			})
	*/
	mw := &MyMainWindow{}
	var editMenu, recentMenu *walk.Menu
	var showAbout, openAction *walk.Action
	var toggleSpecialModePB *walk.PushButton
	var isSpecialMode = walk.NewMutableCondition()
	MustRegisterCondition("bMode", isSpecialMode)
	isSpecialMode.SetSatisfied(true)
	err := MainWindow{
		AssignTo: &mw.MainWindow,
		Title:    "walk学习",
		MinSize:  Size{300, 200},
		Layout:   VBox{},
		OnSizeChanged: func() {
			if win.IsIconic(mw.Handle()) {
				mw.Hide()
				mw.ni.SetVisible(true)
			}
		},
		MenuItems: []MenuItem{
			Menu{
				Text: "&File",
				Items: []MenuItem{
					Action{
						AssignTo: &openAction,
						Text:     "&Open",
						Image:    "img/open.png",
						Enabled:  Bind("chkOpen.Checked"),
						//	Visible:     Bind("!openHiddenCB.Checked"),
						Shortcut:    Shortcut{walk.ModControl, walk.KeyO},
						OnTriggered: mw.onOpenFile,
					},
					Separator{},
					Menu{
						AssignTo: &recentMenu,
						Text:     "&Recent",
					},
					Separator{},
					Action{
						Text:        "E&xit",
						OnTriggered: func() { mw.Close() },
					},
				},
			},
			Menu{
				Text:     "&Edit",
				AssignTo: &editMenu,
				Items: []MenuItem{
					Action{
						//		AssignTo:    &openAction,
						Text:        "画画",
						OnTriggered: mw.onDrawing,
					},
					Action{
						Text:        "打开模态对话框",
						OnTriggered: mw.onOpenModalDlg,
					},
					Action{
						//		AssignTo:    &openAction,
						Text:        "系统设置",
						OnTriggered: mw.onSysFunc,
					},
					Action{
						Text:        "拖放图片文件到对话框中",
						OnTriggered: mw.onDropFile,
					},
					Action{
						Text:        "自定义windows控件",
						OnTriggered: mw.onExWidget,
					},
				},
			},
			Menu{
				Text: "&View",
				Items: []MenuItem{
					Action{
						//		AssignTo:    &openAction,
						Text:    "Open/Special Enable",
						Image:   "img/open.png",
						Checked: Bind("chkOpen.Visible"),
						//	Enabled:     Bind("enabledCB.Checked"),
						//	Visible:     Bind("!openHiddenCB.Checked"),
						//	Shortcut: Shortcut{walk.ModControl, walk.KeyO},
						//	OnTriggered: mw.openAction_Triggered,
					},
					Action{
						//		AssignTo:    &openAction,
						Text: "Open Hide",
						//		Image:   "img/open.png",
						Checked: Bind("chkHidden.Visible"),
						//	Enabled:     Bind("enabledCB.Checked"),
						//	Visible:     Bind("!openHiddenCB.Checked"),
						//	Shortcut: Shortcut{walk.ModControl, walk.KeyO},
						//	OnTriggered: mw.openAction_Triggered,
					},
				},
			},
			Menu{
				Text: "&Help",
				Items: []MenuItem{
					Action{
						AssignTo:    &showAbout,
						Text:        "About",
						OnTriggered: mw.showAboutBoxAction_Triggered,
					},
				},
			},
		},
		Children: []Widget{
			CheckBox{
				Name:    "chkOpen",
				Text:    "Open / Special Enabled",
				Checked: true,
				Accessibility: Accessibility{
					Help: "Enables Open and Special",
				},
			},
			CheckBox{
				Name:    "chkHidden",
				Text:    "Open Hidden",
				Checked: true,
			},
			CheckBox{
				AssignTo:         &mw.chkClose,
				Name:             "chkClose",
				Text:             "关闭按钮",
				Checked:          true,
				OnCheckedChanged: mw.onChkCloseChanged,
			},
			PushButton{
				AssignTo: &toggleSpecialModePB,
				Text:     "Disable Special Mode",
				OnClicked: func() {
					isSpecialMode.SetSatisfied(!isSpecialMode.Satisfied())
					var b = isSpecialMode.Satisfied()
					if isSpecialMode.Satisfied() {
						toggleSpecialModePB.SetText("Disable Special Mode")
						mw.SetFullscreen(b)
					} else {
						toggleSpecialModePB.SetText("Enable Special Mode")
						mw.SetFullscreen(b)
					}
				},
				Accessibility: Accessibility{
					Help: "Toggles special mode",
				},
			},
		},
		ContextMenuItems: []MenuItem{
			ActionRef{&showAbout},
			Action{
				//		AssignTo:    &openAction,
				Text:    "Open/Special Enable",
				Image:   "img/open.png",
				Checked: Bind("chkOpen.Visible"),
				//	Enabled:     Bind("enabledCB.Checked"),
				//	Visible:     Bind("!openHiddenCB.Checked"),
				//	Shortcut: Shortcut{walk.ModControl, walk.KeyO},
				//	OnTriggered: mw.openAction_Triggered,
			},
			Action{
				//		AssignTo:    &openAction,
				Text:  "&Cut",
				Image: "img/open.png",
				//	Enabled:     Bind("enabledCB.Checked"),
				//	Visible:     Bind("!openHiddenCB.Checked"),
				//	Shortcut: Shortcut{walk.ModControl, walk.KeyO},
				//	OnTriggered: mw.openAction_Triggered,
			},
		},
		ToolBar: ToolBar{
			ButtonStyle: ToolBarButtonImageBeforeText,
			Items: []MenuItem{
				ActionRef{&openAction},
				Menu{
					Text:  "New A",
					Image: "img/document-new.png",
					Items: []MenuItem{
						Action{
							Text:        "A",
							OnTriggered: mw.newAction_Triggered,
						},
						Action{
							Text:        "B",
							OnTriggered: mw.newAction_Triggered,
						},
						Action{
							Text:        "C",
							OnTriggered: mw.newAction_Triggered,
						},
					},

					OnTriggered: mw.newAction_Triggered,
				},
				Separator{},
				Menu{
					Text:  "View",
					Image: "img/document-properties.png",
					Items: []MenuItem{
						Action{
							Text:        "X",
							OnTriggered: mw.changeViewAction_Triggered,
						},
						Action{
							Text:        "Y",
							OnTriggered: mw.changeViewAction_Triggered,
						},
						Action{
							Text:        "Z",
							OnTriggered: mw.changeViewAction_Triggered,
						},
					},
				},
				Action{
					Text:        "Special",
					Image:       "img/system-shutdown.png",
					Enabled:     Bind("bMode"),
					OnTriggered: mw.alert("So Special"),
				},
			},
		},
	}.Create()
	if err != nil {
		log.Fatal(err)
	}
	addRecentFileAction := func(texts ...string) {
		for _, text := range texts {
			a := walk.NewAction()
			a.SetText(text)
			a.Triggered().Attach(mw.onOpenFile)
			recentMenu.Actions().Add(a)
		}
	}
	addRecentFileAction("hi", "house", "sea")

	mw.hWnd = mw.Handle()
	mw.AddNotifyIcon()
	mw.Run()
}

func (mw *MyMainWindow) showAboutBoxAction_Triggered() {
	walk.MsgBox(mw, "About", "I will beat you", walk.MsgBoxIconExclamation)
}

func (mw *MyMainWindow) onOpenFile() {
	walk.MsgBox(mw, "Tip", "I am open-minded", walk.MsgBoxIconExclamation)
}
func (mw *MyMainWindow) newAction_Triggered() {
	walk.MsgBox(mw, "Tip", "newAction_Triggered", walk.MsgBoxIconExclamation)
}
func (mw *MyMainWindow) changeViewAction_Triggered() {
	walk.MsgBox(mw, "Tip", "changeViewAction_Triggered", walk.MsgBoxIconExclamation)
}
func (mw *MyMainWindow) onChkCloseChanged() {
	b := mw.chkClose.Checked()
	if b {
		win.GetSystemMenu(mw.hWnd, true)
		mw.Show()
		return
	}
	hMenu := win.GetSystemMenu(mw.hWnd, false)
	win.RemoveMenu(hMenu, win.SC_CLOSE, win.MF_BYCOMMAND)
	mw.Show()
}

func (mw *MyMainWindow) alert(txt string) func() {
	return func() {
		walk.MsgBox(mw, "Tip", txt, walk.MsgBoxIconExclamation)
	}
}

func (mw *MyMainWindow) AddNotifyIcon() {
	var err error
	mw.ni, err = walk.NewNotifyIcon(mw)
	if err != nil {
		log.Fatal(err)
	}
	icon, err := walk.Resources.Image("img/stop.ico")
	if err != nil {
		log.Fatal(err)
	}
	mw.SetIcon(icon)
	mw.ni.SetIcon(icon)
	mw.ni.SetVisible(true)
	mw.ni.MouseDown().Attach(func(x, y int, button walk.MouseButton) {
		if button == walk.LeftButton {
			mw.Show()
			win.ShowWindow(mw.hWnd, win.SW_RESTORE)
		}
	})
}

type MyDialog struct {
	*walk.Dialog
}

func (mw *MyMainWindow) onDrawing() {
	demodlg.RunDrawingDlg(mw)
}

func (mw *MyMainWindow) onSysFunc() {
	demodlg.NewSysFuncDlg()
}

func (mw *MyMainWindow) onOpenModalDlg() {
	if nil == mw.animal {
		mw.animal = new(demodlg.Animal)
		mw.animal.ArrivalDate = time.Now()
		mw.animal.SpeciesId = 2
		//	walk.MsgBox(mw, "info", "new Animal", walk.MsgBoxOK)
	}
	demodlg.RunModalDialog(mw, mw.animal)
}

func (mw *MyMainWindow) onDropFile() {
	demodlg.RunDropFileDlg(mw)
}
func (mw *MyMainWindow) onExWidget() {
	demodlg.RunExWidgetDlg(mw)
}
