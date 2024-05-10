package handlers

type EmptyForm struct{}

func (f EmptyForm) IsEmpty() bool {
	return true
}

type RegNumsForm struct{
	RegNums []string `json:"regNums"`
}

func (f RegNumsForm) IsEmpty() bool {
	return false
}

type PatchMapForm map[string]any

func (f PatchMapForm) IsEmpty() bool {
	return false
}
