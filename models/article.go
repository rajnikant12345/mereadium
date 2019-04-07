package models

type Article struct {
	ID      string   `json:"id"`
	Writer  string   `json:"writer"`
	Heading string   `json:"heading"`
	Tags    []string `json:"tags,omitempty"`
	URI     string   `json:"uri"`
}

func (m *Model) GetArticleList() []Article {
	a1 := Article{}
	a2 := Article{}
	a1.ID = "rajni"
	a2.ID = "sharad"
	return []Article{a1, a2}
}
