package model

type Banner struct {
	Id int `json:"banner_id,omitempty" db:"id"`
	// Title string `json:"title" db:"title"`
	// Text string `json:"text" db:"text"`
	// Url string `json:"url" db:"url"`
	FeatureId int `json:"feature_id,omitempty" db:"feature_id"`
	Content map[string]interface{} `json:"content" db:"-"`
	JSONContent string `json:"-" db:"content"`
	TagIds []int `json:"tag_ids,omitempty" db:"-"`
	IsActive bool `json:"is_active" db:"is_active"`
	CreatedAt string `json:"created_at,omitempty" db:"created_at"`
	UpdatedAt string `json:"updated_at,omitempty" db:"updated_at"`
}

func (f Banner) IsEmpty() bool {
	return false
}