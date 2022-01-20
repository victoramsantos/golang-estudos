package domain

type Poem struct {
	Content string `json:"content" validate:"required"`
	Author  string `json:"author" validate:"required"`
}

type PoemUsecase interface {
	FetchUsecase() []Poem
	FetchRandomUsecase() Poem
	FetchStaticUsecase() Poem
	Store(*Poem) error
}

type PoemRepository interface {
	FetchPoems() []Poem
	FetchRandomPoem() Poem
	FetchStaticPoem() Poem
	Store(*Poem) error
}
