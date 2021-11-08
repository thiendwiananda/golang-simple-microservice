package usecase

import (
	"sync"

	"github.com/thiendwiananda/golang-simple-microservice/domain"
)

type MovieUsecase struct {
	MovieRepository domain.MovieRepository
	LogRepository   domain.LogRepository
}

func NewMovieUsecase(movieRepo domain.MovieRepository, logRepo domain.LogRepository) domain.MovieUsecase {
	return &MovieUsecase{
		MovieRepository: movieRepo,
		LogRepository:   logRepo,
	}
}

func (_mu *MovieUsecase) GetById(id string) (result domain.Movie, err error) {
	payload, err := _mu.MovieRepository.GetById(id)
	if err != nil {
		return
	}

	result = payload

	return
}

func (_mu *MovieUsecase) SearchByTitleWithPagination(title string, page string) (result []domain.Movie, err error) {
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		_mu.LogRepository.Store(&domain.Log{
			Content: "Logging search detail",
		})
	}()

	result, err = _mu.MovieRepository.SearchByTitleWithPagination(title, page)
	if err != nil {
		return
	}

	movieChan := make(chan domain.Movie, 100)

	for _, m := range result {
		wg.Add(1)
		go func(m domain.Movie) {
			defer wg.Done()
			movie, _ := _mu.MovieRepository.GetById(m.Id)
			movieChan <- movie
		}(m)
	}

	for i := range result {
		result[i] = <-movieChan
	}

	wg.Wait()

	return
}
