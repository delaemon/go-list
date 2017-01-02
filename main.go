package main

import (
	"container/list"
	"fmt"
)

type Smap struct {
	LastName string
}

type V6 struct {
	FirstName string
}

func main() {
	data := list.New()
	data.PushFront(V6{"Go"})
	data.PushFront(V6{"Ken"})
	data.PushFront(V6{"Junichi"})
	data.PushBack(Smap{"Kusanagi"})
	data.PushBack(Smap{"Katori"})
	data.PushBack(Smap{"Inagaki"})
	data.PushBack(Smap{"Kimura"})
	data.PushBack(Smap{"Nakai"})
	fmt.Printf("First: %s\n", data.Front().Value)
	fmt.Printf("Last: %s\n", data.Back().Value)

	target := Smap{"Kimura"}
	e := data.Front()
	for e != nil {
		fmt.Printf("check: %s \n", e.Value)
		if target == e.Value {
			fmt.Println("Hit")
			break
		}
		e = e.Next()
	}
}
