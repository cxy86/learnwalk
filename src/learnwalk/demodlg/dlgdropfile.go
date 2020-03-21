package demodlg

import (
	// "bufio"
	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
	//	"github.com/lxn/win"
)

type myDropFileDlg struct {
	*walk.MainWindow
	paintWidget *walk.CustomWidget
	picFile     string
}

func RunDropFileDlg(owner walk.Form) (int, error) {
	mw := new(myDropFileDlg)

	return MainWindow{
		AssignTo: &mw.MainWindow,
		Title:    "DropFile",
		//	MinSize:  Size{320, 240},
		//	Size:     Size{800, 600},
		MinSize: Size{320, 240},
		Size:    Size{400, 300},
		Layout:  VBox{MarginsZero: true},
		OnDropFiles: func(files []string) {
			if len(files) < 1 {
				return
			}
			mw.picFile = files[0]
			mw.Invalidate()
		},
		Children: []Widget{
			CustomWidget{
				AssignTo:            &mw.paintWidget,
				ClearsBackground:    true,
				InvalidatesOnResize: true,
				Paint:               mw.drawStuff,
			},
		},
	}.Run()

}

func (dlg *myDropFileDlg) drawStuff(canvas *walk.Canvas, updateBounds walk.Rectangle) error {
	bmp, err := dlg.createBitmap()
	if err != nil {
		return err
	}
	defer bmp.Dispose()
	if err := canvas.DrawImage(bmp, walk.Point{0, 0}); err != nil {
		return err
	}

	return nil
}

func (dlg *myDropFileDlg) createBitmap() (*walk.Bitmap, error) {
	//	bounds := dlg.paintWidget.Size()

	if dlg.picFile == "" {
		bmp, err := walk.NewBitmap(dlg.paintWidget.Size())
		return bmp, err
	} else {
		bmp, err := walk.NewBitmapFromFile(dlg.picFile)
		return bmp, err
	}

}
