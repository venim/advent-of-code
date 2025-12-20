package main

import (
	_ "embed"
	"flag"
	"math"
	"slices"
	"strings"
	"time"

	"github.com/golang/glog"
	"github.com/venim/advent-of-code/util"
)

type Machine struct {
	Indicators []int
	Goal       []int
	Buttons    []Button
	Joltage    []int
}

type Button struct {
	Indicators []int
}

func NewMachine(input string) *Machine {
	m := new(Machine)
	parts := strings.Fields(input)
	indicators := parts[0]
	joltage := parts[len(parts)-1]
	for i := 1; i < len(indicators)-1; i++ {
		switch indicators[i] {
		case '.':
			m.Goal = append(m.Goal, 0)
		case '#':
			m.Goal = append(m.Goal, 1)
		}
	}
	m.Indicators = make([]int, len(m.Goal))
	js := strings.SplitSeq(joltage[1:len(joltage)-1], ",")
	for j := range js {
		m.Joltage = append(m.Joltage, util.MustAtoi(j))
	}
	for i := 1; i < len(parts)-1; i++ {
		ns := strings.SplitSeq(parts[i][1:len(parts[i])-1], ",")
		b := Button{}
		for n := range ns {
			b.Indicators = append(b.Indicators, util.MustAtoi(n))
		}
		m.Buttons = append(m.Buttons, b)
	}

	return m
}

type state struct {
	Indicators []int
	Steps      int
}

// Less makes state compatible with util.MinHeap.
func (s state) Less(other state) bool {
	return s.Steps < other.Steps
}

func (m Machine) solve() int {
	q := []state{{Indicators: m.Indicators, Steps: 0}}
	visited := make(map[string]struct{})

	// Helper to create a string key from the indicator slice for the visited map
	key := func(indicators []int) string {
		var sb strings.Builder
		for _, i := range indicators {
			sb.WriteRune(rune('0' + i))
		}
		return sb.String()
	}

	visited[key(m.Indicators)] = struct{}{}

	for len(q) > 0 {
		current := q[0]
		q = q[1:]

		if slices.Equal(current.Indicators, m.Goal) {
			return current.Steps
		}

		for _, button := range m.Buttons {
			nextIndicators := make([]int, len(current.Indicators))
			copy(nextIndicators, current.Indicators)
			for _, i := range button.Indicators {
				nextIndicators[i] = 1 - nextIndicators[i] // Flip the indicator
			}

			if _, ok := visited[key(nextIndicators)]; !ok {
				visited[key(nextIndicators)] = struct{}{}
				q = append(q, state{Indicators: nextIndicators, Steps: current.Steps + 1})
			}
		}
	}
	return -1 // No solution found
}

const inf = 1<<31 - 1

