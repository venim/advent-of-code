package main

import (
	_ "embed"
	"flag"
	"slices"
	"strings"
	"time"

	"github.com/golang/glog"
	"github.com/venim/advent-of-code/util"
)

var (
	//go:embed in.txt
	input     string
	cardRank  = []byte{'A', 'K', 'Q', 'J', 'T', '9', '8', '7', '6', '5', '4', '3', '2'}
	cardRank2 = []byte{'A', 'K', 'Q', 'T', '9', '8', '7', '6', '5', '4', '3', '2', 'J'}
)

type hand struct {
	cards  string
	score  int
	result HandResult
}

type HandResult int

const (
	HighCard HandResult = iota + 1
	OnePair
	TwoPair
	ThreeOfAKind
	FullHouse
	FourOfAKind
	FiveOfAKind
)

func getResult(cards string) HandResult {
	cardFreq := map[rune]int{}
	for _, c := range cards {
		cardFreq[c]++
	}
	if len(cardFreq) == 1 {
		return FiveOfAKind
	}
	if len(cardFreq) == 2 {
		for _, n := range cardFreq {
			if n == 4 {
				return FourOfAKind
			}
		}
		return FullHouse
	}
	nPairs := 0
	for _, n := range cardFreq {
		if n == 3 {
			return ThreeOfAKind
		}
		if n == 2 {
			nPairs++
		}
	}
	switch nPairs {
	case 2:
		return TwoPair
	case 1:
		return OnePair
	default:
		return HighCard
	}
}

func sortHands(a, b hand, rank []byte) int {
	if a.result < b.result {
		return -1
	}
	if a.result > b.result {
		return 1
	}
	for i := 0; i < 5; i++ {
		aRank := slices.Index(rank, a.cards[i])
		bRank := slices.Index(rank, b.cards[i])
		if aRank > bRank {
			return -1
		}
		if aRank < bRank {
			return 1
		}
	}
	return 0
}

func part1(lines []string) (res int) {
	hands := []hand{}
	for _, line := range lines {
		if line == "" {
			continue
		}
		parts := strings.Split(line, " ")
		h := hand{
			cards:  parts[0],
			score:  util.MustAtoi(parts[1]),
			result: getResult(parts[0]),
		}
		hands = append(hands, h)
	}

	slices.SortFunc(hands, func(a, b hand) int {
		return sortHands(a, b, cardRank)
	})
	for i, h := range hands {
		res += h.score * (i + 1)
	}
	return
}

func replaceJ(hand string) string {
	n := map[rune]int{}
	j := 0
	for _, c := range hand {
		if c == 'J' {
			j++
		} else {
			n[c]++
		}
	}
	maxC := '0'
	maxN := 0
	for c, n := range n {
		if n > maxN {
			maxN = n
			maxC = c
		}
	}
	return strings.Replace(hand, "J", string(maxC), -1)
}

func part2(lines []string) (res int) {
	hands := []hand{}
	for _, line := range lines {
		if line == "" {
			continue
		}
		parts := strings.Split(line, " ")
		h := hand{
			cards:  parts[0],
			score:  util.MustAtoi(parts[1]),
			result: getResult(replaceJ(parts[0])),
		}
		hands = append(hands, h)
	}

	slices.SortFunc(hands, func(a, b hand) int {
		return sortHands(a, b, cardRank2)
	})
	for i, h := range hands {
		res += h.score * (i + 1)
	}
	return
}

func init() {
	flag.Set("logtostderr", "true")
}

func main() {
	var (
		t   time.Time
		res int
	)
	flag.Parse()
	lines := strings.Split(input, "\n")

	t = time.Now()
	res = part1(lines)
	glog.Infof("[Part 1] = %v", res)
	glog.Infof("took %s", time.Since(t))

	t = time.Now()
	res = part2(lines)
	glog.Infof("[Part 2] = %v", res)
	glog.Infof("took %s", time.Since(t))
}
