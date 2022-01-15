package service

import "placeholder/repository"

func GetStaticContent() repository.Poem {
	return repository.GetStaticContent()
}

func GetDynamicContent() repository.Poem {
	return repository.GetDynamicContent()
}
