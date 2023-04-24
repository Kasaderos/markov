package env

type Indicator int

type Env struct {
	I     []Indicator
	state int
}

func (e *Env) Update() {
	e.I = loadIndicators()
	e.state = 0
	for _, v := range e.I {
		e.state = e.state*10 + int(v)
	}
}

func (e *Env) State() int {
	return e.state
}

func loadIndicators() []Indicator {
	return nil
}
