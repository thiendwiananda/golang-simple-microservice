package http

import (
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/thiendwiananda/golang-simple-microservice/domain"
)

type MovieHandler struct {
	MovieUsecase domain.MovieUsecase
}

func NewMovieHandler(server *echo.Echo, movieUsecase domain.MovieUsecase) {
	controller := &MovieHandler{
		MovieUsecase: movieUsecase,
	}

	server.GET("/api/movies/:id", controller.GetById)
	server.GET("/api/movies/searches", controller.FetchWithPagination)
}

func (_m *MovieHandler) GetById(context echo.Context) error {
	payload, _ := _m.MovieUsecase.GetById(context.Param("id"))

	return context.JSON(http.StatusOK, payload)
}

func (_m *MovieHandler) FetchWithPagination(context echo.Context) error {
	searchWord := strings.ReplaceAll(context.QueryParam("searchword"), " ", "+")

	payload, _ := _m.MovieUsecase.SearchByTitleWithPagination(searchWord, context.QueryParam("pagination"))

	return context.JSON(http.StatusOK, payload)
}
