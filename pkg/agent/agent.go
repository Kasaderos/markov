package agent

import "github.com/kasaderos/markov/pkg/env"

type Matrix [][]int

func (m *Matrix) Init(nrow int, ncol int) {
	a := make([][]int, 0, nrow)
	for i := 0; i < nrow; i++ {
		a = append(a, make([]int, ncol))
	}
	*m = Matrix(a)
}

type Agent struct {
	env *env.Env
	Q   Matrix
}

func (a *Agent) Explore() {
	// get Q
	//
}

func (a *Agent) Trade() {
}

func (a *Agent) ClosePositions() {

}
