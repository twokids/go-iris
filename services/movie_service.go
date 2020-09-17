package services

import (
	"fmt"
	"go-iris/repositories"
)

type MovieService interface {
	ShowMovieName() string
}

type MovieServiceManager struct {
	repo repositories.MovieRepository
}

func NewMovieServiceManager(repo repositories.MovieRepository) MovieService {
	return &MovieServiceManager{repo: repo}
}

func (m *MovieServiceManager) ShowMovieName() string {
	fmt.Println("国庆风景区推荐：" + m.repo.GetMovieName())
	return "国庆风景区推荐：" + m.repo.GetMovieName()
}
