package main

import "fmt"

type collection interface {
	createIterator() iterator
}

type iterator interface {
	hasNext() bool
	getNext() *data
}

type data struct {
	some  int
	some2 string
}

type dataIterator struct {
	index int
	data  []*data
}

// Некототорая коллекция данных
type dataCollection struct {
	data []*data
}

func (u *dataCollection) createIterator() iterator {
	return &dataIterator{
		data: u.data,
	}
}

func (u *dataIterator) hasNext() bool {
	if u.index < len(u.data) {
		return true
	}
	return false

}
func (u *dataIterator) getNext() *data {
	if u.hasNext() {
		data := u.data[u.index]
		u.index++
		return data
	}
	return nil
}

func main() {
	data1 := &data{
		some:  1,
		some2: "some",
	}
	data2 := &data{
		some:  2,
		some2: "some",
	}

	dataCollection := &dataCollection{
		data: []*data{data1, data2},
	}

	iterator := dataCollection.createIterator()

	for iterator.hasNext() {
		data := iterator.getNext()
		fmt.Printf("data is %+v\n", data)
	}
}
