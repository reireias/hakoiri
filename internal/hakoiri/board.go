package hakoiri

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

// Direction is panel move direction
type Direction int

// directions
const (
	Left Direction = iota
	Right
	Top
	Bottom
)

// Directions is all values of Direction
var Directions = []Direction{Left, Right, Top, Bottom}

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

var sizeMap = map[Panel]string{
	PanelGirlTopLeft:    "2x2",
	PanelFatherTop:      "2x1",
	PanelMotherTop:      "2x1",
	PanelGrandFatherTop: "2x1",
	PanelGrandMotherTop: "2x1",
	PanelBrotherLeft:    "1x2",
	PanelKoto:           "1x1",
	PanelFlower:         "1x1",
	PanelCalligraphy:    "1x1",
	PanelTea:            "1x1",
}

type key struct {
	direction Direction
	size      string
}

type delta struct {
	height int
	width  int
}

// 移動させる際にEmptyかをチェックすべき位置を相対値で表すmap
var moveCheckMap = map[key][]delta{
	{direction: Left, size: "2x2"}:   {delta{height: 0, width: -1}, delta{height: 1, width: -1}},
	{direction: Right, size: "2x2"}:  {delta{height: 0, width: 2}, delta{height: 1, width: 2}},
	{direction: Top, size: "2x2"}:    {delta{height: -1, width: 0}, delta{height: -1, width: 1}},
	{direction: Bottom, size: "2x2"}: {delta{height: 2, width: 0}, delta{height: 2, width: 1}},
	{direction: Left, size: "2x1"}:   {delta{height: 0, width: -1}, delta{height: 1, width: -1}},
	{direction: Right, size: "2x1"}:  {delta{height: 0, width: 1}, delta{height: 1, width: 1}},
	{direction: Top, size: "2x1"}:    {delta{height: -1, width: 0}},
	{direction: Bottom, size: "2x1"}: {delta{height: 2, width: 0}},
	{direction: Left, size: "1x2"}:   {delta{height: 0, width: -1}},
	{direction: Right, size: "1x2"}:  {delta{height: 0, width: 2}},
	{direction: Top, size: "1x2"}:    {delta{height: -1, width: 0}, delta{height: -1, width: 1}},
	{direction: Bottom, size: "1x2"}: {delta{height: 1, width: 0}, delta{height: 1, width: 1}},
	{direction: Left, size: "1x1"}:   {delta{height: 0, width: -1}},
	{direction: Right, size: "1x1"}:  {delta{height: 0, width: 1}},
	{direction: Top, size: "1x1"}:    {delta{height: -1, width: 0}},
	{direction: Bottom, size: "1x1"}: {delta{height: 1, width: 0}},
}

type swapDelta struct {
	h1, w1, h2, w2 int
}

