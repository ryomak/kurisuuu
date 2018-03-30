package interactor

import (
	"github.com/ryomak/kurisuuu/domain"
	"io/ioutil"
	"path/filepath"
	"github.com/ryomak/kurisuuu/core"
	"github.com/ryomak/kurisuuu/core/const_variable"
)

func ResMovies()([]domain.Movie,error){
	return GetMovies()
}

func GetMovies()([]domain.Movie,error){
	list, err := ioutil.ReadDir(const_variable.MoviePath)
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


