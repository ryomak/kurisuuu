package interactor

import (
	"github.com/ryomak/kurisuuu/domain"
	"os"
	"io/ioutil"
	"path/filepath"
	"github.com/ryomak/kurisuuu/core"
)

func ResMovies()([]domain.Movie,error){
	return GetMovies()
}

func GetMovies()([]domain.Movie,error){
	moviePath := os.Getenv("GOPATH") + "/src/github.com/ryomak/kurisuuu/public/movie"
	list, err := ioutil.ReadDir(moviePath)
	movies := make([]domain.Movie,0)
	for _, fileInfo := range list {
		movie := domain.Movie{
			Name:core.RemoveExtension(core.GetNameBySnakeCase(fileInfo.Name())),
			Path:filepath.Join("public/movie",fileInfo.Name()),
		}
		movies = append(movies,movie)
	}
	return movies,err
}


