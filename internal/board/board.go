package board

import (
	"strings"
)

// Height is board panel height
const Height = 5

// Width is board panel width
const Width = 4

// PanelHeight is lines in a panel
const PanelHeight = 3

// StringHeight is lines in a board
const StringHeight = Height * PanelHeight

// StringWidth is runes of line in a board
const StringWidth = 4

// Panel shows panel type in board
type Panel int

// panel types
const (
	PanelGirlTopLeft Panel = iota
	PanelGirlTopRight
	PanelGirlBottomLeft
	PanelGirlBottomRight
	PanelFatherTop
	PanelFatherBottom
	PanelMotherTop
	PanelMotherBottom
	PanelGrandFatherTop
	PanelGrandFatherBottom
	PanelGrandMotherTop
	PanelGrandMotherBottom
	PanelBrotherLeft
	PanelBrotherRight
	PanelKoto
	PanelFlower
	PanelCalligraphy
	PanelTea
	PanelEmpty
)

// PanelStringMap is a map for Panel to string
var PanelStringMap = map[Panel][]string{
	PanelGirlTopLeft: {
		"+------",
		"|      ",
		"|   箱 ",
	},
	PanelGirlTopRight: {
		"-----+ ",
		"     | ",
		"入   | ",
	},
	PanelGirlBottomLeft: {
		"|   り ",
		"|      ",
		"+------",
	},
	PanelGirlBottomRight: {
		"娘   | ",
		"     | ",
		"-----+ ",
	},
	PanelFatherTop: {
		"+----+ ",
		"| 父 | ",
		"|    | ",
	},
	PanelFatherBottom: {
		"|    | ",
		"| 親 | ",
		"+----+ ",
	},
	PanelMotherTop: {
		"+----+ ",
		"| 母 | ",
		"|    | ",
	},
	PanelMotherBottom: {
		"|    | ",
		"| 親 | ",
		"+----+ ",
	},
	PanelGrandFatherTop: {
		"+----+ ",
		"| 祖 | ",
		"|    | ",
	},
	PanelGrandFatherBottom: {
		"|    | ",
		"| 父 | ",
		"+----+ ",
	},
	PanelGrandMotherTop: {
		"+----+ ",
		"| 祖 | ",
		"|    | ",
	},
	PanelGrandMotherBottom: {
		"|    | ",
		"| 母 | ",
		"+----+ ",
	},
	PanelBrotherLeft: {
		"+------",
		"| 兄   ",
		"+------",
	},
	PanelBrotherRight: {
		"-----+ ",
		"  弟 | ",
		"-----+ ",
	},
	PanelKoto: {
		"+----+ ",
		"| 琴 | ",
		"+----+ ",
	},
	PanelFlower: {
		"+----+ ",
		"| 華 | ",
		"+----+ ",
	},
	PanelCalligraphy: {
		"+----+ ",
		"| 書 | ",
		"+----+ ",
	},
	PanelTea: {
		"+----+ ",
		"| 茶 | ",
		"+----+ ",
	},
	PanelEmpty: {
		"       ",
		"       ",
		"       ",
	},
}

// Board is state of hakoiri puzzle
type Board struct {
	Panels [Height][Width]Panel
}

// ToString returns board state as string
func (b *Board) ToString() string {
	strBoard := [StringHeight][StringWidth]string{}
	for i := 0; i < Height; i++ {
		for j := 0; j < Width; j++ {
			panel := b.Panels[i][j]
			for k := 0; k < PanelHeight; k++ {
				strBoard[i*3+k][j] = PanelStringMap[panel][k]
			}
		}
	}

	lines := [StringHeight]string{}
	for i := 0; i < StringHeight; i++ {
		lines[i] = strings.Join(strBoard[i][:], "")
	}
	return strings.Join(lines[:], "\n")
}
