package structs

type Option struct {
	ID            int    `json:"id,omitempty"`
	Option        string `json:"option"`
	SelectedCount int    `json:"selectedCount"`
}
