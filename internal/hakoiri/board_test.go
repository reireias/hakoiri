package hakoiri

import (
	"testing"
)

func TestToString(t *testing.T) {
	board := Board{Panels: DefaultPanels}
	expect := `+----+ +-----------+ +----+ 
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
+----+               +----+ `
	if board.ToString() != expect {
		t.Fail()
	}
}

func TestToHash(t *testing.T) {
	board := Board{Panels: DefaultPanels}
	expect := "B1A1A2B1B2A3A4B2B1C1C2B1B2D1D1B2D1E1E1D1"
	if board.ToHash() != expect {
		t.Fail()
	}
}

func TestIsGoal(t *testing.T) {
	p := deepCopy(DefaultPanels)
	b := Board{Panels: p}
	if b.IsGoal() {
		t.Fail()
	}

	p[3][1] = PanelGirlTopLeft
	p[3][2] = PanelGirlTopRight
	p[4][1] = PanelGirlBottomLeft
	p[4][2] = PanelGirlBottomRight
	b = Board{Panels: p}
	if !b.IsGoal() {
		t.Fail()
	}
}
