package mocks

import (
	"github.com/stretchr/testify/mock"
	"github.com/thiendwiananda/golang-simple-microservice/domain"
)

type MovieRepository struct {
	mock.Mock
}

func (_mr *MovieRepository) GetById(id string) (result domain.Movie, err error) {
	return
}

func (_mr *MovieRepository) SearchByTitleWithPagination(title string, page string) (result []domain.Movie, err error) {
	return
}
