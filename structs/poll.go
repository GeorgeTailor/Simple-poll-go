package structs

type Poll struct {
	ID         int
	Question   string
	OptionList []Option
}
