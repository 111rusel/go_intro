package transformersGame

type Autobot struct {
	Name        string
	Damage      int
	HealthPoint int
}

func (a *Autobot) Fire() int {
	return a.Damage
}

func (a *Autobot) TakeDamage(damage int) {
	a.HealthPoint = a.HealthPoint - damage
}

func (a *Autobot) GetName() string {
	return a.Name
}

func (a *Autobot) GetHP() int {
	return a.HealthPoint
}

type Desepticon struct {
	Name        string
	Damage      int
	HealthPoint int
}

func (d *Desepticon) Fire() int {
	return d.Damage
}

func (d *Desepticon) TakeDamage(damage int) {
	d.HealthPoint = d.HealthPoint - damage
}

func (d *Desepticon) GetName() string {
	return d.Name
}

func (d *Desepticon) GetHP() int {
	return d.HealthPoint
}

type Game struct {
	Player1 Transformer
	Player2 Transformer
}

func (g *Game) GetWinner() string {
	if g.Player1.GetHP() <= 0 {
		return g.Player2.GetName()
	} else {
		return g.Player1.GetName()
	}

}

func (g *Game) Fight() {
	for {
		damage := g.Player1.Fire()
		g.Player2.TakeDamage(damage)
		if g.Player2.GetHP() <= 0 {
			break
		}
		damage = g.Player2.Fire()
		g.Player1.TakeDamage(damage)
		if g.Player1.GetHP() <= 0 {
			break
		}
	}

}

type Transformer interface {
	Fire() int
	TakeDamage(damage int)
	GetName() string
	GetHP() int
}

func CreateGame(player1, player2 Transformer) Game {
	return Game{
		Player1: player1,
		Player2: player2,
	}
}
