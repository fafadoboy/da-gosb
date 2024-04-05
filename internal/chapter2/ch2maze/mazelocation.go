package ch2maze

type MazeLocation struct {
	row, column int
}

func NewMazeLocation(row, column int) MazeLocation {
	return MazeLocation{row: row, column: column}
}
