package repository

import (
	"math/rand"
	"placeholder/domain"
	"time"
)

type poemRepository struct {
	poems []domain.Poem
}

func NewPoemRepository() domain.PoemRepository {
	poemRepository := poemRepository{
		poems: make([]domain.Poem, 0),
	}
	poemRepository.poems = append(
		poemRepository.poems,
		domain.Poem{
			Content: "Nem alegre nem triste, poeta",
			Author:  "Someone",
		},
	)
	return &poemRepository
}

func (pr *poemRepository) FetchPoems() []domain.Poem {
	return pr.poems
}

func (pr *poemRepository) FetchRandomPoem() domain.Poem {
	randomizer := sanitizeRand()
	var randPos = randomizer.Intn(len(pr.poems))
	return pr.poems[randPos]
}

func (*poemRepository) FetchStaticPoem() domain.Poem {
	return domain.Poem{
		Content: "Nem alegre nem triste, poeta",
		Author:  "Someone",
	}
}

func (pr *poemRepository) Store(poem *domain.Poem) error {
	pr.poems = append(
		pr.poems,
		*poem,
	)
	return nil
}

func sanitizeRand() *rand.Rand {
	return rand.New(rand.NewSource(time.Now().UnixNano()))
}
