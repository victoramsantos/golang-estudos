package repository

var _POEM = Poem{
	Content: "Nem alegre nem triste, poeta",
	Author:  "Someone",
}

func GetStaticContent() Poem {
	return _POEM
}
