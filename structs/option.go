package structs

type Option struct {
	ID            int `json:"id,omitempty"`
	PollID        int
	Option        string `json:"option"`
	SelectedCount int    `json:"selectedCount"`
}
