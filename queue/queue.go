package queue

import (
	"errors"
	"guia2/stack"
)

// Queue implementa una cola genérica sobre un arreglo dinámico.
type Queue struct {
	cola []any
}

// Enqueue agrega un elemento a la cola. O(1)
func (q *Queue) Enqueue(v any) {
	q.cola = append(q.cola, v)
}

// Dequeue saca un elemento de la cola. O(1)
func (q *Queue) Dequeue() (any, error) {
	if len(q.cola) == 0 {
		return nil, errors.New("queue is empty")
	}
	head := (q.cola)[0]
	q.cola = (q.cola)[1:]
	return head, nil
}

// Front devuelve el elemento del frente de la cola. O(1)
func (q *Queue) Front() (any, error) {
	if q.IsEmpty() {
		return nil, errors.New("queue is empty")
	}
	return q.cola[0], nil
}

// IsEmpty verifica si la cola esta vacia. O(1)
func (q *Queue) IsEmpty() bool {
	return len(q.cola) == 0
}

type QueueS struct {
	pila stack.Stack
}

func (q *QueueS) Enqueue(v any) {
	//TODO
}

func (q *QueueS) Dequeue() (any, error) {
	//TODO
}

func (q *QueueS) IsEmpty() bool {

}

func (q *QueueS) Front() (any, error) {
	//TODO
}
