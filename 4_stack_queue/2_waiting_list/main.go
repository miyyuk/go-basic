// キャンセル待ちのお客様リストを想定します。
// お客様の情報には、名前、電話番号、メールアドレスを登録します。
// キャンセル待ちのお客様情報をリストに登録していきましょう。
// キャンセルが出たら、一番最初に登録されたお客様の情報をリストから削除し、出力しましょう。

package main

import (
	"fmt"
)

type Customer struct {
	name             string
	telephone_number string
	email_address    string
}

type Node struct {
	data *Customer
	prev *Node
	next *Node
}

type Queue struct {
	front *Node
	rear  *Node
}

type QueueI interface {
	isEmpty() bool
	Enqueue(Customer)
	Dequeue() *Node
}

var _ QueueI = &Queue{}

func (q *Queue) isEmpty() bool {
	return q.front == q.rear && q.front == nil
}

func (q *Queue) Enqueue(c Customer) {
	var node *Node = &Node{data: &c}
	if q.isEmpty() { // Queue が空のとき、 front も rear も node をさす
		q.front = node
		q.rear = node
	} else { // Queue が空ではないとき、
		q.rear.next = node // リストの最後尾に node をリンクする
		node.prev = q.rear // node の　prev はそれまでリスト最後尾だった Node になる
		q.rear = node      // リストの最後尾を node にする
	}
}

func (q *Queue) Dequeue() *Node {
	if q.isEmpty() {
		fmt.Println("No one is waiting for cancellation")
		return nil
	}

	node := q.front        // 取り除く Node
	q.front = q.front.next // リストの先頭の Node は先頭の次の Node になる
	if q.front != nil {    // リストが空ではない
		q.front.prev = nil
	} else { // リストが空
		q.rear = nil
	}

	fmt.Printf("\nSomeone canceled!! %s can now be booked.\n", node.data.name)
	return node
}

// キャンセル待ちリストの出力
func (q *Queue) Traverse() {
	node := q.front
	for i := 1; ; i++ {
		if node == nil {
			break
		} else {
			fmt.Printf("%d: %v\n", i, *node.data)
			node = node.next
		}
	}
}

func main() {
	var q *Queue = &Queue{} // Queue を初期化

	var c1 *Customer = &Customer{name: "Sakura", telephone_number: "090-1234-5678", email_address: "sakura@example.com"}
	fmt.Printf("\n%s is waiting for cancellation\n", c1.name)
	q.Enqueue(*c1) // キャンセル待ちリストに追加
	q.Traverse()

	var c2 *Customer = &Customer{name: "Kaede", telephone_number: "080-1234-5678", email_address: "kaede@example.com"}
	fmt.Printf("\n%s is waiting for cancellation\n", c2.name)
	q.Enqueue(*c2)
	q.Traverse()

	q.Dequeue() // キャンセルが出たので、キャンセル待ちリストの先頭を Dequeue する
	q.Traverse()

	var c3 *Customer = &Customer{name: "Yuzu", telephone_number: "070-1234-5678", email_address: "yuzu@example.com"}
	fmt.Printf("\n%s is waiting for cancellation\n", c3.name)
	q.Enqueue(*c3)
	q.Traverse()
}

// ############# result #############

// Sakura is waiting for cancellation
// 1: {Sakura 090-1234-5678 sakura@example.com}

// Kaede is waiting for cancellation
// 1: {Sakura 090-1234-5678 sakura@example.com}
// 2: {Kaede 080-1234-5678 kaede@example.com}

// Someone canceled!! Sakura can now be booked.
// 1: {Kaede 080-1234-5678 kaede@example.com}

// Yuzu is waiting for cancellation
// 1: {Kaede 080-1234-5678 kaede@example.com}
// 2: {Yuzu 070-1234-5678 yuzu@example.com}
