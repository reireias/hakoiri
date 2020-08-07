package main

import (
	"fmt"
	"log"
	"time"

	"github.com/jroimartin/gocui"
	"github.com/reireias/hakoiri/internal/hakoiri"
)

const dummy = `
+----+ +-----------+ +----+
| 父 | |           | | 母 |
|    | |   箱 入   | |    |
|    | |   り 娘   | |    |
| 親 | |           | | 親 |
+----+ +-----------+ +----+
+----+ +-----------+ +----+
| 祖 | | 兄     弟 | | 祖 |
|    | +-----------+ |    |
|    | +----+ +----+ |    |
| 父 | | 華 | | 書 | | 母 |
+----+ +----+ +----+ +----+
+----+               +----+
| 琴 |               | 茶 |
+----+               +----+
`

const p1 = `+----+
| 琴  |
+----+`

func main() {
	hakoiri.Solve(hakoiri.Board{Panels: hakoiri.DefaultPanels})
	return
	g, err := gocui.NewGui(gocui.OutputNormal)
	if err != nil {
		log.Panicln(err)
	}
	defer g.Close()

	g.SetManagerFunc(layout)

	// setUpdate(g)

	if err := g.SetKeybinding("", gocui.KeyCtrlC, gocui.ModNone, quit); err != nil {
		log.Panicln(err)
	}

	if err := g.SetKeybinding("", 'q', gocui.ModNone, quit); err != nil {
		log.Panicln(err)
	}

	if err := g.MainLoop(); err != nil && err != gocui.ErrQuit {
		log.Panicln(err)
	}
}

func layout(g *gocui.Gui) error {
	// v1, _ := g.SetView("v1", 0, 0, 8, 5)
	// v1.Frame = false
	// fmt.Fprintln(v, p1)
	v1, err := g.SetView("hello", 0, 0, 8, 5)
	if err != nil {
		// if err != gocui.ErrUnknownView {
		// 	return err
		// }
		v1.Frame = false
		// fmt.Fprintln(v, "Hello world!")
		fmt.Fprintln(v1, p1)
	}
	return nil
}

func setUpdate(g *gocui.Gui) {
	go func() {
		time.Sleep(5 * time.Second)
		g.Update(func(gui *gocui.Gui) error {
			v, _ := g.View("hello")
			v.Clear()
			fmt.Fprintln(v, dummy)
			return nil
		})
	}()
}

func quit(g *gocui.Gui, v *gocui.View) error {
	return gocui.ErrQuit
}
