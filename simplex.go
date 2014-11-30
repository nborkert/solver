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

const EPSILON float64 = 0.0000000001

//var a [][]float64
//var M int
//var N int
//var basis []int

func CreateSimplexRoster() []Player {
	return nil
}

// sets up the simplex tableaux and 
func CreateSimplexTableaux(A [][]float64, b []float64, c []float64) {
	M := len(b)
	N := len(c)
	//var a [M + 1][N + M + 1]float64
	a := make([][]float64, M+1)
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

	basis := make([]int, M)

	for i := 0; i < M; i++ {
		basis[i] = N + i
	}

	Solve()

}

func Solve() []Player {

	return nil
}
