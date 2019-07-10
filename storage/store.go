package storage

import (
	"../structs"
)

var store structs.MessageList
var currentMaxID = 1

//Get is ...
func Get() structs.MessageList {
	return store
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
