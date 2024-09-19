package main

import (
	"bytes"
	"reflect"
	"testing"
)

type FakeSleeper struct {
	Calls int
}

func (s *FakeSleeper) Sleep() {
	s.Calls++
}

type MockCountDawnOperations struct {
	Calls []string
}

func (m *MockCountDawnOperations) Sleep() {
	m.Calls = append(m.Calls, "Sleep")
}

func (m *MockCountDawnOperations) Write(p []byte) (n int, err error) {
	m.Calls = append(m.Calls, "Write")
	return n, err
}

func TestMain(t *testing.T) {
	t.Run("Test Mocking", func(t *testing.T) {
		buffer := &bytes.Buffer{}
		s := &FakeSleeper{}
		CountDown(buffer, s)
		got := buffer.String()
		want := `3
2
1
Go!`

		if got != want {
			t.Errorf("want %q got %q", want, got)
		}
		if s.Calls != 3 {
			t.Errorf("want 3 sleeps got %d", s.Calls)
		}
	})

	t.Run("Sleep before every print", func(t *testing.T) {
		s := &MockCountDawnOperations{}
		CountDown(s, s)

		want := []string{
			"Write",
			"Sleep",
			"Write",
			"Sleep",
			"Write",
			"Sleep",
			"Write",
		}
		got := s.Calls

		if !reflect.DeepEqual(want, got) {
			t.Errorf("want %v got %v", want, got)
		}
	})
}
