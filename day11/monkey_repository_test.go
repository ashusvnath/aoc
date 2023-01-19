package main

import "testing"

func TestRepository(t *testing.T) {
	t.Run("Repository add and get", func(t *testing.T) {
		m0 := NewMonkeyBuilder().Id("0").Build()
		repo.Add(m0)
		m1 := NewMonkeyBuilder().Id("1").Build()
		repo.Add(m1)
		assertEqual("0", repo.Get("0").id, t)
		assertEqual("1", repo.Get("1").id, t)
	})

}
