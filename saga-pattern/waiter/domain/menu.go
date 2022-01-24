package domain

type Menu struct {
	Foods  []Item `json:"foods"`
	Drinks []Item `json:"drinks"`
}

type Item struct {
	Id   uint8  `json:"id"`
	Name string `json:"name"`
}

type MenuUsecase interface {
	ShowMenu() Menu
}

type MenuRepository interface {
	ShowMenu() Menu
}
