package hakoiri

// Solve returns Board list
func Solve(initBoard Board) Board {
	pastBoards := make(map[string]struct{})

	// board queue for check next
	q := []Board{initBoard}
	pastBoards[initBoard.ToHash()] = struct{}{}

	for {
		currentBoard := q[0]
		q = q[1:]
		for _, b := range nextBoards(currentBoard) {
			b.Turn = currentBoard.Turn + 1
			b.Prev = &currentBoard

			if b.IsGoal() {
				return b
			}
			h := b.ToHash()
			if _, ok := pastBoards[h]; !ok {
				pastBoards[h] = struct{}{}
				q = append(q, b)
			}
		}
	}
}

// アルゴリズム
// 左上から順にPanelを1枚ずつ走査する。
// Empty Panelに隣接しているかチェックし、その後1手で動かせる次の状態を算出する。
func nextBoards(b Board) []Board {
	result := []Board{}
	for h := 0; h < Height; h++ {
		for w := 0; w < Width; w++ {
			result = append(result, next(b, h, w, false)...)
		}
	}
	return result
}

func next(b Board, h int, w int, extra bool) []Board {
	p := b.Panels
	result := []Board{}
	switch p[h][w] {
	case PanelGirlTopLeft:
		// タテヨコ 2x2の左上

		// 左に移動
		if isEmpty(p, [2]int{h, w - 1}, [2]int{h + 1, w - 1}) {
			result = append(result, Board{Panels: moveLeft(p, h, w, 2, 2)})
		}
		// 右に移動
		if isEmpty(p, [2]int{h, w + 2}, [2]int{h + 1, w + 2}) {
			result = append(result, Board{Panels: moveRight(p, h, w, 2, 2)})
		}
		// 上に移動
		if isEmpty(p, [2]int{h - 1, w}, [2]int{h - 1, w + 1}) {
			result = append(result, Board{Panels: moveTop(p, h, w, 2, 2)})
		}
		// 下に移動
		if isEmpty(p, [2]int{h + 2, w}, [2]int{h + 2, w + 1}) {
			result = append(result, Board{Panels: moveBottom(p, h, w, 2, 2)})
		}
	case PanelFatherTop, PanelMotherTop, PanelGrandFatherTop, PanelGrandMotherTop:
		// タテヨコ 2x1の上

		// 左に移動
		if isEmpty(p, [2]int{h, w - 1}, [2]int{h + 1, w - 1}) {
			result = append(result, Board{Panels: moveLeft(p, h, w, 2, 1)})
		}
		// 右に移動
		if isEmpty(p, [2]int{h, w + 1}, [2]int{h + 1, w + 1}) {
			result = append(result, Board{Panels: moveRight(p, h, w, 2, 1)})
		}
		// 上に移動
		if isEmpty(p, [2]int{h - 1, w}) {
			result = append(result, Board{Panels: moveTop(p, h, w, 2, 1)})
		}
		// 上に2移動
		if isEmpty(p, [2]int{h - 1, w}, [2]int{h - 2, w}) {
			newP := moveTop(moveTop(p, h, w, 2, 1), h-1, w, 2, 1)
			result = append(result, Board{Panels: newP})
		}
		// 下に移動
		if isEmpty(p, [2]int{h + 2, w}) {
			result = append(result, Board{Panels: moveBottom(p, h, w, 2, 1)})
		}
		// 下に2移動
		if isEmpty(p, [2]int{h + 2, w}, [2]int{h + 3, w}) {
			newP := moveBottom(moveBottom(p, h, w, 2, 1), h+1, w, 2, 1)
			result = append(result, Board{Panels: newP})
		}
	case PanelBrotherLeft:
		// タテヨコ 1x2の左

		// 左に移動
		if isEmpty(p, [2]int{h, w - 1}) {
			result = append(result, Board{Panels: moveLeft(p, h, w, 1, 2)})
		}
		// 左に2移動
		if isEmpty(p, [2]int{h, w - 1}, [2]int{h, w - 2}) {
			newP := moveLeft(moveLeft(p, h, w, 1, 2), h, w-1, 1, 2)
			result = append(result, Board{Panels: newP})
		}
		// 右に移動
		if isEmpty(p, [2]int{h, w + 2}) {
			result = append(result, Board{Panels: moveRight(p, h, w, 1, 2)})
		}
		// 右に2移動
		if isEmpty(p, [2]int{h, w + 2}, [2]int{h, w + 3}) {
			newP := moveRight(moveRight(p, h, w, 1, 2), h, w+1, 1, 2)
			result = append(result, Board{Panels: newP})
		}
		// 上に移動
		if isEmpty(p, [2]int{h - 1, w}, [2]int{h - 1, w + 1}) {
			result = append(result, Board{Panels: moveTop(p, h, w, 1, 2)})
		}
		// 下に移動
		if isEmpty(p, [2]int{h + 1, w}, [2]int{h + 1, w + 1}) {
			result = append(result, Board{Panels: moveBottom(p, h, w, 1, 2)})
		}
	case PanelKoto, PanelFlower, PanelCalligraphy, PanelTea:
		// タテヨコ 1x1

		// 左に移動
		if isEmpty(p, [2]int{h, w - 1}) {
			newP := moveLeft(p, h, w, 1, 1)
			result = append(result, Board{Panels: newP})
			if !extra {
				result = append(result, next(Board{Panels: newP}, h, w-1, true)...)
			}
		}
		// 右に移動
		if isEmpty(p, [2]int{h, w + 1}) {
			newP := moveRight(p, h, w, 1, 1)
			result = append(result, Board{Panels: newP})
			if !extra {
				result = append(result, next(Board{Panels: newP}, h, w+1, true)...)
			}
		}
		// 上に移動
		if isEmpty(p, [2]int{h - 1, w}) {
			newP := moveTop(p, h, w, 1, 1)
			result = append(result, Board{Panels: newP})
			if !extra {
				result = append(result, next(Board{Panels: newP}, h-1, w, true)...)
			}
		}
		// 下に移動
		if isEmpty(p, [2]int{h + 1, w}) {
			newP := moveBottom(p, h, w, 1, 1)
			result = append(result, Board{Panels: newP})
			if !extra {
				result = append(result, next(Board{Panels: newP}, h+1, w, true)...)
			}
		}
	default:
	}
	return result
}

