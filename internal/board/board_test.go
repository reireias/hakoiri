package board

import (
	"testing"
)

var DefaultPanels = [Height][Width]Panel{
	{PanelFatherTop, PanelGirlTopLeft, PanelGirlTopRight, PanelMotherTop},
	{PanelFatherBottom, PanelGirlBottomLeft, PanelGirlBottomRight, PanelMotherBottom},
	{PanelGrandFatherTop, PanelBrotherLeft, PanelBrotherRight, PanelGrandMotherTop},
	{PanelGrandFatherBottom, PanelFlower, PanelCalligraphy, PanelGrandMotherBottom},
	{PanelKoto, PanelEmpty, PanelEmpty, PanelTea},
}

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
