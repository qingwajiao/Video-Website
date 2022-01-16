package main

import (
	"fmt"
	"sync"
)

func Myprint() int {
	return 5
}

func getMap(m *sync.Map) {
	s := []string{"a", "b", "c", "d"}

	for index, value := range s {
		//fmt.Println("v:",v,"i:",i)

		m.Store(index, value)
	}
}

func main() {
	//fmt.Println(Myprint())
	ma := &sync.Map{}
	getMap(ma)
	ma.Range(func(key, value interface{}) bool {
		fmt.Printf("key:%v,value:%v\n", key, value)
		return true
	})
	//fmt.Println()
}
