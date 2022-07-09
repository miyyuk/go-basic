// 名前、電話番号、メールアドレスが記載されたアドレス帳を想定します。

// Singly Linked Listでアドレス帳を作りましょう
// Listの最後尾にアドレスを追加するfunctionを作りましょう
// Listにあるアドレスを表示するfunctionを作りましょう
// Listからアドレスを削除するfunctionを作りましょう

package main

import "fmt"

type Address struct {
	name          string
	phone_number  string
	email_address string
}

type Node struct {
	addr *Address
	next *Node
}

type AddressList struct {
	root *Node
	len  int
}

type AddressListI interface {
	Insert(cur *Node, name, phone_number, email_address string) *Node
	GetAddress(n int) *Node
	Remove(n int, cur *Node) (*Node, error)
	Traverse()
}

var _ AddressListI = &AddressList{}

// ListにAddressを追加
func (a *AddressList) Insert(cur *Node, name, phone_number, email_address string) *Node {
	address := &Address{
		name:          name,
		phone_number:  phone_number,
		email_address: email_address,
	}

	var node *Node = &Node{addr: address}
	// 図① 参照
	if a.len != 0 {
		node.next = cur.next // Insert する Node の next を cur の next にする
		cur.next = node      // cur の next を Insert する Node にする
		cur = cur.next       // cur を Insert した Node にする
	} else { // Listが空のときは、 root を Insert したい Node にすればよい
		a.root = node
		cur = a.root
	}

	a.len++
	return cur
}

// n 番目の Node を取得
func (a *AddressList) GetAddress(n int) *Node {
	if n > a.len { // Node が存在しないので、 nil を return
		return nil
	}

	ptr := a.root
	for i := 1; i < n; i++ { // n 番目の Node まで、 root から辿っていく
		ptr = ptr.next
	}
	return ptr
}

// n 番目の Node を List から削除
func (a *AddressList) Remove(n int, cur *Node) (*Node, error) {
	if a.len == 0 {
		return cur, fmt.Errorf("List is empty")
	}

	target := a.GetAddress(n) // 削除する Node を取得
	if target == nil {
		return cur, fmt.Errorf("Address not found")
	}

	// 図② 参照
	if n > 1 {
		prev := a.GetAddress(n - 1) // 削除する Node の前の Node を取得
		prev.next = target.next     // 削除する Node の前の Node の next を、削除する Node の次の Node にする
		if n == a.len {             // 最後の Node を削除する場合、 cur を最後の Node の前の Node にする
			cur = prev
		}
	} else {
		a.root = target.next // 最初の Node を削除する場合、 root を削除する Node の次の Nodeに変更する
	}

	a.len--
	return cur, nil
}

// List の要素を全て表示
func (a *AddressList) Traverse() {
	if a.len == 0 {
		fmt.Println("List is empty")
	}

	fmt.Printf("Size of AddressList: %d\n", a.len)
	ptr := a.root
	for i := 1; i <= a.len; i++ {
		fmt.Printf("%d name: %s, phone_number: %s, email_address: %s\n", i, ptr.addr.name, ptr.addr.phone_number, ptr.addr.email_address)
		ptr = ptr.next
	}
}

func main() {
	var addressList *AddressList = &AddressList{} // AddressList を初期化
	fmt.Println("\n~~Init AddressList~~")
	addressList.Traverse()

	var cur *Node = &Node{} // current Node を初期化
	cur = addressList.Insert(cur, "Sakura", "090-1234-5678", "sakura@example.com")
	cur = addressList.Insert(cur, "Kaede", "080-1234-5678", "kaede@example.com")
	cur = addressList.Insert(cur, "Yuzu", "070-1234-5678", "yuzu@example.com")
	fmt.Println("\n~~After Insert Addresses~~")
	addressList.Traverse()

	cur, err := addressList.Remove(2, cur)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println("\n~~After Remove Address~~")
	addressList.Traverse()
}

// ############# result #############
// ~~Init AddressList~~
// List is empty
// Size of AddressList: 0

// ~~After Insert Addresses~~
// Size of AddressList: 3
// 1 name: Sakura, phone_number: 090-1234-5678, email_address: sakura@example.com
// 2 name: Kaede, phone_number: 080-1234-5678, email_address: kaede@example.com
// 3 name: Yuzu, phone_number: 070-1234-5678, email_address: yuzu@example.com

// ~~After Remove Address~~
// Size of AddressList: 2
// 1 name: Sakura, phone_number: 090-1234-5678, email_address: sakura@example.com
// 2 name: Yuzu, phone_number: 070-1234-5678, email_address: yuzu@example.com
