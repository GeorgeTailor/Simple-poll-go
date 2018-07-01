package httprequest

import "../../structs"

type AddPollRequest struct {
	Question   string
	OptionList []structs.Option
}
