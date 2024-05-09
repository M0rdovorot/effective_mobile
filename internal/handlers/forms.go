package handlers

type LoginForm struct{
	Body struct {
		Token string `json:"token"`
	} `json:"body"`
}

type EmptyForm struct{}

func (f EmptyForm) IsEmpty() bool {
	return true
}