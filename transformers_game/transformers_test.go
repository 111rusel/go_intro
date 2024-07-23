package transformersGame

import "testing"

func TestTransformers(t *testing.T) {
	t.Run("Optimus Prime Fire", func(t *testing.T) {
		want := 10
		op := Autobot{
			Name:   "Optimus Prime",
			Damage: 10,
		}
		got := op.Fire()
		if want != got {
			t.Errorf("want %d got %d", want, got)
		}
	})

	t.Run("Optimus Prime Take Damage", func(t *testing.T) {
		want := 99
		op := Autobot{
			Name:        "Optimus Prime",
			Damage:      10,
			HealthPoint: 100,
		}
		op.TakeDamage(1)
		if op.HealthPoint != want {
			t.Errorf("want %d got %d", want, op.HealthPoint)
		}
	})

	t.Run("Megatrom Fire", func(t *testing.T) {
		want := 15
		mg := Desepticon{
			Name:   "Megatron",
			Damage: 15,
		}
		got := mg.Fire()
		if want != got {
			t.Errorf("want %d got %d", want, got)
		}
	})

	t.Run("Megatron Take Damage", func(t *testing.T) {
		want := 94
		mg := Desepticon{
			Name:        "Megatron",
			Damage:      15,
			HealthPoint: 95,
		}
		mg.TakeDamage(1)
		if mg.HealthPoint != want {
			t.Errorf("want %d got %d", want, mg.HealthPoint)
		}
	})

	t.Run("Prepear Game", func(t *testing.T) {
		wantPlayer1 := "Optimus Prime"
		wantPlayer2 := "Megatron"
		op := Autobot{
			Name:        "Optimus Prime",
			Damage:      10,
			HealthPoint: 100,
		}
		mg := Desepticon{
			Name:        "Megatron",
			Damage:      15,
			HealthPoint: 95,
		}
		got := CreateGame(&op, &mg)
		if got.Player1.GetName() != wantPlayer1 {
			t.Errorf("want %s got %s", wantPlayer1, got.Player1.GetName())
		}
		if got.Player2.GetName() != wantPlayer2 {
			t.Errorf("want %s got %s", wantPlayer2, got.Player2.GetName())
		}
	})

	t.Run("Prepear Game", func(t *testing.T) {
		wantPlayer2 := "Optimus Prime"
		wantPlayer1 := "Megatron"
		op := Autobot{
			Name:        "Optimus Prime",
			Damage:      10,
			HealthPoint: 100,
		}
		mg := Desepticon{
			Name:        "Megatron",
			Damage:      15,
			HealthPoint: 95,
		}
		got := CreateGame(&mg, &op)
		if got.Player1.GetName() != wantPlayer1 {
			t.Errorf("want %s got %s", wantPlayer1, got.Player1.GetName())
		}
		if got.Player2.GetName() != wantPlayer2 {
			t.Errorf("want %s got %s", wantPlayer2, got.Player2.GetName())
		}
	})

	t.Run("Game Fight", func(t *testing.T) {
		wantHpPlayer1 := -5
		wantHpPlayer2 := 25
		op := Autobot{
			Name:        "Optimus Prime",
			Damage:      10,
			HealthPoint: 100,
		}
		mg := Desepticon{
			Name:        "Megatron",
			Damage:      15,
			HealthPoint: 95,
		}
		got := CreateGame(&op, &mg)
		got.Fight()
		if got.Player1.GetHP() != wantHpPlayer1 {
			t.Errorf("want %d got %d", wantHpPlayer1, got.Player1.GetHP())
		}
		if got.Player2.GetHP() != wantHpPlayer2 {
			t.Errorf("want %d got %d", wantHpPlayer2, got.Player2.GetHP())
		}

	})

	t.Run("Game Winner", func(t *testing.T) {
		want := "Megatron"
		op := Autobot{
			Name:        "Optimus Prime",
			Damage:      10,
			HealthPoint: 100,
		}
		mg := Desepticon{
			Name:        "Megatron",
			Damage:      15,
			HealthPoint: 95,
		}
		g := CreateGame(&op, &mg)
		g.Fight()
		got := g.GetWinner()
		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	})
}
