package main

import (
	"math/big"
	"testing"
)

func TestRepository(t *testing.T) {
	t.Run("Repository add and get", func(t *testing.T) {
		repo := GetMonkeyRepository()
		m0 := NewMonkeyBuilder(Identity()).Id("0").Divisor(big.NewInt(1)).Build()
		repo.Add(m0)
		m1 := NewMonkeyBuilder(Identity()).Id("1").Divisor(big.NewInt(1)).Build()
		repo.Add(m1)
		assertEqual("0", repo.Get("0").id, t)
		assertEqual("1", repo.Get("1").id, t)
	})

}
