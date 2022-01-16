package usecase

import "placeholder/domain"

//domain.PoemUsecase implementation
type poemUsecase struct {
	poemRepository domain.PoemRepository
}

func NewPoemUsecase(poemRepository domain.PoemRepository) domain.PoemUsecase {
	return &poemUsecase{
		poemRepository: poemRepository,
	}
}

func (p *poemUsecase) FetchUsecase() []domain.Poem {
	return p.poemRepository.FetchPoems()
}

func (p *poemUsecase) FetchRandomUsecase() domain.Poem {
	return p.poemRepository.FetchRandomPoem()
}

func (p *poemUsecase) FetchStaticUsecase() domain.Poem {
	return p.poemRepository.FetchStaticPoem()
}

func (p *poemUsecase) Store(poem *domain.Poem) error {
	return p.poemRepository.Store(poem)
}