func (m Machine) solveForJoltage() int {
	// Overview:
	// We have a system of linear equations. Each "counter" on the machine is a goal we must reach.
	// Each "button" adds 1 to a specific set of counters.
	// We want to find out how many times to press each button (let's call the number of presses for button i "x_i")
	// so that the total added to each counter matches its goal.
	// This gives us equations like:
	//   (Button 1's effect on Counter A)*x_1 + (Button 2's effect on Counter A)*x_2 + ... = Goal for Counter A
	//
	// We want to find the solution that uses the *fewest total button presses* (minimize sum of x_i).
	// Also, x_i must be non-negative integers (can't press a button -1 or 1.5 times).

	numVars := len(m.Buttons) // The variables x_0, x_1, ...
	numEqs := len(m.Joltage)  // The equations (one per counter)

	// Step 1: Establish Limits
	// We can't press a button more times than the goal itself. If a button adds 1 to a counter with goal 5,
	// pressing it 6 times would overshoot. We find the tightest limit for each button.
	bounds := make([]int, numVars)
	for i := 0; i < numVars; i++ {
		minGoal := inf
		affected := false
		for _, ind := range m.Buttons[i].Indicators {
			affected = true
			if m.Joltage[ind] < minGoal {
				minGoal = m.Joltage[ind]
			}
		}
		if !affected {
			bounds[i] = 0 // Button affects nothing useful.
		} else {
			bounds[i] = minGoal
		}
	}

	// Step 2: Build the Matrix
	// We organize our equations into a grid (matrix).
	// Rows = equations (counters). Columns = variables (buttons).
	// We add an extra column at the end for the "Right Hand Side" (RHS) which holds the goal values.
	// Example: 1*x_0 + 0*x_1 = 5 becomes row [1, 0 | 5]
	mat := make([][]float64, numEqs)
	for r := 0; r < numEqs; r++ {
		mat[r] = make([]float64, numVars+1)
		// Set coefficients (1 if button affects counter, 0 otherwise)
		for c := 0; c < numVars; c++ {
			affects := false
			for _, ind := range m.Buttons[c].Indicators {
				if ind == r {
					affects = true
					break
				}
			}
			if affects {
				mat[r][c] = 1.0
			}
		}
		// Set the goal value for this counter
		mat[r][numVars] = float64(m.Joltage[r])
	}

	// Step 3: Gaussian Elimination
	// This is a systematic way to simplify the equations.
	// We try to isolate one variable per equation.
	// For example, turn:
	//   x + y = 5
	//   x - y = 1
	// Into:
	//   x = 3
	//   y = 2
	pivotRow := 0
	pivots := make(map[int]int) // Maps a variable column to the row where it is solved (pivoted)
	pivotCols := make([]int, 0) // Keeps track of the order we solved variables

	for col := 0; col < numVars && pivotRow < numEqs; col++ {
		// Find the best equation (row) to solve for this variable (col).
		// We pick the one with the largest coefficient to avoid numerical issues (though here they are all 0 or 1).
		sel := pivotRow
		for r := pivotRow + 1; r < numEqs; r++ {
			if math.Abs(mat[r][col]) > math.Abs(mat[sel][col]) {
				sel = r
			}
		}
		// If the coefficient is 0, this variable doesn't appear in the remaining equations. Skip it.
		if math.Abs(mat[sel][col]) < 1e-9 {
			continue
		}

		// Swap the selected row to the top (pivotRow)
		mat[pivotRow], mat[sel] = mat[sel], mat[pivotRow]

		// Normalize the row so the coefficient for our variable is 1.
		div := mat[pivotRow][col]
		for c := col; c <= numVars; c++ {
			mat[pivotRow][c] /= div
		}

		// Eliminate this variable from all other equations.
		// If we know x = 3, we replace 'x' with '3' in other equations to simplify them.
		for r := 0; r < numEqs; r++ {
			if r != pivotRow {
				factor := mat[r][col]
				for c := col; c <= numVars; c++ {
					mat[r][c] -= factor * mat[pivotRow][c]
				}
			}
		}

		// Record that we found a pivot for this column
		pivots[col] = pivotRow
		pivotCols = append(pivotCols, col)
		pivotRow++
	}

	// Step 4: Consistency Check
	// If we end up with a row like [0, 0, ... | 5], it means 0 = 5, which is impossible.
	// This happens if the goals are contradictory.
	for r := pivotRow; r < numEqs; r++ {
		if math.Abs(mat[r][numVars]) > 1e-5 {
			return inf // No solution exists
		}
	}

	// Step 5: Identify "Free" Variables
	// Variables that we found a pivot for are "Dependent" (their value is fixed by the equations).
	// Variables we didn't find a pivot for are "Free".
	// This means the system has multiple solutions. We can pick *any* value for the free variables,
	// and the dependent variables will adjust to match.
	var freeVars []int
	for c := 0; c < numVars; c++ {
		if _, ok := pivots[c]; !ok {
			freeVars = append(freeVars, c)
		}
	}

	minTotal := inf

	// Step 6: Search for the Best Solution
	// Since we want integer solutions and minimum total presses, we can't just pick any value.
	// We iterate through all valid integer values for the Free Variables.
	// Because the problem space is small, this is fast.
	var search func(idx int, currentAssignment []int)
	currentAssignment := make([]int, numVars)

	search = func(idx int, assignment []int) {
		// Base Case: All free variables have been assigned a value.
		if idx == len(freeVars) {
			valid := true
			currentSum := 0

			// Sum up the presses for the free variables
			for _, fv := range freeVars {
				currentSum += assignment[fv]
			}

			// Calculate the Dependent Variables based on the chosen Free Variables.
			// The equation for a pivot variable x_p is effectively:
			// x_p + (coeff * x_free1) + ... = RHS
			// So: x_p = RHS - (coeff * x_free1) - ...
			for _, pc := range pivotCols {
				r := pivots[pc]
				val := mat[r][numVars]
				for _, fv := range freeVars {
					val -= mat[r][fv] * float64(assignment[fv])
				}

				// Check if the resulting value is a valid non-negative integer.
				intVal := int(math.Round(val))
				if math.Abs(val-float64(intVal)) > 1e-5 && math.Abs(val-float64(intVal)) < 0.99999 {
					valid = false // Not an integer
					break
				}
				if intVal < 0 || intVal > bounds[pc] {
					valid = false // Negative or exceeds bounds
					break
				}
				assignment[pc] = intVal
				currentSum += intVal
			}

			// If valid, check if this is the best (lowest) total so far.
			if valid {
				if currentSum < minTotal {
					minTotal = currentSum
				}
			}
			return
		}

		// Recursive Step: Try every possible value for the current free variable.
		fv := freeVars[idx]
		for val := 0; val <= bounds[fv]; val++ {
			assignment[fv] = val
			search(idx+1, assignment)
		}
	}

	search(0, currentAssignment)

	return minTotal
}

var (
	//go:embed in.txt
	input string
)

func part1(lines []string) (res int) {
	for _, line := range lines {
		m := NewMachine(line)
		res += m.solve()
	}
	return
}

func part2(lines []string) (res int) {
	for _, line := range lines {
		m := NewMachine(line)
		// The recursive solution returns `inf` if no solution is found.
		if r := m.solveForJoltage(); r != inf {
			res += r
		}
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
