package storage

import (
	"math"

	"../structs"
)

var store structs.MessageList
var currentMaxID = 1

//Get is ...
func Get(input float64) float64 {
	for i := uint64(0); i < 4000; i++ {
		Factorial(uint64(i))
	}
	return math.Tan(input)
}

//Add is ...
func Add(message structs.Message) int {
	message.ID = currentMaxID
	currentMaxID++
	store = append(store, message)
	return message.ID
}

//Remove is ...
func Remove(id int) bool {
	index := -1

	for i, message := range store {
		if message.ID == id {
			index = i
		}
	}

	if index != -1 {
		store = append(store[:index], store[index+1:]...)
	}

	// Returns true if item was found & removed
	return index != -1
}

func Factorial(n uint64) (result uint64) {
	if n > 0 {
		result = n * Factorial(n-1)
		return result
	}
	return 1
}
