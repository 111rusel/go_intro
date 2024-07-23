package animals

import "testing"

func TestAnimals(t *testing.T) {
	t.Run("CatVoice", func(t *testing.T) {
		want := "Meow"
		c := Cat{
			Voice: "Meow",
		}
		got := c.Say()
		if want != got {
			t.Errorf("want %s got %s", want, got)
		}
	})

	t.Run("CatGo", func(t *testing.T) {
		want := "Run"
		c := Cat{
			Voice: "Meow",
			Move:  "Run",
		}
		got := c.Go()
		if want != got {
			t.Errorf("want %s got %s", want, got)
		}
	})

	t.Run("RabbitVoice", func(t *testing.T) {
		want := "uwu"
		r := Rabbit{
			Voice: "uwu",
		}
		got := r.Say()
		if want != got {
			t.Errorf("want %s got %s", want, got)
		}
	})

	t.Run("RabbitGo", func(t *testing.T) {
		want := "Jump"
		r := Rabbit{
			Voice: "uwu",
			Move:  "Jump",
		}
		got := r.Go()
		if want != got {
			t.Errorf("want %s got %s", want, got)
		}
	})

	t.Run("CallCat", func(t *testing.T) {
		want := "Meow"
		c := Cat{
			Voice: "Meow",
			Move:  "Run",
		}
		got := Call(c)
		if want != got {
			t.Errorf("want %s got %s", want, got)
		}
	})

	t.Run("CallRabbit", func(t *testing.T) {
		want := "uwu"
		r := Rabbit{
			Voice: "uwu",
			Move:  "Jump",
		}
		got := Call(r)
		if want != got {
			t.Errorf("want %s got %s", want, got)
		}
	})

	t.Run("BeatCat", func(t *testing.T) {
		want := "Run"
		c := Cat{
			Voice: "Meow",
			Move:  "Run",
		}
		got := Beat(c)
		if want != got {
			t.Errorf("want %s got %s", want, got)
		}
	})

	t.Run("BeatRabbit", func(t *testing.T) {
		want := "Jump"
		r := Rabbit{
			Voice: "uwu",
			Move:  "Jump",
		}
		got := Beat(r)
		if want != got {
			t.Errorf("want %s got %s", want, got)
		}
	})
}
