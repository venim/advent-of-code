package util

import (
	"log/slog"
	"os"
	"strconv"
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
