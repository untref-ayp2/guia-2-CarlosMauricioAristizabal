package ejercicios

import (
	"guia2/queue"
	"guia2/stack"
)

func InvertirCadena(cadena string) string {
	var pila stack.Stack
	var cadenaInvertida string
	for _, valor := range cadena {
		pila.Push(string(valor))
	}
	for !pila.IsEmpty() {
		valor, _ := pila.Pop()
		cadenaInvertida += valor.(string)
	}
	return cadenaInvertida
}

func Palindromo(texto string) bool {
	textoInvertido := InvertirCadena(texto)

	return texto == textoInvertido
}

func Balanceada(cadena string) bool {
	balanceada := true
	var pila stack.Stack
	i := 0
	for i < len(cadena) && balanceada {
		llave := string(cadena[i])

		if (llave == "{") || (llave == "(") || (llave == "[") {
			pila.Push(llave)

		} else {
			value, _ := pila.Top()
			switch {
			case (llave == string(")")) && (value == string("(")):
				pila.Pop()
			case (llave == string("}")) && (value == string("{")):
				pila.Pop()
			case (llave == string("]")) && (value == string("[")):
				pila.Pop()
			default:
				balanceada = false
			}
		}
		i++
	}
	return balanceada
}

func UnirColas(q1, q2 queue.Queue) queue.Queue {
	for !q2.IsEmpty() {
		valor, _ := q2.Dequeue()
		q1.Enqueue(valor)
	}
	return q1
}
