package storage

import (
	"../structs"
)

var Store structs.OptionList
var currentMaxId = 1

func GetPollOptions() structs.OptionList {
	return Store
}

func SelectPollOption(ID int) (structs.Option, bool) {
	var str = Store
	for i := 0; i < len(str); i++ {
		if str[i].ID == ID {
			str[i].SelectedCount++
			Store = str
			return (str)[i], true
		}
	}
	empty := new(structs.Option)
	return (*empty), false
}

func Add(message structs.Option) int {
	message.ID = currentMaxId
	currentMaxId++
	Store = append(Store, message)
	return message.ID
}

func Remove(id int) bool {
	index := -1

	for i, message := range Store {
		if message.ID == id {
			index = i
		}
	}

	if index != -1 {
		Store = append(Store[:index], Store[index+1:]...)
	}

	// Returns true if item was found & removed
	return index != -1
}
