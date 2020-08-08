package cui

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/jroimartin/gocui"
	"github.com/reireias/hakoiri/internal/hakoiri"
)

// State is current cui state
type State struct {
	result []hakoiri.Board
	turn   int
}

var state = State{turn: 0}

// Start func executes CUI
func Start() {
	g := initialize()
	defer g.Close()
	keybind(g)
	main(g)
}

const puzzleViewWidth = 29
const puzzleViewHeight = 17
const headerHeight = 2

func initialize() *gocui.Gui {
	state.result = hakoiri.Solve(hakoiri.Board{Panels: hakoiri.DefaultPanels})
	g, err := gocui.NewGui(gocui.OutputNormal)
	if err != nil {
		log.Panicln(err)
	}

	g.SetManagerFunc(layout)

	return g
}

func keybind(g *gocui.Gui) {
	if err := g.SetKeybinding("", 'n', gocui.ModNone, next); err != nil {
		log.Panicln(err)
	}

	if err := g.SetKeybinding("", 'p', gocui.ModNone, prev); err != nil {
		log.Panicln(err)
	}

	if err := g.SetKeybinding("", gocui.KeyCtrlC, gocui.ModNone, quit); err != nil {
		log.Panicln(err)
	}

	if err := g.SetKeybinding("", 'q', gocui.ModNone, quit); err != nil {
		log.Panicln(err)
	}
}

func main(g *gocui.Gui) {
	if err := g.MainLoop(); err != nil && err != gocui.ErrQuit {
		log.Panicln(err)
	}
}

func layout(g *gocui.Gui) error {
	maxX, maxY := g.Size()
	puzzleView, err := g.SetView(
		"puzzle",
		(maxX-puzzleViewWidth)/2,
		(maxY-puzzleViewHeight)/2,
		(maxX+puzzleViewWidth)/2,
		(maxY+puzzleViewHeight)/2,
	)
	if err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		puzzleView.Frame = false
		fmt.Fprintln(puzzleView, toViewString(state.result[0].ToString()))
	}

	headerView, err := g.SetView(
		"header",
		(maxX-puzzleViewWidth)/2,
		(maxY-puzzleViewHeight)/2-1-headerHeight,
		(maxX+puzzleViewWidth)/2,
		(maxY-puzzleViewHeight)/2-1,
	)
	if err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		headerView.Frame = false
		fmt.Fprintln(headerView, "Turn: 0")
	}
	return nil
}

// handlers

func next(g *gocui.Gui, v *gocui.View) error {
	if state.turn == len(state.result)-1 {
		return nil
	}
	g.Update(func(gui *gocui.Gui) error {
		state.turn++
		update(g, state.turn)
		return nil
	})
	return nil
}

func prev(g *gocui.Gui, v *gocui.View) error {
	if state.turn == 0 {
		return nil
	}
	g.Update(func(gui *gocui.Gui) error {
		state.turn--
		update(g, state.turn)
		return nil
	})
	return nil
}

func update(g *gocui.Gui, turn int) {
	puzzleView, _ := g.View("puzzle")
	puzzleView.Clear()
	fmt.Fprintln(puzzleView, toViewString(state.result[turn].ToString()))

	headerView, _ := g.View("header")
	headerView.Clear()
	fmt.Fprintln(headerView, "Turn: "+strconv.Itoa(turn))
}

func quit(g *gocui.Gui, v *gocui.View) error {
	return gocui.ErrQuit
}

// gocuiが利用しているtcellは全角文字でも半角サイズとして扱う
// そのため、全角文字の場合は、後ろにスペースを挿入する
func toViewString(s string) string {
	result := []string{}
	for _, c := range s {
		str := string([]rune{c})
		result = append(result, str)
		if len(str) > 1 {
			result = append(result, " ")
		}
	}
	return strings.Join(result, "") + "\n           出  口"
}
