package tests

import (
	"guia2/ejercicios"
	"guia2/queue"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestInvertirCadena(t *testing.T) {
	salida := ejercicios.InvertirCadena("abcd")
	if "dcba" != salida {
		t.Error("InvertirCadena falla")
	}
}

func TestPalindromo(t *testing.T) {
	if true != ejercicios.Palindromo("1456541") ||
		true != ejercicios.Palindromo("145541") {
		t.Error("Palindromo falla")
	}
}

func TestBalanceada(t *testing.T) {
	if true != ejercicios.Balanceada("[()]{}{[()()]()}") ||
		false != ejercicios.Balanceada("[(])") {
		t.Error("Balanceada falla")
	}
}

func TestUnirColas(t *testing.T) {
	var q1 queue.Queue
	q1.Enqueue(1)
	q1.Enqueue(2)
	q1.Enqueue(3)

	var q2 queue.Queue
	q2.Enqueue(5)
	q2.Enqueue(7)
	var q12_esperado queue.Queue
	q12_esperado.Enqueue(1)
	q12_esperado.Enqueue(2)
	q12_esperado.Enqueue(3)
	q12_esperado.Enqueue(5)
	q12_esperado.Enqueue(7)

	q12_dado := ejercicios.UnirColas(q1, q2)
	if !cmp.Equal(q12_dado, q12_esperado) {
		t.Error("UnirColas falla")
	}
}
