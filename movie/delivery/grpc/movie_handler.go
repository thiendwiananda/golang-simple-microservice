package grpc

import (
	context "context"

	"github.com/thiendwiananda/golang-simple-microservice/domain"
	"google.golang.org/grpc"
)

func NewMovieServerGrpc(grpcServer *grpc.Server, movieUsecase domain.MovieUsecase) {
	newServer := &server{
		movieUsecase: movieUsecase,
	}

	RegisterMovieHandlerServer(grpcServer, newServer)
}

type server struct {
	UnimplementedMovieHandlerServer
	movieUsecase domain.MovieUsecase
}

func (server *server) GetMovie(ctx context.Context, req *SingleRequest) (movie *Movie, err error) {
	id := ""

	if req != nil {
		id = req.Id
	}

	data, err := server.movieUsecase.GetById(id)

	if err != nil {
		return
	}

	movie = &Movie{
		ImdbID:     data.Id,
		Title:      data.Title,
		Year:       data.Year,
		Rated:      data.Rated,
		Released:   data.Released,
		Runtime:    data.Runtime,
		Genre:      data.Genre,
		Director:   data.Director,
		Writer:     data.Writer,
		Actors:     data.Actors,
		Plot:       data.Plot,
		Language:   data.Language,
		Country:    data.Country,
		Awards:     data.Awards,
		Poster:     data.Poster,
		Ratings:    []*MovieRatings{},
		Metascore:  data.Metascore,
		ImdbRating: data.ImdbRating,
		ImdbVotes:  data.ImdbVotes,
		Type:       data.Type,
		DVD:        data.Dvd,
		BoxOffice:  data.BoxOffice,
		Production: data.Production,
		Website:    data.Website,
		Response:   data.Response,
	}

	for _, rating := range data.Ratings {
		movie.Ratings = append(movie.Ratings, &MovieRatings{
			Source: rating.Source,
			Value:  rating.Value,
		})
	}

	return
}

func (server *server) SearchMovie(ctx context.Context, req *SearchRequest) (movieList *MovieList, err error) {
	searchWord := ""
	pagination := ""

	if req != nil {
		searchWord = req.Searchword
		pagination = req.Pagination
	}

	data, err := server.movieUsecase.SearchByTitleWithPagination(searchWord, pagination)

	for _, movie := range data {
		data := &Movie{
			ImdbID:     movie.Id,
			Title:      movie.Title,
			Year:       movie.Year,
			Rated:      movie.Rated,
			Released:   movie.Released,
			Runtime:    movie.Runtime,
			Genre:      movie.Genre,
			Director:   movie.Director,
			Writer:     movie.Writer,
			Actors:     movie.Actors,
			Plot:       movie.Plot,
			Language:   movie.Language,
			Country:    movie.Country,
			Awards:     movie.Awards,
			Poster:     movie.Poster,
			Ratings:    []*MovieRatings{},
			Metascore:  movie.Metascore,
			ImdbRating: movie.ImdbRating,
			ImdbVotes:  movie.ImdbVotes,
			Type:       movie.Type,
			DVD:        movie.Dvd,
			BoxOffice:  movie.BoxOffice,
			Production: movie.Production,
			Website:    movie.Website,
			Response:   movie.Response,
		}

		for _, rating := range movie.Ratings {
			data.Ratings = append(data.Ratings, &MovieRatings{
				Source: rating.Source,
				Value:  rating.Value,
			})
		}

		movieList.Movies = append(movieList.Movies, data)
	}

	return
}
