package main

import (
	_ "embed"
	"flag"
	"regexp"
	"strings"
	"time"

	"github.com/golang/glog"
	"github.com/venim/advent-of-code/util"
)

var (
	//go:embed in.txt
	input string
)

// parseNumbers extracts all integer numbers from a line using strings.Fields.
func parseNumbers(line string) (nums []int) {
	for f := range strings.FieldsSeq(line) {
		nums = append(nums, util.MustAtoi(f))
	}
	return
}

// operRe is a regular expression to find operators (+ or *) and optional trailing spaces.
var operRe = regexp.MustCompile(`([\+\*])\s*`)

// parseOperands extracts all operator symbols from a line.
func parseOperands(line string) (opers []string) {
	matches := operRe.FindAllStringSubmatch(line, -1)
	for _, m := range matches {
		opers = append(opers, m[1])
	}
	return
}

// calculate performs the arithmetic reduction (sum or product) on a slice of numbers.
func calculate(nums []int, op string) int {
	if len(nums) == 0 {
		return 0
	}
	// Initialize the result with the first number.
	n := nums[0]
	// Apply the operation iteratively to the rest of the numbers.
	for _, val := range nums[1:] {
		switch op {
		case "*":
			n *= val
		case "+":
			n += val
		}
	}
	return n
}

func part1(lines []string) (res int) {
	if len(lines) == 0 {
		return 0
	}

	// Separate data lines (numbers) from the operand line.
	dataLines := lines[:len(lines)-1]
	opLine := lines[len(lines)-1]

	// Parse operands from the last line.
	operands := parseOperands(opLine)

	// Parse all numbers into a 2D grid.
	var grid [][]int
	for _, line := range dataLines {
		grid = append(grid, parseNumbers(line))
	}

	// Iterate through each column (corresponding to an operand) and calculate its result.
	for i, op := range operands {
		// Ensure valid column index.
		if len(grid) == 0 || i >= len(grid[0]) {
			continue
		}
		// Collect numbers vertically for the current column.
		var nums []int
		for _, row := range grid {
			if i < len(row) {
				nums = append(nums, row[i])
			}
		}
		// Add the calculated column result to the total.
		res += calculate(nums, op)
	}
	return
}

func part2(lines []string) (res int) {
	if len(lines) == 0 {
		return 0
	}

	// Separate data lines (numbers) from the operand line.
	dataLines := lines[:len(lines)-1]
	opLine := lines[len(lines)-1]

	// Find all operator column boundaries in the last line.
	matches := operRe.FindAllStringIndex(opLine, -1)

	// Process each identified column.
	for _, loc := range matches {
		start, end := loc[0], loc[1]
		// Extract the operator.
		op := opLine[start : start+1]

		// Determine the effective width of the number chunk for this column.
		chunkWidth := end - start
		// Adjust chunk width if it's not the last column (to account for trailing space).
		if end != len(opLine) {
			chunkWidth--
		}

		// Extract numbers vertically for the current column chunk.
		var nums []int
		for c := 0; c < chunkWidth; c++ {
			var val int
			var hasDigit bool
			colIdx := start + c
			// Iterate through data lines to build the number from characters in the current vertical slice.
			for _, line := range dataLines {
				if colIdx < len(line) {
					b := line[colIdx]
					// Accumulate digits to form the integer value.
					if b >= '0' && b <= '9' {
						val = val*10 + int(b-'0')
						hasDigit = true
					}
				}
			}
			// Only append if at least one digit was found in the column slice.
			if hasDigit {
				nums = append(nums, val)
			}
		}

		// Add the calculated column result to the total.
		res += calculate(nums, op)
	}

	return
}

func init() {
	_ = flag.Set("logtostderr", "true")
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
