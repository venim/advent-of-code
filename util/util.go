package util

import (
	"log/slog"
	"os"
	"strconv"
	"strings"
)

var (
	Directions = map[byte]Pos{'>': {0, 1}, '<': {0, -1}, '^': {-1, 0}, 'v': {1, 0}}
	Logger     = slog.New(slog.NewJSONHandler(os.Stderr, nil))
)

func MustAtoi(s string) int {
	i, _ := strconv.Atoi(s)
	return i
}

type Pos struct {
	X, Y int
}

func (p Pos) Add(p2 Pos) Pos {
	return Pos{X: p.X + p2.X, Y: p.Y + p2.Y}
}

func (p Pos) Sub(p2 Pos) Pos {
	return Pos{X: p.X - p2.X, Y: p.Y - p2.Y}
}

func (p Pos) IsOutOfBounds(x, y int) bool {
	return p.Y < 0 || p.Y >= y || p.X < 0 || p.X >= x
}

func SplitLines(s string) []string {
	return strings.Split(strings.ReplaceAll(s, "\r", ""), "\n")
}
