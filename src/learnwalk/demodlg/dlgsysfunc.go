package demodlg

import (
	// "bufio"
	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
	"github.com/lxn/win"
	//	"strings"
	"log"
)

type MySysFuncWin struct {
	*walk.MainWindow
	hWnd win.HWND
}

func NewSysFuncDlg() {
	mw := &MySysFuncWin{}
	err := MainWindow{
		AssignTo: &mw.MainWindow,
		Title:    "the dialog",
		MinSize:  Size{300, 200},
		Layout:   VBox{},
		Children: []Widget{
			CheckBox{
				Name:    "chkOpen",
				Text:    "Open / Special Enabled",
				Checked: true,
			},
			PushButton{
				Text:      "去掉最小化",
				OnClicked: mw.onRemoveMinimizeBox,
			},
			PushButton{
				Text:      "启用最小化",
				OnClicked: mw.onAddMinimizeBox,
			},
		},
	}.Create()
	if err != nil {
		log.Fatal(err)
	}
	mw.hWnd = mw.Handle()
	mw.Run()
}

func (mw *MySysFuncWin) onRemoveMinimizeBox() {
	mw.removeStyle(win.WS_MINIMIZEBOX)
	//	mw.removeStyle(win.WS_SYSMENU)
	mw.Show()
}
func (mw *MySysFuncWin) onAddMinimizeBox() {
	mw.addStyle(win.WS_MINIMIZEBOX)
	//	mw.removeStyle(win.WS_SYSMENU)
	mw.Show()
}

func (mw *MySysFuncWin) onRemoveMaximize() {
	mw.removeStyle(win.WS_MAXIMIZEBOX)
	mw.Show()
}

func (mw *MySysFuncWin) addStyle(style int32) {
	curstyle := win.GetWindowLong(mw.hWnd, win.GWL_STYLE)
	win.SetWindowLong(mw.hWnd, win.GWL_STYLE, curstyle|style)
}

func (mw *MySysFuncWin) removeStyle(style int32) {
	curstyle := win.GetWindowLong(mw.hWnd, win.GWL_STYLE)
	win.SetWindowLong(mw.hWnd, win.GWL_STYLE, curstyle&(^(style)))
}
