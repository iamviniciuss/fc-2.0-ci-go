package main

import (
	"testing"
)

func TestSoma(t *testing.T) {
	total := Soma(2, 3)

	if total != 5 {
		t.Errorf("O resultado esperado era %d e o recebido foi %d", 5, total)
	}
}
