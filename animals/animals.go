package animals

type Cat struct {
	Voice string
	Move  string
}

func (c Cat) Say() string {
	return c.Voice
}

func (c Cat) Go() string {
	return c.Move
}

type Rabbit struct {
	Voice string
	Move  string
}

func (r Rabbit) Say() string {
	return r.Voice
}

func (r Rabbit) Go() string {
	return r.Move
}

func Call(c Animal) string {
	return c.Say()
}

type Animal interface {
	Say() string
	Go() string
}

func Beat(a Animal) string {
	return a.Go()
}