// サイズと移動方向をkey、swapすべき座標のリストをvalueにもつmap
var moveSwapMap = map[key][]swapDelta{
	{direction: Left, size: "2x2"}: {
		swapDelta{h1: 0, w1: -1, h2: 0, w2: 0},
		swapDelta{h1: 0, w1: 0, h2: 0, w2: 1},
		swapDelta{h1: 1, w1: -1, h2: 1, w2: 0},
		swapDelta{h1: 1, w1: 0, h2: 1, w2: 1},
	},
	{direction: Right, size: "2x2"}: {
		swapDelta{h1: 0, w1: 2, h2: 0, w2: 1},
		swapDelta{h1: 0, w1: 1, h2: 0, w2: 0},
		swapDelta{h1: 1, w1: 2, h2: 1, w2: 1},
		swapDelta{h1: 1, w1: 1, h2: 1, w2: 0},
	},
	{direction: Top, size: "2x2"}: {
		swapDelta{h1: -1, w1: 0, h2: 0, w2: 0},
		swapDelta{h1: 0, w1: 0, h2: 1, w2: 0},
		swapDelta{h1: -1, w1: 1, h2: 0, w2: 1},
		swapDelta{h1: 0, w1: 1, h2: 1, w2: 1},
	},
	{direction: Bottom, size: "2x2"}: {
		swapDelta{h1: 2, w1: 0, h2: 1, w2: 0},
		swapDelta{h1: 1, w1: 0, h2: 0, w2: 0},
		swapDelta{h1: 2, w1: 1, h2: 1, w2: 1},
		swapDelta{h1: 1, w1: 1, h2: 0, w2: 1},
	},
	{direction: Left, size: "2x1"}: {
		swapDelta{h1: 0, w1: -1, h2: 0, w2: 0},
		swapDelta{h1: 1, w1: -1, h2: 1, w2: 0},
	},
	{direction: Right, size: "2x1"}: {
		swapDelta{h1: 0, w1: 0, h2: 0, w2: 1},
		swapDelta{h1: 1, w1: 0, h2: 1, w2: 1},
	},
	{direction: Top, size: "2x1"}: {
		swapDelta{h1: -1, w1: 0, h2: 0, w2: 0},
		swapDelta{h1: 0, w1: 0, h2: 1, w2: 0},
	},
	{direction: Bottom, size: "2x1"}: {
		swapDelta{h1: 2, w1: 0, h2: 1, w2: 0},
		swapDelta{h1: 1, w1: 0, h2: 0, w2: 0},
	},
	{direction: Left, size: "1x2"}: {
		swapDelta{h1: 0, w1: -1, h2: 0, w2: 0},
		swapDelta{h1: 0, w1: 0, h2: 0, w2: 1},
	},
	{direction: Right, size: "1x2"}: {
		swapDelta{h1: 0, w1: 2, h2: 0, w2: 1},
		swapDelta{h1: 0, w1: 1, h2: 0, w2: 0},
	},
	{direction: Top, size: "1x2"}: {
		swapDelta{h1: -1, w1: 0, h2: 0, w2: 0},
		swapDelta{h1: -1, w1: 1, h2: 0, w2: 1},
	},
	{direction: Bottom, size: "1x2"}: {
		swapDelta{h1: 0, w1: 0, h2: 1, w2: 0},
		swapDelta{h1: 0, w1: 1, h2: 1, w2: 1},
	},
	{direction: Left, size: "1x1"}:   {swapDelta{h1: 0, w1: 0, h2: 0, w2: -1}},
	{direction: Right, size: "1x1"}:  {swapDelta{h1: 0, w1: 0, h2: 0, w2: 1}},
	{direction: Top, size: "1x1"}:    {swapDelta{h1: 0, w1: 0, h2: -1, w2: 0}},
	{direction: Bottom, size: "1x1"}: {swapDelta{h1: 0, w1: 0, h2: 1, w2: 0}},
}

// サイズのみでpanelの種類を区別せずに盤面をhash化する際に利用するmap
var sizeTypeMap = map[Panel]string{
	PanelGirlTopLeft:       "A1",
	PanelGirlTopRight:      "A2",
	PanelGirlBottomLeft:    "A3",
	PanelGirlBottomRight:   "A4",
	PanelFatherTop:         "B1",
	PanelFatherBottom:      "B2",
	PanelMotherTop:         "B1",
	PanelMotherBottom:      "B2",
	PanelGrandFatherTop:    "B1",
	PanelGrandFatherBottom: "B2",
	PanelGrandMotherTop:    "B1",
	PanelGrandMotherBottom: "B2",
	PanelBrotherLeft:       "C1",
	PanelBrotherRight:      "C2",
	PanelKoto:              "D1",
	PanelFlower:            "D1",
	PanelCalligraphy:       "D1",
	PanelTea:               "D1",
	PanelEmpty:             "E1",
}

// DefaultPanels is standard initial panels
var DefaultPanels = [Height][Width]Panel{
	{PanelFatherTop, PanelGirlTopLeft, PanelGirlTopRight, PanelMotherTop},
	{PanelFatherBottom, PanelGirlBottomLeft, PanelGirlBottomRight, PanelMotherBottom},
	{PanelGrandFatherTop, PanelBrotherLeft, PanelBrotherRight, PanelGrandMotherTop},
	{PanelGrandFatherBottom, PanelFlower, PanelCalligraphy, PanelGrandMotherBottom},
	{PanelKoto, PanelEmpty, PanelEmpty, PanelTea},
}

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
	Turn   int
	Prev   *Board
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

// ToHash returns string hash of board's panels
func (b *Board) ToHash() string {
	lines := [Height]string{}
	for i := 0; i < Height; i++ {
		line := [Width]string{}
		for j := 0; j < Width; j++ {
			line[j] = sizeTypeMap[b.Panels[i][j]]
		}
		lines[i] = strings.Join(line[:], "")
	}
	return strings.Join(lines[:], "")
}

// IsGoal returns true if GirlPanel can go out
func (b *Board) IsGoal() bool {
	p := b.Panels
	return p[3][1] == PanelGirlTopLeft
}
