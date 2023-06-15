package pattern

import "math/rand"

type Room struct {
	chunks [][]int
	Doors  [][]bool // другой способ?
}

func initialization(patternOfChunks [][]int, DoorsOfChunks [][]bool) *Room {
	R := &Room{
		chunks: patternOfChunks,
		Doors:  DoorsOfChunks,
	}

	return R
}

func randomChoice(patternOfChunks Room) {

	Random := rand.Int()

}

func main() {

	patternOfChunks := [][]int{
		{
			// 1
			4, 4, 4, 4, 4, 4, 1, 2, 2, 1, 4, 4, 4, 4, 4, 4,
			4, 1, 1, 1, 1, 1, 1, 2, 2, 1, 1, 1, 1, 1, 1, 4,
			4, 1, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 1, 4,
			4, 1, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 1, 4,
			4, 1, 2, 2, 3, 3, 2, 2, 2, 2, 3, 3, 2, 2, 1, 4,
			4, 1, 2, 2, 3, 3, 2, 2, 2, 2, 3, 3, 2, 2, 1, 4,
			1, 1, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 1, 1,
			2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2,
			2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2,
			1, 1, 2, 2, 3, 3, 2, 2, 2, 2, 3, 3, 2, 2, 1, 1,
			4, 1, 2, 2, 3, 3, 2, 2, 2, 2, 3, 3, 2, 2, 1, 4,
			4, 1, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 1, 4,
			4, 1, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 1, 4,
			4, 1, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 1, 4,
			4, 1, 1, 1, 1, 1, 1, 2, 2, 1, 1, 1, 1, 1, 1, 4,
			4, 4, 4, 4, 4, 4, 1, 2, 2, 1, 4, 4, 4, 4, 4, 4,
		},
	}

	DoorsOfChunks := [][]bool{
		{
			// 1
			true, // Up
			true, // Right
			true, // Down
			true, // Left
		},
	}

	patternOfRooms := initialization(patternOfChunks, DoorsOfChunks)
}
