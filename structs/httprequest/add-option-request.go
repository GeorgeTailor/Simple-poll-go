package httprequest

import "../../structs"

type AddOptionRequest struct {
	PollID int
	Option structs.Option
}
