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

func (u *poemUsecase) FetchUsecase() []domain.Poem {
	return u.poemRepository.FetchPoems()
}

func (u *poemUsecase) FetchRandomUsecase() domain.Poem {
	return u.poemRepository.FetchRandomPoem()
}

func (u *poemUsecase) FetchStaticUsecase() domain.Poem {
	return u.poemRepository.FetchStaticPoem()
}

func (u *poemUsecase) Store(poem *domain.Poem) error {
	return u.poemRepository.Store(poem)
}
