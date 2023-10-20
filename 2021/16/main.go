package main

import (
	_ "embed"
	"flag"
	"fmt"
	"slices"
	"strconv"
	"strings"
	"time"

	"github.com/golang/glog"
)

var (
	//go:embed in.txt
	input string
)

func toBits(input string) string {
	bits := ""
	for _, r := range input {
		i, _ := strconv.ParseInt(string(r), 16, 64)
		bits += fmt.Sprintf("%04b", i)
	}
	return bits
}

type transmission struct {
	bits       string
	pos        int
	versionSum int
}

func (t *transmission) parseN(len int) int {
	chunk := t.bits[t.pos : t.pos+len]
	i, _ := strconv.ParseInt(chunk, 2, 64)
	t.pos += len
	return int(i)
}

func (t *transmission) parseHeader() int {
	version := t.parseN(3)
	t.versionSum += version
	return t.parseN(3)
}

func (t *transmission) parseLiteral() int {
	bits := ""
	for start := '1'; start == '1'; t.pos += 5 {
		start = rune(t.bits[t.pos])
		bits += t.bits[t.pos+1 : t.pos+5]

	}
	i, _ := strconv.ParseInt(bits, 2, 64)
	return int(i)
}

func (t *transmission) parseOperator() (subPackets []int) {
	lengthTypeId := t.parseN(1)
	switch lengthTypeId {
	case 0:
		// total length
		nBits := t.parseN(15)
		end := t.pos + nBits
		for t.pos < end {
			subPackets = append(subPackets, t.parsePacket())
		}
	case 1:
		// n subpackets
		nPackets := t.parseN(11)
		for i := 0; i < nPackets; i++ {
			subPackets = append(subPackets, t.parsePacket())
		}
	}
	return
}

func (t *transmission) parsePacket() int {
	typeId := t.parseHeader()

	if typeId == 4 {
		return t.parseLiteral()
	}

	subPackets := t.parseOperator()

	return score(typeId, subPackets)
}

func score(typeId int, packets []int) int {
	switch typeId {
	case 0:
		sum := 0
		for _, p := range packets {
			sum += p
		}
		return sum
	case 1:
		product := 1
		for _, p := range packets {
			product *= p
		}
		return product
	case 2:
		return slices.Min(packets)
	case 3:
		return slices.Max(packets)
	case 5:
		if packets[0] > packets[1] {
			return 1
		}
	case 6:
		if packets[0] < packets[1] {
			return 1
		}
	case 7:
		if packets[0] == packets[1] {
			return 1
		}
	}
	return 0
}

func part1(lines []string) (res int) {
	t := &transmission{
		bits: toBits(lines[0]),
	}
	t.parsePacket()
	return t.versionSum
}

func part2(lines []string) (res int) {
	t := &transmission{
		bits: toBits(lines[0]),
	}
	return t.parsePacket()
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
