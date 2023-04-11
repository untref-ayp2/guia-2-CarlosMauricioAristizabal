package ejercicios

import (
	"guia2/queue"
	"guia2/stack"
)

func InvertirCadena(cadena string) string {
	//TODO
	var s stack.Stack
	var nuevaCadena string
	for _, v := range cadena {
		s.Push(string(v))
	}

	c, err := s.Pop()
	for err == nil {

		c, err = s.Pop()
	}
	return nuevaCadena
}

func Palindromo(string) bool {
	//TODO
	bool := true
	return bool
}

func Balanceada(string) bool {
	//TODO
	bool := true
	return bool
}

func UnirColas(q1, q2 queue.Queue) queue.Queue {
	// TODO
	return queue.Queue{}
}
