package storage

import (
	"../structs"
)

// Store - local DB
var Store structs.PollList

//currentMaxPollID - local 'sequence' for polls
var currentMaxPollID = 1

//currentMaxOptionID - local 'sequence' for options
var currentMaxOptionID = 1

// ListPolls returns full storage with
// all available polls
func ListPolls() structs.PollList {
	return Store
}

// GetPoll returns specific poll by given PollId
// or returns empty struct when poll with such id does not exist in storage
func GetPoll(PollID int) (structs.Poll, bool) {
	var str = Store
	for i := 0; i < len(str); i++ {
		if str[i].ID == PollID {
			return str[i], true
		}
	}
	empty := new(structs.Poll)
	return (*empty), false
}

// SelectPollOption increments the number of how many times the
// option by given ID was selected in the poll with given PollID
func SelectPollOption(PollID int, OptionID int) bool {
	var str = Store
	for i := 0; i < len(str); i++ {
		if str[i].ID == PollID {
			var poll = str[i].OptionList
			for j := 0; j < len(poll); j++ {
				if poll[j].ID == OptionID {
					poll[j].SelectedCount++
					Store = str
					return true
				}
			}
		}
	}
	return false
}

// AddPoll adds Poll struct to the storage
// and returns its ID
func AddPoll(question string, optionList []structs.Option) int {
	poll := new(structs.Poll)
	poll.Question = question
	poll.ID = currentMaxPollID
	currentMaxPollID++
	for i := 0; i < len(optionList); i++ {
		optionList[i].ID = currentMaxOptionID
		optionList[i].PollID = poll.ID
		currentMaxOptionID++
	}
	poll.OptionList = optionList

	Store = append(Store, *poll)
	return poll.ID
}

// AddOption adds Option struct to the poll with given ID in the storage
// returns poll ID
func AddOption(option structs.Option, pollID int) (int, bool) {
	var str = Store
	for i := 0; i < len(str); i++ {
		if str[i].ID == pollID {
			var optionList = str[i].OptionList
			option.ID = currentMaxOptionID
			currentMaxOptionID++
			optionList = append(optionList, option)
			str[i].OptionList = optionList
			Store = str
			return option.ID, true
		}
	}
	return 0, false
}

// RemovePoll - remove poll from storage by given ID
// returns true when poll was found and removed or false when poll was not found
func RemovePoll(pollID int) bool {
	for i, poll := range Store {
		if poll.ID == pollID {
			Store = append(Store[:i], Store[i+1:]...)
			return true
		}
	}
	return false
}

// RemoveOption - remove option from poll by given ID
// returns true when option was found and removed or false when option was not found
func RemoveOption(pollID int, optionID int) bool {
	var str = Store
	for i, poll := range Store {
		if poll.ID == pollID {
			var optionList = poll.OptionList
			for j, option := range optionList {
				if option.ID == optionID {
					optionList = append(optionList[:j], optionList[j+1:]...)
					poll.OptionList = optionList

					str[i] = poll
					Store = str
					return true
				}
			}
		}
	}
	return false
}
