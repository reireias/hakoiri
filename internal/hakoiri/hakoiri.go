package hakoiri

// Solve returns Board list
func Solve(initBoard Board) []Board {
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
				return makeResult(&b)
			}
			h := b.ToHash()
			if _, ok := pastBoards[h]; !ok {
				pastBoards[h] = struct{}{}
				q = append(q, b)
			}
		}
	}
}

// Algorithm
// 1. Scan panels from the upper left.
// 2. Check if Empty Panel is nearby.
// 3. Calculate the next state.
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
	currentHash := b.ToHash()
	p := b.Panels
	result := []Board{}

	size, ok := sizeMap[p[h][w]]
	if !ok {
		return result
	}

	for _, direction := range Directions {
		deltas := moveCheckMap[key{direction: direction, size: size}]
		checks := [][2]int{}
		for _, d := range deltas {
			checks = append(checks, [2]int{h + d.height, w + d.width})
		}
		if isEmpty(p, checks...) {
			movedP := move(p, h, w, direction)
			result = append(result, Board{
				Panels: movedP,
				Moved:  p[h][w],
			})
			if !extra {
				moveDelta := moveCheckMap[key{direction: direction, size: "1x1"}][0]
				extras := next(Board{Panels: movedP}, h+moveDelta.height, w+moveDelta.width, true)
				for _, extraBoard := range extras {
					if extraBoard.ToHash() != currentHash {
						result = append(result, extraBoard)
					}
				}
			}
		}
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

func move(p [Height][Width]Panel, h int, w int, direction Direction) [Height][Width]Panel {
	newP := deepCopy(p)
	size := sizeMap[p[h][w]]
	k := key{direction: direction, size: size}
	for _, d := range moveSwapMap[k] {
		swap(&newP, h+d.h1, w+d.w1, h+d.h2, w+d.w2)
	}
	return newP
}

func swap(p *[Height][Width]Panel, h1, w1, h2, w2 int) {
	tmp := p[h1][w1]
	p[h1][w1] = p[h2][w2]
	p[h2][w2] = tmp
}

func makeResult(b *Board) []Board {
	result := make([]Board, b.Turn+1)

	current := b
	for i := 0; i < b.Turn+1; i++ {
		result[b.Turn-i] = *current
		current = current.Prev
	}
	return result
}
