package main

import "fmt"

var max int = 10000

type Queue struct {
	data []int
}

type QueueI interface {
	isEmpty() bool
	isFull() bool
	enqueue(int)
	dequeue() int
}

var _ QueueI = &Queue{}

func (q *Queue) isEmpty() bool {
	return len(q.data) == 0
}

func (q *Queue) isFull() bool {
	return len(q.data) == max
}

func (q *Queue) enqueue(n int) {
	if q.isFull() {
		fmt.Println("error: queue is full")
		return
	}
	q.data = append(q.data, n)
}

func (q *Queue) dequeue() int {
	if q.isEmpty() {
		fmt.Println("error: queue is empty")
		return 0
	}
	dequeueData := q.data[0]
	q.data = q.data[1:]
	return dequeueData
}

func main() {
	var q *Queue = &Queue{}

	q.enqueue(3) // [] -> [3]
	q.enqueue(5) // [3] -> [3, 5]
	q.enqueue(7) // [3, 5] -> [3, 5, 7]

	fmt.Println(q.dequeue()) // [3, 5, 7] -> [5, 7] で 3 を出力
	fmt.Println(q.dequeue()) // [5, 7] -> [5] で 5 を出力

	q.enqueue(9)  // [7] -> [7, 9]
	q.enqueue(11) // [7, 9] -> [7, 9, 11]

	q.dequeue() // [7, 9, 11] -> [9, 11]
	q.dequeue() // [9, 11] -> [11]
	q.dequeue() // [11] -> []

	if q.isEmpty() {
		fmt.Println("empty")
	} else {
		fmt.Println("not empty")
	}
}
