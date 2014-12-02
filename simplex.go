//   Full credit to http://algs4.cs.princeton.edu/65reductions/Simplex.java.html from
//   http://algs4.cs.princeton.edu/65reductions/ by Robert Sedgewick and Kevin Wayne.
//
//   Given an M-by-N matrix A, an M-length vector b, and an
//   N-length vector c, solve the  LP { max cx : Ax <= b, x >= 0 }.
//   Assumes that b >= 0 so that x = 0 is a basic feasible solution.
//
//   Creates an (M+1)-by-(N+M+1) simplex tableaux with the
//   RHS in column M+N, the objective function in row M, and
//   slack variables in columns M through M+N-1.

package solver

import (
	"fmt"
)

const EPSILON float64 = 0.0000000001

var a [][]float64
var M int
var N int
var basis []int

//Entry point. Creates matrices based on AllPlayers array of arrays,
//executes the solver, performs post-processing adjustments,
//and returns roster.
func CreateSimplexRoster() []Player {
	//Create matrices needed for CreateSimplexTableaux
	//Matrices are A: a 2-D array where the first row is player salaries. 
	//Subsequent rows are filled with "1" or "0" where a "1" indicates that 
	//the player in that position of the array plays the position held by that row.
	//For example, a QB in the third position of SingleList would have his salary
	//in the third element of the first row, a "1" in the third element of the second row, 
	//and "0"s in all other rows.
	//Position rows are in order of QB, RB, WR, TE, K, and D.
	A := make([][]float64, 7)
	for i := range A {
		A[i] = make([]float64, len(SingleList))
	}

	//Matrix b is a 1-D array of constraints where the first element is the total 
	//allowable salary for a roster, and other elements indicate the number of 
	//players at the indicated position on the roster. Position of the element
	//matches the position rows found in matrix A.
	b := make([]float64, 7)

	//Matrix c is a 1-D array of projected points per player.
	c := make([]float64, len(SingleList))

	//c := []float64{13.0, 23.0}
	//b := []float64{480.0, 160.0, 1190.0}

	//A := [][]float64{{5.0, 15.0}, {4.0, 4.0}, {35.0, 20.0},}
	/*
	A := make([][]float64, 3)
	for i := range A {
		A[i] = make([]float64, 2)
	}
	A[0][0] = 5.0
	A[0][1] = 15.0
	A[1][0] = 4.0
	A[1][1] = 4.0
	A[2][0] = 35.0
	A[2][1] = 20.0
*/
	//Call CreateSimplexTableaux
	createSimplexTableaux(A, b, c)

	//Call Solve, returns nil if there was an error or unbounded solution
	if !solve() {
		return nil
	}

	//Call Primal to get decision variable vector
	fmt.Printf("Result = %v\n", primal())

	//Check vector for anomalies like picking the same player twice

	//Adjust roster as needed

	//create roster of Player structs

	return nil
}

// sets up the simplex tableaux and
func createSimplexTableaux(A [][]float64, b []float64, c []float64) {
	M = len(b)
	N = len(c)
	//var a [M + 1][N + M + 1]float64
	a = make([][]float64, M+1)
	for i := range a {
		a[i] = make([]float64, N+M+1)
	}

	for i := 0; i < M; i++ {
		for j := 0; j < N; j++ {
			a[i][j] = A[i][j]
		}
	}

	for i := 0; i < M; i++ {
		a[i][N+i] = 1.0
	}
	for j := 0; j < N; j++ {
		a[M][j] = c[j]
	}
	for i := 0; i < M; i++ {
		a[i][M+N] = b[i]
	}

	basis = make([]int, M)

	for i := 0; i < M; i++ {
		basis[i] = N + i
	}
}

// print tableaux.
// This is not needed during normal operation, only used for testing
func show() {
	fmt.Printf("M = %v\n", M)
	fmt.Printf("N = %v\n", N)
	for i := 0; i <= M; i++ {
		for j := 0; j <= M+N; j++ {
			fmt.Printf("%v ", a[i][j])
		}
		fmt.Printf(" %v\n", 10)
	}
	fmt.Printf("value = %v\n ", -a[M][M+N])
	for i := 0; i < M; i++ {
		if basis[i] < N {
			fmt.Printf("x_%v = %v ", basis[i], a[i][M+N])
		}
	}
	fmt.Printf(" %v\n", 10)
}

//returns true if successful finding a solution, false if not
func solve() bool {
	var result bool = false
	for true {
		// find entering column q
		q := bland()
		if q == -1 {
			result = true
			break // optimal
		}

		// find leaving row p
		p := minRatioRule(q)
		if p == -1 {
			result = false
			break //unbounded
		}

		// pivot
		pivot(p, q)

		// update basis
		basis[p] = q

	}

	return result
}

// lowest index of a non-basic column with a positive cost
func bland() int {
	for j := 0; j < M+N; j++ {
		if a[M][j] > 0 {
			return j
		}
	}
	return -1 // optimal
}

// index of a non-basic column with most positive cost
func dantzig() int {
	q := 0
	for j := 1; j < M+N; j++ {
		if a[M][j] > a[M][q] {
			q = j
		}
	}

	if a[M][q] <= 0 {
		return -1 // optimal
	} else {
		return q
	}
}

// find row p using min ratio rule (-1 if no such row)
func minRatioRule(q int) int {
	p := -1
	for i := 0; i < M; i++ {
		if a[i][q] <= 0 {
			continue
		} else if p == -1 {
			p = i
		} else if (a[i][M+N] / a[i][q]) < (a[p][M+N] / a[p][q]) {
			p = i
		}
	}
	return p
}

// pivot on entry (p, q) using Gauss-Jordan elimination
func pivot(p int, q int) {

	// everything but row p and column q
	for i := 0; i <= M; i++ {
		for j := 0; j <= M+N; j++ {
			if i != p && j != q {
				a[i][j] -= a[p][j] * a[i][q] / a[p][q]
			}
		}
	}

	// zero out column q
	for i := 0; i <= M; i++ {
		if i != p {
			a[i][q] = 0.0
		}
	}

	// scale row p
	for j := 0; j <= M+N; j++ {
		if j != q {
			a[p][j] /= a[p][q]
		}
	}

	a[p][q] = 1.0
}

// return optimal objective value
func value() float64 {
	return -a[M][M+N]
}

// return primal solution vector
func primal() []float64 {
	x := make([]float64, N)
	for i := 0; i < M; i++ {
		if basis[i] < N {
			x[basis[i]] = a[i][M+N]
		}
	}

	return x
}