func isEmpty(p [Height][Width]Panel, targets ...[2]int) bool {
	for _, t := range targets {
		// check range
		if t[0] < 0 || Height <= t[0] {
			return false
		}
		if t[1] < 0 || Width <= t[1] {
			return false
		}
		// check panel
		if p[t[0]][t[1]] != PanelEmpty {
			return false
		}
	}
	return true
}

func deepCopy(p [Height][Width]Panel) [Height][Width]Panel {
	newP := [Height][Width]Panel{}
	for h := 0; h < Height; h++ {
		for w := 0; w < Width; w++ {
			newP[h][w] = p[h][w]
		}
	}
	return newP
}

func moveLeft(p [Height][Width]Panel, h, w, ph, pw int) [Height][Width]Panel {
	newP := deepCopy(p)

	swap(&newP, h, w-1, h, w)
	if pw == 2 {
		swap(&newP, h, w, h, w+1)
	}

	if ph == 2 {
		swap(&newP, h+1, w-1, h+1, w)
		if pw == 2 {
			swap(&newP, h+1, w, h+1, w+1)
		}
	}

	return newP
}

func moveRight(p [Height][Width]Panel, h, w, ph, pw int) [Height][Width]Panel {
	newP := deepCopy(p)

	swap(&newP, h, w, h, w+1)
	if pw == 2 {
		swap(&newP, h, w, h, w+2)
	}

	if ph == 2 {
		swap(&newP, h+1, w, h+1, w+1)
		if pw == 2 {
			swap(&newP, h+1, w, h+1, w+2)
		}
	}

	return newP
}

func moveTop(p [Height][Width]Panel, h, w, ph, pw int) [Height][Width]Panel {
	newP := deepCopy(p)

	swap(&newP, h, w, h-1, w)
	if ph == 2 {
		swap(&newP, h, w, h+1, w)
	}

	if pw == 2 {
		swap(&newP, h, w+1, h-1, w+1)
		if ph == 2 {
			swap(&newP, h, w+1, h+1, w+1)
		}
	}

	return newP
}

func moveBottom(p [Height][Width]Panel, h, w, ph, pw int) [Height][Width]Panel {
	newP := deepCopy(p)

	swap(&newP, h, w, h+1, w)
	if ph == 2 {
		swap(&newP, h, w, h+2, w)
	}

	if pw == 2 {
		swap(&newP, h, w+1, h+1, w+1)
		if ph == 2 {
			swap(&newP, h, w+1, h+2, w+1)
		}
	}

	return newP
}

func swap(p *[Height][Width]Panel, h1, w1, h2, w2 int) {
	tmp := p[h1][w1]
	p[h1][w1] = p[h2][w2]
	p[h2][w2] = tmp
}
