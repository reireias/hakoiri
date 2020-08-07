package hakoiri

import (
	"testing"
)

func TestSolve(t *testing.T) {
	b := Board{Panels: deepCopy(DefaultPanels), Turn: 0, Prev: nil}
	goal := Solve(b)
	if goal.Turn != 81 {
		t.Fail()
	}
}

func TestNext2x2(t *testing.T) {
	p := deepCopy(DefaultPanels)
	// cannot move
	b := Board{Panels: p}
	boards := next(b, 0, 1, false)
	if len(boards) != 0 {
		t.Fail()
	}

	// can move left
	p[0][0] = PanelEmpty
	p[1][0] = PanelEmpty
	b = Board{Panels: p}
	boards = next(b, 0, 1, false)
	if len(boards) != 1 {
		t.Fail()
	}
	if boards[0].Panels[0][0] != PanelGirlTopLeft {
		t.Fail()
	}

	// can move bottom
	p = deepCopy(DefaultPanels)
	p[2][1] = PanelEmpty
	p[2][2] = PanelEmpty
	b = Board{Panels: p}
	boards = next(b, 0, 1, false)
	if len(boards) != 1 {
		t.Fail()
	}
	if boards[0].Panels[1][1] != PanelGirlTopLeft {
		t.Fail()
	}
}

func TestNext2x1(t *testing.T) {
	p := deepCopy(DefaultPanels)
	// cannot move
	b := Board{Panels: p}
	boards := next(b, 0, 0, false)
	if len(boards) != 0 {
		t.Fail()
	}

	// can move right
	p[0][1] = PanelEmpty
	p[1][1] = PanelEmpty
	b = Board{Panels: p}
	boards = next(b, 0, 0, false)
	if len(boards) != 1 {
		t.Fail()
	}
	if boards[0].Panels[0][1] != PanelFatherTop {
		t.Fail()
	}

	// can move bottom, 2 * bottom
	p = deepCopy(DefaultPanels)
	p[2][0] = PanelEmpty
	p[3][0] = PanelEmpty
	b = Board{Panels: p}
	boards = next(b, 0, 0, false)
	if len(boards) != 2 {
		t.Fail()
	}
}

func TestNext1x2(t *testing.T) {
	p := deepCopy(DefaultPanels)
	// cannot move
	b := Board{Panels: p}
	boards := next(b, 2, 1, false)
	if len(boards) != 0 {
		t.Fail()
	}

	// can move right, 2 * right
	p[2][0] = PanelBrotherLeft
	p[2][1] = PanelBrotherRight
	p[2][2] = PanelEmpty
	p[2][3] = PanelEmpty
	b = Board{Panels: p}
	boards = next(b, 2, 0, false)
	if len(boards) != 2 {
		t.Fail()
	}

	// can move bottom
	p = deepCopy(DefaultPanels)
	p[3][1] = PanelEmpty
	p[3][2] = PanelEmpty
	b = Board{Panels: p}
	boards = next(b, 2, 1, false)
	if len(boards) != 1 {
		t.Fail()
	}
}

func TestNext1x1(t *testing.T) {
	p := deepCopy(DefaultPanels)
	// can move rignt, 2 * right
	b := Board{Panels: p}
	boards := next(b, 4, 0, false)
	if len(boards) < 2 {
		t.Fail()
	}

	// can move bottom, bottom + right
	boards = next(b, 3, 1, false)
	if len(boards) < 2 {
		t.Fail()
	}
}

func TestMove(t *testing.T) {
	p := deepCopy(DefaultPanels)

	// 0, 1の座標に左上がある2x2のPanelを左へ移動
	moved := moveLeft(p, 0, 1, 2, 2)

	if moved[0][0] != p[0][1] ||
		moved[0][1] != p[0][2] ||
		moved[0][2] != p[0][0] ||
		moved[1][0] != p[1][1] ||
		moved[1][1] != p[1][2] ||
		moved[1][2] != p[1][0] {
		t.Fail()
	}

	// 戻す
	moved = moveRight(moved, 0, 0, 2, 2)
	if moved[0][0] != p[0][0] ||
		moved[1][1] != p[1][1] {
		t.Fail()
	}
}

func TestDeepCopy(t *testing.T) {
	p := DefaultPanels
	newP := deepCopy(p)

	newP[0][0] = PanelEmpty

	if p[0][0] == PanelEmpty {
		t.Fail()
	}
}
