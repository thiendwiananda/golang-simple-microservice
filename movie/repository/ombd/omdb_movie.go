package omdb

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"

	"github.com/thiendwiananda/golang-simple-microservice/domain"
)

type OmdbMovieRepository struct {
	ApiKey  string
	baseUrl string
}

func NewOmdbMoviewRepository(apikey string) domain.MovieRepository {
	return &OmdbMovieRepository{
		ApiKey:  apikey,
		baseUrl: "https://www.omdbapi.com/?apikey=",
	}
}

func (_o *OmdbMovieRepository) GetById(id string) (result domain.Movie, err error) {
	request, err := http.Get(_o.baseUrl + _o.ApiKey + "&i=" + id)
	if err != nil {
		return
	}

	body, err := ioutil.ReadAll(request.Body)
	if err != nil {
		return
	}

	var response map[string]json.RawMessage

	err = json.Unmarshal(body, &response)
	if err != nil {
		return
	}

	if _, exist := response["Error"]; exist {
		err = errors.New(string(response["Error"]))
		return
	}

	json.Unmarshal(body, &result)

	return
}

func (_o *OmdbMovieRepository) SearchByTitleWithPagination(title string, page string) (result []domain.Movie, err error) {
	request, err := http.Get(_o.baseUrl + _o.ApiKey + "&s=" + title + "&page=" + page)
	if err != nil {
		return
	}

	body, err := ioutil.ReadAll(request.Body)
	if err != nil {
		return
	}

	var response map[string]json.RawMessage

	err = json.Unmarshal(body, &response)
	if err != nil {
		return
	}

	if _, exist := response["Error"]; exist {
		err = errors.New(string(response["Error"]))
		return
	}

	json.Unmarshal(response["Search"], &result)
	if err != nil {
		return
	}

	return
}
