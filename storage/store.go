package storage

import (
	"../structs"
)

var store structs.OptionList
var currentMaxId = 1

func GetPollOptions() structs.OptionList {
	return store
}

func Get() structs.OptionList {
	return store
}

func Add(message structs.Option) int {
	message.ID = currentMaxId
	currentMaxId++
	store = append(store, message)
	return message.ID
}

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
