package demodlg

import (
	// "bufio"
	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
	"github.com/lxn/win"
	"log"
)

func RunExWidgetDlg(owner walk.Form) (int, error) {
	var mw *walk.Dialog

	err := Dialog{
		AssignTo: &mw,
		Title:    "Drawing",
		//	MinSize:  Size{320, 240},
		//	Size:     Size{800, 600},
		MinSize:  Size{320, 240},
		Size:     Size{400, 300},
		Layout:   HBox{MarginsZero: true},
		Children: []Widget{},
	}.Create(owner)
	if err != nil {
		log.Fatal(err)
	}
	for _, name := range []string{"a", "b", "c"} {
		if w, err := NewMyWidget(mw); err != nil {
			log.Fatal(err)
		} else {
			w.SetName(name)
		}
	}
	mpb, err := NewMyPushButton(mw)
	if err != nil {
		log.Fatal(err)
	}
	mpb.SetText("This MyCustomButton")
	return mw.Run(), err
}

type MyWidget struct {
	walk.WidgetBase
}

const myWidgetWindowClass = "MyWidget Class"

func init() {
	walk.AppendToWalkInit(func() {
		walk.MustRegisterWindowClass(myWidgetWindowClass)
	})
}

func (*MyWidget) CreateLayoutItem(ctx *walk.LayoutContext) walk.LayoutItem {
	return &myWidgetLayoutItem{idealSize: walk.SizeFrom96DPI(walk.Size{50, 50}, ctx.DPI())}
}

func (w *MyWidget) WndProc(hwnd win.HWND, msg uint32, wParam, lParam uintptr) uintptr {
	switch msg {
	case win.WM_LBUTTONDOWN:
		log.Printf("%s: WM_LBUTTONDOWN", w.Name())
	}
	return w.WidgetBase.WndProc(hwnd, msg, wParam, lParam)
}

func NewMyWidget(parent walk.Container) (*MyWidget, error) {
	w := new(MyWidget)
	if err := walk.InitWidget(w, parent, myWidgetWindowClass, win.WS_VISIBLE, 0); err != nil {
		return nil, err
	}
	bg, err := walk.NewSolidColorBrush(walk.RGB(0, 255, 0))
	if err != nil {
		return nil, err
	}
	w.SetBackground(bg)

	return w, nil
}

type myWidgetLayoutItem struct {
	walk.LayoutItemBase
	idealSize walk.Size // in native pixels
}

func (li *myWidgetLayoutItem) LayoutFlags() walk.LayoutFlags {
	return 0
}

func (li *myWidgetLayoutItem) IdealSize() walk.Size {
	return li.idealSize
}

type MyPushButton struct {
	*walk.PushButton
}

func NewMyPushButton(parent walk.Container) (*MyPushButton, error) {
	pb, err := walk.NewPushButton(parent)
	if err != nil {
		return nil, err
	}
	mpb := &MyPushButton{pb}
	if err := walk.InitWrapperWindow(mpb); err != nil {
		return nil, err
	}
	return mpb, nil
}

func (mpb *MyPushButton) WndProc(hwnd win.HWND, msg uint32, wParam, lParam uintptr) uintptr {
	switch msg {
	case win.WM_LBUTTONDOWN:
		log.Printf("%s: WM_LBUTTONDOWN", mpb.Text())
	}

	return mpb.PushButton.WndProc(hwnd, msg, wParam, lParam)
}
